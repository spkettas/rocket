package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Image       *ebiten.Image
	Width       int
	Height      int
	X           float64
	Y           float64
	SpeedFactor float64
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.X, bullet.Y)

	screen.DrawImage(bullet.Image, op)
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)

	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	return &Bullet{
		Image:       img,
		Width:       cfg.BulletWidth,
		Height:      cfg.BulletHeight,
		X:           ship.X + float64(ship.width-cfg.BulletWidth)/2,
		Y:           float64(cfg.ScreenHeight - ship.height - cfg.BulletHeight),
		SpeedFactor: cfg.BulletSpeedFactor,
	}
}
