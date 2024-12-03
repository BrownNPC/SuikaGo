package systems

import (
	"game/components"
	"game/engine"
)

// update the sprite using Transform component
func UpdateSpriteSystem(e *engine.Entity) {
	comp, ok := e.GetComponent(components.TransformComponentId)
	if !ok {
		return
	}
	transform := comp.(*components.Transform)

	comp, ok = e.GetComponent(components.SpriteComponentId)
	if !ok {
		return
	}
	sprite := comp.(*components.Sprite)

	sprite.Dest.X = float32(transform.Position.X)
	sprite.Dest.Y = float32(transform.Position.Y)
}
