package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	msg string
}

func (i *Input) Update(ship *Ship, cfg *Config) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		fmt.Println("←←←←←←←←←←←←←←←←←←←←←←←")
		i.msg = "left pressed"
		ship.x -= cfg.ShipSpeedFactor
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		fmt.Println("→→→→→→→→→→→→→→→→→→→→→→→")
		i.msg = "right pressed"
		ship.x += cfg.ShipSpeedFactor
	}
}
