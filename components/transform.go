package components

import (
	"github.com/jakecoffman/cp"
)

// Transform component
type Transform struct {
	Position cp.Vector
	Velocity cp.Vector

	Angle float64
}

func (c Transform) ID() int {
	return TransformComponentId
}

func (c *Transform) WithPositionX(PositionX float64) *Transform {
	c.Position.X = PositionX
	return c
}

func (c *Transform) WithPostionY(PositionY float64) *Transform {
	c.Position.Y = PositionY
	return c
}

func NewTransform() *Transform {
	return &Transform{}
}
