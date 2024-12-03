package systems

import (
	"game/components"
	"game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func InputSystem(e *engine.Entity) {
	c, ok := e.GetComponent(components.InputComponentId)
	if !ok {
		return
	}
	inputComponent := c.(*components.Input)
	inputComponent.Left = rl.IsKeyDown(rl.KeyLeft)
	inputComponent.Right = rl.IsKeyDown(rl.KeyRight)
	inputComponent.Action = rl.IsKeyDown(rl.KeySpace)

}
