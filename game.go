package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Input   *Input               // 键盘输入
	Ship    *Ship                // 飞船
	Cfg     *Config              // 配置
	bullets map[*Bullet]struct{} // 多个子弹
}

func NewGame() *Game {
	cfg := loadConfig()
	log.Printf("cfg=%v", cfg)

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	return &Game{
		Input:   &Input{},
		Ship:    NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		Cfg:     cfg,
		bullets: make(map[*Bullet]struct{}),
	}
}

// addBullet 添加子弹
func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

func (g *Game) Update() error {
	// 更新飞船
	g.Input.Update(g)

	// 更新子弹的垂直移动
	for bullet := range g.bullets {
		bullet.Y -= bullet.SpeedFactor
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.Cfg.BgColor)
	ebitenutil.DebugPrint(screen, g.Input.Msg)

	// 飞船
	g.Ship.Draw(screen, g.Cfg)

	// 子弹
	for bullet := range g.bullets {
		bullet.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Cfg.ScreenWidth, g.Cfg.ScreenHeight
	//return g.Cfg.ScreenWidth / 2, g.Cfg.ScreenHeight / 2
}
