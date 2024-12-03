package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Dest rl.Rectangle
}

func (c *Sprite) WithDest(dest rl.Rectangle) *Sprite {
	c.Dest = dest
	return c
}

func NewSprite() *Sprite {
	return &Sprite{}
}

func (c Sprite) ID() int {
	return SpriteComponentId
}
