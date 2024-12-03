package resources

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AssetManager struct {
	config   AssetConfig
	Fonts    map[string]rl.Font
	Textures map[string]rl.Texture2D
	Sounds   map[string]rl.Sound
	Music    map[string]rl.Music
}

func NewAssetManager(assetConfigFilePath string) *AssetManager {
	am := &AssetManager{
		Fonts:    make(map[string]rl.Font),
		Textures: make(map[string]rl.Texture2D),
		Sounds:   make(map[string]rl.Sound),
		Music:    make(map[string]rl.Music),
	}
	am.config = loadAssetsFromConfigFile(assetConfigFilePath)
	for _, v := range am.config.Textures {
		tex := rl.LoadTexture(v.Path)
		if rl.IsTextureReady(tex) {
			am.AddTexture(v.Name, tex)
		}
	}
	return am
}

func (a *AssetManager) GetFont(name string) rl.Font {
	return a.Fonts[name]
}

// return texture defined in config, and a dest rl.Rectangle
func (a *AssetManager) GetTexture(name string) (Texture rl.Texture2D, dest rl.Rectangle) {
	dest = rl.Rectangle{}

	v, exists := a.config.TextureConfigMap[name]
	if !exists {
		log.Panicln("ERROR texture with the name ", name, " does not exist in the texture config file")
	}
	dest.X = v.Position[0]
	dest.Y = v.Position[1]
	dest.Width = v.Size[0]
	dest.Height = v.Size[1]
	return a.Textures[name], dest
}

func (a *AssetManager) GetDest(name string) rl.Rectangle {
	dest := rl.Rectangle{}

	v, exists := a.config.TextureConfigMap[name]
	if !exists {
		log.Panicln("ERROR texture with the name ", name, " does not exist in the texture config file")
	}
	dest.X = v.Position[0]
	dest.Y = v.Position[1]
	dest.Width = v.Size[0]
	dest.Height = v.Size[1]
	return dest
}

func (a *AssetManager) GetSound(name string) rl.Sound {
	return a.Sounds[name]
}

func (a *AssetManager) GetMusic(name string) rl.Music {
	return a.Music[name]
}

func (a *AssetManager) AddFont(name string, font rl.Font) {
	a.Fonts[name] = font
}

func (a *AssetManager) AddTexture(name string, texture rl.Texture2D) {
	a.Textures[name] = texture
}

func (a *AssetManager) AddSound(name string, sound rl.Sound) {
	a.Sounds[name] = sound
}

func (a *AssetManager) AddMusic(name string, music rl.Music) {
	a.Music[name] = music
}

func (a *AssetManager) Unload() {
	for _, v := range a.Fonts {
		rl.UnloadFont(v)
	}
	for _, v := range a.Textures {
		rl.UnloadTexture(v)
	}
	for _, v := range a.Sounds {
		rl.UnloadSound(v)
	}
	for _, v := range a.Music {
		rl.UnloadMusicStream(v)

	}
}
