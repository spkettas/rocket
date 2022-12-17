package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeOver
)

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
	FlushCnt        = 0 // &&&
)

type Game struct {
	Mode  Mode    // 模式
	Input *Input  // 键盘输入
	Ship  *Ship   // 飞船
	Cfg   *Config // 配置

	bullets map[*Bullet]struct{} // 多个子弹
	aliens  map[*Alien]struct{}  // 多个外星人

	FailCount int // 失败次数
}

// CreateFonts 创建字体
func (g *Game) CreateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.Cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.Cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.Cfg.SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) init() {
	g.CreateAliens()
	g.CreateFonts()
}

func NewGame() *Game {
	cfg := loadConfig()
	log.Printf("cfg=%v", cfg)

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	g := &Game{
		Input:     &Input{},
		Ship:      NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		Cfg:       cfg,
		bullets:   make(map[*Bullet]struct{}),
		aliens:    make(map[*Alien]struct{}),
		FailCount: 0,
	}

	// 创建一批外星人、字体
	g.init()
	return g
}

// addBullet 添加子弹
func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

// addBullet 添加子弹
func (g *Game) addAlien(alien *Alien) {
	g.aliens[alien] = struct{}{} // &&& 赋值方式有意思
}

func (g *Game) CreateAliens() {
	// 创建二行外星人
	for row := 0; row < 2; row++ {
		alien := NewAlien(g.Cfg)

		availableSpaceX := g.Cfg.ScreenWidth - 2*alien.Width()
		numAliens := availableSpaceX / (2 * alien.Width())

		for i := 0; i < numAliens; i++ {
			alien = NewAlien(g.Cfg)
			alien.x = float64(alien.Width() + 2*alien.Width()*i)
			alien.y = float64(alien.Height()*row) * 1.5

			g.addAlien(alien)
		}
	}
}

func (g *Game) CheckCollision() {
	for alien := range g.aliens {
		// 子弹是否击中外星人
		for bullet := range g.bullets {
			if CheckCollision(bullet, alien) {
				delete(g.aliens, alien)
				delete(g.bullets, bullet)
			}
		}

		// 外星人是否移出屏幕
		if alien.OutOfScreen(g.Cfg) {
			g.FailCount++
			delete(g.aliens, alien)
			continue
		}

		// 外星人是否击中飞船
		if CheckCollision(alien, g.Ship) {
			g.FailCount++
			delete(g.aliens, alien)
			continue
		}
	}
}

// update 更新参数，先update再draw
func (g *Game) Update() error {
	switch g.Mode {
	case ModeTitle: // 开始游戏
		if g.Input.IsSpacePressed() {
			g.Mode = ModeGame
		}
	case ModeGame: // 游戏中
		// 更新子弹的垂直移动
		for bullet := range g.bullets {
			if bullet.OutOfScreen() { // 超出屏幕
				delete(g.bullets, bullet)
			} else {
				bullet.y -= bullet.SpeedFactor
			}
		}

		// 更新外星人轨迹
		for alien := range g.aliens {
			alien.y += alien.SpeedFactor
		}

		// 更新飞船
		g.Input.Update(g)

		// 检测碰撞
		g.CheckCollision()
	case ModeOver: // 游戏结束
		if g.Input.IsSpacePressed() {
			g.init()
			g.Mode = ModeTitle
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.Cfg.BgColor)
	//ebitenutil.DebugPrint(screen, g.Input.Msg)

	var titleTexts []string
	var texts []string

	switch g.Mode {
	case ModeTitle:
		titleTexts = []string{"ALIEN INVASION"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}
	case ModeGame:
		// 飞船
		g.Ship.Draw(screen, g.Cfg)

		// 子弹
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}

		// 外星人
		for alien := range g.aliens {
			alien.Draw(screen)
		}
	case ModeOver:
		texts = []string{"", "GAME OVER!"}
	}

	for i, l := range titleTexts {
		x := (g.Cfg.ScreenWidth - len(l)*g.Cfg.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.Cfg.TitleFontSize, color.White)
	}

	for i, l := range texts {
		x := (g.Cfg.ScreenWidth - len(l)*g.Cfg.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.Cfg.FontSize, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Cfg.ScreenWidth, g.Cfg.ScreenHeight
}
