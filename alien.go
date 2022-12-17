package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Alien 外星人
type Alien struct {
	GameObject

	Image       *ebiten.Image
	SpeedFactor float64
}

func NewAlien(cfg *Config) *Alien {
	img, _, err := ebitenutil.NewImageFromFile("./images/alien.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	return &Alien{
		Image: img,
		GameObject: GameObject{ // 匿名结构赋值
			width:  width,
			height: height,
			x:      0,
			y:      0,
		},
		SpeedFactor: cfg.AlienSpeedFactor,
	}
}

func (alien *Alien) OutOfScreen(cfg *Config) bool {
	return alien.Y() > float64(cfg.ScreenHeight)+float64(alien.Height())
}

func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.X(), alien.Y())

	screen.DrawImage(alien.Image, op)
}
