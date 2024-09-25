package main

import (
	"github.com/jakecoffman/cp/v2"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Entity struct {
	id      int
	tag     string
	active  bool
	radius  float64
	points  int32
	FruitId int
	// Components

	CSprite    *CSprite
	CInput     *CInput
	CRectangle *sdl.Rect
	CFillColor sdl.Color

	Vec2  cp.Vector
	Body  *cp.Body
	Shape *cp.Shape
	Font  *ttf.Font
}

func NewEntity(id int, tag string) *Entity {
	return &Entity{
		id:     id,
		tag:    tag,
		active: true,
	}
}

func (e *Entity) Kill() {
	e.active = false
}
