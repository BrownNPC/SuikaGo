package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene interface {
	Update(*Game)
	sRender(*Game)
	LoadAssets(*Game)
	UnloadAssets(*Game)
}

type Game struct {
	Scenes map[string]Scene

	Running      bool
	CurrentScene string
	Window       *sdl.Window
	Renderer     *sdl.Renderer

	Assets     *GameAssets
	FpsManager *gfx.FPSmanager

	Debug      bool
	StartTime  uint64
	FrameDelay float64
	FPS        string
}

func NewGame() *Game {
	g := &Game{
		Running: true,
		Scenes:  make(map[string]Scene),

		StartTime: sdl.GetTicks64(),

		CurrentScene: "main",

		FpsManager: &gfx.FPSmanager{},

		Assets: LoadJsonConfig(),
	}
	g.Window, _ = sdl.CreateWindow("Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, g.Assets.Config.WindowWidth, g.Assets.Config.WindowHeight, sdl.WINDOW_SHOWN)
	g.Renderer, _ = sdl.CreateRenderer(g.Window, -1, sdl.RENDERER_ACCELERATED)
	g.AddScene("main", InitMainScene(g))
	g.SetCurrentScene("main")
	g.Debug = g.Assets.Config.DebugDraw
	gfx.InitFramerate(g.FpsManager)
	gfx.SetFramerate(g.FpsManager, g.Assets.Config.FPS)

	return g
}

func (g *Game) AddScene(name string, scene Scene) {
	g.Scenes[name] = scene
}

func (g *Game) SetCurrentScene(name string) {

	if g.CurrentScene != "" {
		g.Scenes[g.CurrentScene].UnloadAssets(g)
	}
	g.CurrentScene = name

	g.Scenes[g.CurrentScene].LoadAssets(g)
}

func (g *Game) DebugUpdate(delayCounter *uint64, updateDelayms uint64) {
	g.FrameDelay = (float64(gfx.FramerateDelay(g.FpsManager)) * 0.001)

	if sdl.GetTicks64()-*delayCounter > updateDelayms {
		// update FPS counter
		g.FPS = fmt.Sprintf("%.0f", 1/g.FrameDelay)
		*delayCounter = sdl.GetTicks64()

	}
}

func (g *Game) DebugDraw() {

	gfx.StringRGBA(g.Renderer, 0, 0, fmt.Sprint("FPS: ", g.FPS), 20, 255, 20, 255)
}

func (g *Game) Run() {

	var DebugUpdatedelayCounter uint64 // Keep track of FPS updates in milliseconds
	for g.Running {
		// used to get g.FPS and g.FrameDelay
		g.DebugUpdate(&DebugUpdatedelayCounter, 1000)

		g.Scenes[g.CurrentScene].Update(g)

		g.Renderer.SetDrawColor(0, 0, 0, 255)
		g.Renderer.Clear()
		g.Scenes[g.CurrentScene].sRender(g)
		if g.Debug {
			g.DebugDraw()
		}
		g.Renderer.Present()
	}
	sdl.Quit()

}
