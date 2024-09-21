package main

import (
	"github.com/jakecoffman/cp/v2"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// MAIN GAME SCENE

type SceneMain struct {
	EM            *EntityManager
	Score         int
	MagicNums     *MagicNums
	CurrentFruit  int
	NextFruit     int
	CanSpawnFruit bool
	Space         *cp.Space
}

func InitMainScene(g *Game) *SceneMain {
	scene := &SceneMain{
		EM:        NewEntityManager(),
		MagicNums: LoadMagicNumsJson(),
		Space:     cp.NewSpace(),
	}

	scene.Space.SetGravity(cp.Vector{X: 0, Y: scene.MagicNums.Physics.Gravity})
	scene.Space.SetDamping(scene.MagicNums.Physics.Damping)
	return scene
}

func (s *SceneMain) LoadAssets(g *Game) {

	aSprites := g.Assets.Sprites //asset sprites
	// only background is stored inside scene manager, other things are entities.
	bg := s.EM.CreateEntity("background")
	bg.CSprite = NewCSprite(g.Renderer, aSprites.Background.FileName, aSprites.Background.Width, aSprites.Background.Height, 0, 0)
	bg.Vec2 = cp.Vector{X: 0, Y: 0}

	line := s.EM.CreateEntity("line")
	line.CRectangle = &sdl.Rect{X: 0, Y: 0, W: s.MagicNums.LineWidth, H: s.MagicNums.LineLength}

	cloud := s.EM.CreateEntity("cloud")
	cloud.CSprite = NewCSprite(g.Renderer, aSprites.Cloud.FileName, aSprites.Cloud.Width, aSprites.Cloud.Height, 0, 0)
	cloud.CInput = NewCInput()
	cloud.Vec2 = cp.Vector{X: float64(g.Assets.Config.WindowWidth) / 2, Y: s.MagicNums.CloudHeight}
	//load fruit textures into entities

	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Cherry.FileName, aSprites.Cherry.Width, aSprites.Cherry.Height, aSprites.Cherry.OffsetX, aSprites.Cherry.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Strawberry.FileName, aSprites.Strawberry.Width, aSprites.Strawberry.Height, aSprites.Strawberry.OffsetX, aSprites.Strawberry.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Grapes.FileName, aSprites.Grapes.Width, aSprites.Grapes.Height, aSprites.Grapes.OffsetX, aSprites.Grapes.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Dekopon.FileName, aSprites.Dekopon.Width, aSprites.Dekopon.Height, aSprites.Dekopon.OffsetX, aSprites.Dekopon.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Orange.FileName, aSprites.Orange.Width, aSprites.Orange.Height, aSprites.Orange.OffsetX, aSprites.Orange.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Apple.FileName, aSprites.Apple.Width, aSprites.Apple.Height, aSprites.Apple.OffsetX, aSprites.Apple.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Pear.FileName, aSprites.Pear.Width, aSprites.Pear.Height, aSprites.Pear.OffsetX, aSprites.Pear.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Peach.FileName, aSprites.Peach.Width, aSprites.Peach.Height, aSprites.Peach.OffsetX, aSprites.Peach.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Pineapple.FileName, aSprites.Pineapple.Width, aSprites.Pineapple.Height, aSprites.Pineapple.OffsetX, aSprites.Pineapple.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Melon.FileName, aSprites.Melon.Width, aSprites.Melon.Height, aSprites.Melon.OffsetX, aSprites.Melon.OffsetY)
	s.EM.CreateEntity("fruit_sprites").CSprite = NewCSprite(g.Renderer, aSprites.Watermelon.FileName, aSprites.Watermelon.Width, aSprites.Watermelon.Height, aSprites.Watermelon.OffsetX, aSprites.Watermelon.OffsetY)
	s.EM.Update()
	for i, e := range s.EM.GetEntitiesByTag("fruit_sprites") {
		e.points = s.MagicNums.Points[i]
		e.radius = s.MagicNums.FruitRadii[i]
	}

	// add container walls in physics space
	BottomWall := s.EM.CreateEntity("walls")
	BottomWall.Body = cp.NewStaticBody()
	BottomWall.Shape = cp.NewSegment(BottomWall.Body, s.MagicNums.Wall.BottomLeft(), s.MagicNums.Wall.BottomRight(), 2.0)
	s.Space.AddBody(BottomWall.Body)
	s.Space.AddShape(BottomWall.Shape)
	BottomWall.Shape.SetFriction(s.MagicNums.Physics.Wall_friction)

}
func (s *SceneMain) UnloadAssets(g *Game) {
}

func (s *SceneMain) sRender(g *Game) {

	s.EM.GetEntities()

	for _, e := range s.EM.GetEntities() {
		if !e.active {
			continue
		}
		if e.tag == "line" {
			gfx.ThickLineColor(
				g.Renderer,
				e.CRectangle.X+s.MagicNums.LineOffsetX, e.CRectangle.Y+s.MagicNums.LineOffsetY,
				e.CRectangle.X+s.MagicNums.LineOffsetX, e.CRectangle.Y+s.MagicNums.LineOffsetY+s.MagicNums.LineLength,
				s.MagicNums.LineWidth,
				sdl.Color{R: 255, G: 255, B: 255, A: 255},
			)
			continue
		}
		if e.CSprite != nil && e.tag != "fruit_sprites" && e.tag != "fruits" {
			e.CSprite.Render(g.Renderer, int32(e.Vec2.X), int32(e.Vec2.Y), 0)
		}

		if e.CSprite != nil && e.tag == "fruits" {
			e.CSprite.RenderCentered(g.Renderer, int32(e.Vec2.X), int32(e.Vec2.Y), 0)
		}

		if e.tag == "cloud" {
			// cloud holding fruit
			s.EM.GetEntitiesByTag("fruit_sprites")[s.CurrentFruit].CSprite.RenderCentered(g.Renderer,
				s.MagicNums.CurrentFruitOffsetX+int32(e.Vec2.X),
				e.CSprite.height+int32(e.Vec2.Y), 0)
		}
	}
	fruits := s.EM.GetEntitiesByTag("fruit_sprites")
	// show next fruit
	fruits[s.NextFruit].CSprite.RenderCentered(g.Renderer, g.Assets.Config.WindowWidth-s.MagicNums.NextFruitOffsetX, s.MagicNums.NextFruitOffsetY, 0)

}

func (s *SceneMain) sInput(g *Game) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			g.Running = false
		case *sdl.KeyboardEvent:
			if e.Keysym.Sym == sdl.K_ESCAPE {
				g.Running = false
			}
			if e.Type == sdl.KEYDOWN {
				var cloud_ENT *Entity = s.EM.GetEntitiesByTag("cloud")[0]
				switch e.Keysym.Sym {
				case sdl.K_a:
					cloud_ENT.CInput.Left = true
				case sdl.K_LEFT:
					cloud_ENT.CInput.Left = true
				case sdl.K_d:
					cloud_ENT.CInput.Right = true
				case sdl.K_RIGHT:
					cloud_ENT.CInput.Right = true
				case sdl.K_9:
					s.MagicNums = LoadMagicNumsJson()
				case sdl.K_SPACE:
					s.sSpawnFruit()
				}
			} else if e.Type == sdl.KEYUP {
				var cloud_ENT *Entity = s.EM.GetEntitiesByTag("cloud")[0]
				switch e.Keysym.Sym {
				case sdl.K_a:
					cloud_ENT.CInput.Left = false
				case sdl.K_LEFT:
					cloud_ENT.CInput.Left = false
				case sdl.K_d:
					cloud_ENT.CInput.Right = false
				case sdl.K_RIGHT:
					cloud_ENT.CInput.Right = false
				}
			}

		}
	}
}

func (s *SceneMain) Update(g *Game) {
	s.EM.Update()
	s.sInput(g)
	s.sMovement()
	s.sPhysics()
	// s.sSpawnFruit(g)

}
