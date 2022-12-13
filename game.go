package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Input *Input  // 键盘输入
	Ship  *Ship   // 飞船
	Cfg   *Config // 配置

	bullets map[*Bullet]struct{} // 多个子弹
	aliens  map[*Alien]struct{}  // 多个外星人
}

func NewGame() *Game {
	cfg := loadConfig()
	log.Printf("cfg=%v", cfg)

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	g := &Game{
		Input:   &Input{},
		Ship:    NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		Cfg:     cfg,
		bullets: make(map[*Bullet]struct{}),
		aliens:  make(map[*Alien]struct{}),
	}

	// 创建一批外星人
	g.CreateAliens()
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

		availableSpaceX := g.Cfg.ScreenWidth - 2*alien.Width
		numAliens := availableSpaceX / (2 * alien.Width)

		for i := 0; i < numAliens; i++ {
			alien = NewAlien(g.Cfg)
			alien.X = float64(alien.Width + 2*alien.Width*i)
			alien.Y = float64(alien.Height*row) * 1.5

			g.addAlien(alien)
		}
	}
}

func (g *Game) Update() error {
	// 更新飞船
	g.Input.Update(g)

	// 更新子弹的垂直移动
	for bullet := range g.bullets {
		if bullet.OutOfScreen() { // 超出屏幕
			delete(g.bullets, bullet)
		} else {
			bullet.Y -= bullet.SpeedFactor
		}
	}

	// 更新外星人轨迹
	for alien := range g.aliens {
		alien.Y += alien.SpeedFactor
	}
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.Cfg.BgColor)
	//ebitenutil.DebugPrint(screen, g.Input.Msg)

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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Cfg.ScreenWidth, g.Cfg.ScreenHeight
	//return g.Cfg.ScreenWidth / 2, g.Cfg.ScreenHeight / 2
}
