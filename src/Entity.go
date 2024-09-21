package main

import (
	"github.com/jakecoffman/cp/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type Entity struct {
	id     int
	tag    string
	active bool
	radius float64
	points int32
	// Components

	CSprite    *CSprite
	CInput     *CInput
	Vec2       cp.Vector
	CRectangle *sdl.Rect
	CFillColor sdl.Color

	Body  *cp.Body
	Shape *cp.Shape
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
