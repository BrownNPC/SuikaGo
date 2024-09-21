package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type CSprite struct {
	texture *sdl.Texture
	width   int32
	height  int32
	offsetX int32
	offsetY int32
}

func NewCSprite(Renderer *sdl.Renderer, texturePath string, width int32, height int32, offsetX int32, offsetY int32) *CSprite {
	c := &CSprite{
		width:   width,
		height:  height,
		offsetX: offsetX,
		offsetY: offsetY,
	}
	c.texture, _ = img.LoadTexture(Renderer, texturePath)

	return c
}

func (c *CSprite) GetTexture() *sdl.Texture {
	return c.texture
}

func (c *CSprite) GetWidth() int32 {
	return c.width
}

func (c *CSprite) GetHeight() int32 {
	return c.height
}

func (c *CSprite) Destroy() {
	c.texture.Destroy()
}

func (c *CSprite) Render(renderer *sdl.Renderer, x int32, y int32, angle float64) {
	renderer.CopyEx(c.texture, nil, &sdl.Rect{
		X: x + c.offsetX,
		Y: y + c.offsetY,
		W: c.width,
		H: c.height,
	}, angle, &sdl.Point{X: c.width / 2, Y: c.height / 2}, sdl.FLIP_NONE)
}

func (c *CSprite) Free() {
	c.texture.Destroy()
}

func (c *CSprite) RenderCentered(renderer *sdl.Renderer, x int32, y int32, angle float64) {
	renderer.CopyEx(c.texture, nil, &sdl.Rect{
		X: x - c.width/2 + c.offsetX,
		Y: y - c.height/2 + c.offsetY,
		W: c.width,
		H: c.height,
	}, angle, &sdl.Point{X: c.width / 2, Y: c.height / 2}, sdl.FLIP_NONE)

}
