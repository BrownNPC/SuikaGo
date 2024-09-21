package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	img.Init(img.INIT_PNG)

	game := NewGame()
	game.Run()

}
