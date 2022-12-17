package main

import (
	_ "image/png"
	//_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Ship 飞船
type Ship struct {
	GameObject

	image *ebiten.Image
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}

	//op.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
	op.GeoM.Translate(ship.X(), ship.Y())

	screen.DrawImage(ship.image, op)
}

func NewShip(screenWidth, screenHeight int) *Ship {
	path := "./images/ship.png"
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	log.Printf("width=%v height=%v", width, height)

	ship := &Ship{
		GameObject: GameObject{
			width:  width,
			height: height,
			x:      float64(screenWidth-width) / 2, // 居中
			y:      float64(screenHeight - height), // 居中
		},
		image: img,
	}

	return ship
}
