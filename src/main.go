package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	img.Init(img.INIT_PNG)
	ttf.Init()

	game := NewGame()
	game.Run()

}
