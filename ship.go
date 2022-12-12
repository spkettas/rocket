package main

import (
	_ "image/png"
	//_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	X      float64 // x坐标
	Y      float64 // y坐标
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}

	//op.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
	op.GeoM.Translate(ship.X, ship.Y)

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
		image:  img,
		width:  width,
		height: height,
		X:      float64(screenWidth-width) / 2, // 居中
		Y:      float64(screenHeight - height), // 居中
	}

	return ship
}
