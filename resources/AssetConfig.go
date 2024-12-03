package resources

import (
	"encoding/json"
	"log"
	"os"
)

type Texture struct {
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	Size     []float32 `json:"size"`
	Position []float32 `json:"position"`
}

type AssetConfig struct {
	Textures []Texture `json:"Textures"`
	// map texture name properties to texture, for fast retrieval
	TextureConfigMap map[string]Texture
}

func loadAssetsFromConfigFile(path string) AssetConfig {
	configPath := path
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	conf := &AssetConfig{TextureConfigMap: make(map[string]Texture)}
	err = json.NewDecoder(file).Decode(conf)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range conf.Textures {
		conf.TextureConfigMap[v.Name] = v
	}

	return *conf
}
