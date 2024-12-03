package main

import (
	"game/components"
	"game/engine"
	"game/resources"
	"game/systems"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneMain struct {
	engine.BaseScene
	*resources.AssetManager
}

func (s *SceneMain) Init() {
	s.LoadAssets()
	_, plrDest := s.AssetManager.GetTexture("cloud")
	s.EntityManager.CreateEntity("player",
		components.NewSprite().WithDest(plrDest),
		components.NewInput(),
		components.NewPlayer(0.3, 5, 0.3),
		components.NewTransform().
			WithPositionX(float64(plrDest.X)).
			WithPostionY(float64(plrDest.Y)),
	)
}
func (s *SceneMain) Render() {
	tex, dest := s.AssetManager.GetTexture("backdrop")
	s.DrawTexture(tex, dest, rl.White)
	tex, dest = s.AssetManager.GetTexture("container")
	s.DrawTextureRotateCenter(tex, dest, 0, rl.RayWhite)
	tex, _ = s.AssetManager.GetTexture("cloud")
	e := s.EntityManager.GetFirstEntityWithTag("player")
	comp, _ := e.GetComponent(components.SpriteComponentId)
	sprite := comp.(*components.Sprite)
	s.DrawTextureRotateCenter(tex, sprite.Dest, 0, rl.ColorAlpha(rl.White, 1.0))
}
func (s *SceneMain) Update(virtualWidth float32, virtualHeight float32) {
	s.UpdateBaseScene(virtualWidth, virtualHeight)
	s.ForEachEntity(
		func(e *engine.Entity) {
			systems.InputSystem(e)
			systems.MovementSystem(e)
			systems.UpdateSpriteSystem(e)
		},
	)
}

func (s *SceneMain) NextScene() string {
	return "menu"
}
func (s *SceneMain) Unload() {
	s.AssetManager.Unload()
}

func (s *SceneMain) LoadAssets() {
	configPath := "TextureConfigSceneMain.json"
	s.AssetManager = resources.NewAssetManager(configPath)

}
