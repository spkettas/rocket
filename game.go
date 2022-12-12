package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Input *Input
	Ship  *Ship
	Cfg   *Config
}

func NewGame() *Game {
	cfg := loadConfig()
	log.Printf("cfg=%v", cfg)

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	return &Game{
		Input: &Input{},
		Ship:  NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		Cfg:   cfg,
	}
}

func (g *Game) Update() error {
	g.Input.Update(g.Ship, g.Cfg)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.Cfg.BgColor)
	ebitenutil.DebugPrint(screen, g.Input.Msg)

	g.Ship.Draw(screen, g.Cfg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Cfg.ScreenWidth, g.Cfg.ScreenHeight
	//return g.Cfg.ScreenWidth / 2, g.Cfg.ScreenHeight / 2
}
