package main

import (
	"log"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	x      float64 // x坐标
	y      float64 // y坐标
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
	screen.DrawImage(ship.image, op)
}

func NewShip(screenWidth, screenHeight int) *Ship {
	path := "./images/ship.png"
	//path := "/Users/kanesun/mycode/rocket/images/ship.png"

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
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
	}

	return ship
}
