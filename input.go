package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	Msg            string
	lastBulletTime time.Time
}

func (i *Input) IsSpacePressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace)
}

func (i *Input) Update(g *Game) {
	now := time.Now()

	if ebiten.IsKeyPressed(ebiten.KeySpace) { // space
		// 添加子弹
		if len(g.bullets) < g.Cfg.MaxBulletNum &&
			now.Sub(i.lastBulletTime).Milliseconds() > g.Cfg.BulletInterval {
			bullet := NewBullet(g.Cfg, g.Ship)
			g.addBullet(bullet)

			i.lastBulletTime = now
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) { // <-
		i.Msg = "left pressed"
		g.Ship.X -= g.Cfg.ShipSpeedFactor

		// 边界判断
		if g.Ship.X < -float64(g.Ship.width)/2 {
			g.Ship.X = -float64(g.Ship.width) / 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) { // ->
		i.Msg = "right pressed"
		g.Ship.X += g.Cfg.ShipSpeedFactor

		if g.Ship.X > float64(g.Cfg.ScreenWidth)-float64(g.Ship.width)/2 {
			g.Ship.X = float64(g.Cfg.ScreenWidth) - float64(g.Ship.width)/2
		}
	}
}
