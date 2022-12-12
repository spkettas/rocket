// ebitengine gui
// https://darjun.github.io/2022/11/15/godailylib/ebiten1/
// 2022/12/12

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
