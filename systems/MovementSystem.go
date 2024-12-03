package systems

import (
	"game/components"
	"game/engine"
)

// updates the player's velocity and sets its "CanDropFruit" according to input component

func MovementSystem(e *engine.Entity) {
	comp, ok := e.GetComponent(components.InputComponentId)
	if !ok {
		return
	}
	input := comp.(*components.Input)

	comp, ok = e.GetComponent(components.PlayerComponentId)
	if !ok {
		return
	}
	plr := comp.(*components.Player)

	comp, ok = e.GetComponent(components.TransformComponentId)
	if !ok {
		return
	}
	transform := comp.(*components.Transform)

	// Apply movement based on input
	if input.Left {
		transform.Velocity.X -= plr.MoveSpeed
	} else if input.Right {
		transform.Velocity.X += plr.MoveSpeed
	} else {
		// Apply friction when no input is given
		if transform.Velocity.X > 0 {
			transform.Velocity.X -= plr.Friction
			if transform.Velocity.X < 0 { // Prevent overshooting
				transform.Velocity.X = 0
			}
		} else if transform.Velocity.X < 0 {
			transform.Velocity.X += plr.Friction
			if transform.Velocity.X > 0 { // Prevent overshooting
				transform.Velocity.X = 0
			}
		}
	}

	// Clamp velocity to avoid excessive speed
	if transform.Velocity.X > plr.MaxSpeed {
		transform.Velocity.X = plr.MaxSpeed
	} else if transform.Velocity.X < -plr.MaxSpeed {
		transform.Velocity.X = -plr.MaxSpeed
	}

	transform.Position.X += transform.Velocity.X
	transform.Position.Y += transform.Velocity.Y
	plr.CanDropFruit = input.Action
}
