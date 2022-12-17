package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	GameObject

	Image       *ebiten.Image
	SpeedFactor float64
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.X(), bullet.Y())

	screen.DrawImage(bullet.Image, op)
}

func (bullet *Bullet) OutOfScreen() bool {
	return bullet.Y() < -float64(bullet.Height())
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)

	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	return &Bullet{
		Image: img,
		GameObject: GameObject{
			width:  cfg.BulletWidth,
			height: cfg.BulletHeight,
			x:      ship.X() + float64(ship.Width()-cfg.BulletWidth)/2,
			y:      float64(cfg.ScreenHeight - ship.Height() - cfg.BulletHeight),
		},
		SpeedFactor: cfg.BulletSpeedFactor,
	}
}
