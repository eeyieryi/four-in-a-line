package main

import (
	"log"

	"github.com/eeyieryi/four-in-a-row/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	myGame := game.Game{}
	myGame.Setup()
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle(myGame.Title)
	if err := ebiten.RunGame(&myGame); err != nil {
		if err == game.ErrTerminated {
			return
		}
		log.Fatal(err)
	}
}
