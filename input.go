package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	Msg string
}

func (i *Input) Update(ship *Ship, cfg *Config) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		log.Printf("press key left")

		i.Msg = "left pressed"
		ship.X -= cfg.ShipSpeedFactor
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		log.Printf("press key right")

		i.Msg = "right pressed"
		ship.X += cfg.ShipSpeedFactor
	}
}
