package main

import (
	"math/rand"

	"github.com/jakecoffman/cp/v2"
)

func (s *SceneMain) sFruitSpawnerTick() {
	s.LastFruitSpawned++

}

func (s *SceneMain) sSpawnFruit() {

	if s.CanSpawnFruit && s.LastFruitSpawned >= s.MagicNums.FruitSpawnDelayFrames { // if 30 frames have passed since last fruit spawned
		s.LastFruitSpawned = 0
		s.CanSpawnFruit = false

		cloud_ENT := s.EM.GetEntitiesByTag("cloud")[0]
		current_fruit_ENT := s.EM.GetEntitiesByTag("fruit_sprites")[s.CurrentFruit]

		fruit := s.EM.CreateEntity("fruits")
		// reference the loaded sprite
		fruit.CSprite = current_fruit_ENT.CSprite
		fruit.points = current_fruit_ENT.points
		fruit.radius = current_fruit_ENT.radius
		fruit.FruitId = s.CurrentFruit
		// set position inside of cloud's "handle"
		fruit.Vec2 = cp.Vector{X: float64(s.MagicNums.CurrentFruitOffsetX + int32(cloud_ENT.Vec2.X)),
			Y: float64(cloud_ENT.CSprite.height + int32(cloud_ENT.Vec2.Y))}
		// create body, use radius from loaded sprite
		fruit.Body = cp.NewBody(100.0, cp.MomentForCircle(100, current_fruit_ENT.radius, 0, cp.Vector{}))
		fruit.Body.SetPosition(fruit.Vec2)

		fruit.Shape = cp.NewCircle(fruit.Body, current_fruit_ENT.radius, cp.Vector{})
		fruit.Shape.SetFriction(s.MagicNums.Physics.Fruit_friction)
		fruit.Shape.SetElasticity(s.MagicNums.Physics.Elasticity)
		fruit.Shape.UserData = []int{fruit.FruitId, fruit.id}
		fruit.Shape.SetCollisionType(FruitCollisionID)
		s.CurrentFruit = s.NextFruit
		s.NextFruit = rand.Intn(5)
	}
}

func (s *SceneMain) sMovement() {

	// Move Cloud
	var cloud_ENT *Entity = s.EM.GetEntitiesByTag("cloud")[0]
	var line_ENT *Entity = s.EM.GetEntitiesByTag("line")[0]
	if cloud_ENT.CInput.Left {
		if cloud_ENT.Vec2.X > s.MagicNums.MovementWallLeft {
			cloud_ENT.Vec2.X -= s.MagicNums.CloudSpeed
		}

	}
	if cloud_ENT.CInput.Right {
		if cloud_ENT.Vec2.X < s.MagicNums.MovementWallRight {

			cloud_ENT.Vec2.X += s.MagicNums.CloudSpeed
		}
	}
	line_ENT.CRectangle.X = int32(cloud_ENT.Vec2.X)
	line_ENT.CRectangle.Y = int32(cloud_ENT.Vec2.Y)
	cloud_ENT.Vec2.Y = s.MagicNums.CloudHeight
}

func (s *SceneMain) sPhysics() {

	s.EM.Space().Step(1.0 / 44.0)
	fruits := s.EM.GetEntitiesByTag("fruits")
	for _, e := range fruits {
		if !e.active {
			continue
		}
		e.Vec2 = e.Shape.Body().Position()
		e.Body.SetAngle(e.Body.Angle() + e.Body.AngularVelocity())
	}
	// fmt.Println(s.EM.GetEntitiesByTag("walls")[0].Shape.Body().Position())
	// fmt.Println(s.EM.GetEntitiesByTag("cloud")[0].Vec2)
}

func (s *SceneMain) sCollisions(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {

	shape1, shape2 := arb.Shapes()
	if shape1.Body().GetType() == cp.BODY_STATIC || shape2.Body().GetType() == cp.BODY_STATIC {
		s.CanSpawnFruit = true
		return true
	}
	shape1_e := shape1.UserData.([]int)
	shape2_e := shape2.UserData.([]int)
	same := shape1_e[0] == shape2_e[0] // index 0 is the fruit id
	// that corresponds to a fruit like (strawberry, apple)
	s.CanSpawnFruit = true
	if same {
		// index 2  is the entity id
		s.EM.KillByID("fruits", shape1_e[1])
		s.EM.KillByID("fruits", shape2_e[1])
		return false
	}
	return true
}
