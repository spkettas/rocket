package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	input *Input
	ship  *Ship
	cfg   *Config
}

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	return &Game{
		input: &Input{},
		ship:  NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		cfg:   cfg,
	}
}

func (g *Game) Update() error {
	g.input.Update(g.ship, g.cfg)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})
	//ebitenutil.DebugPrint(screen, g.input.msg)

	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen, g.cfg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth / 2, g.cfg.ScreenHeight / 2
}
