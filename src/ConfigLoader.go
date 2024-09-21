package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jakecoffman/cp/v2"
)

// Define structs that match the JSON structure
type Config struct {
	FPS          uint32 `json:"FPS"`
	WindowWidth  int32  `json:"WindowWidth"`
	WindowHeight int32  `json:"WindowHeight"`
	DebugDraw    bool   `json:"DebugDraw"`
}

type Sounds struct {
	BGM string `json:"BGM"`
}

type SpriteData struct {
	FileName string `json:"FileName"`
	Width    int32  `json:"Width"`
	Height   int32  `json:"Height"`
	OffsetX  int32  `json:"OffsetX"`
	OffsetY  int32  `json:"OffsetY"`
}

type Sprites struct {
	Background SpriteData `json:"Background"`
	Cloud      SpriteData `json:"Cloud"`
	Cherry     SpriteData `json:"Cherry"`
	Strawberry SpriteData `json:"Strawberry"`
	Grapes     SpriteData `json:"Grapes"`
	Dekopon    SpriteData `json:"Dekopon"`
	Orange     SpriteData `json:"Orange"`
	Apple      SpriteData `json:"Apple"`
	Pear       SpriteData `json:"Pear"`
	Peach      SpriteData `json:"Peach"`
	Pineapple  SpriteData `json:"Pineapple"`
	Melon      SpriteData `json:"Melon"`
	Watermelon SpriteData `json:"Watermelon"`
}

type GameAssets struct {
	Sprites Sprites `json:"Sprites"`
	Sounds  Sounds  `json:"Sounds"`
	Config  Config  `json:"Config"`
}

func LoadJsonData() *GameAssets {

	// Read the file contents
	jsonData, err := os.ReadFile("Assets.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal JSON data into struct
	var assets GameAssets
	err = json.Unmarshal(jsonData, &assets)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return &assets
}

type Wall struct {
	Left   float64 `json:"Left"`
	Right  float64 `json:"Right"`
	Top    float64 `json:"Top"`
	Bottom float64 `json:"Bottom"`
}

func (w *Wall) TopLeft() cp.Vector {
	return cp.Vector{X: w.Left, Y: w.Top}
}

func (w *Wall) TopRight() cp.Vector {
	return cp.Vector{X: w.Right, Y: w.Top}
}

func (w *Wall) BottomLeft() cp.Vector {
	return cp.Vector{X: w.Left, Y: w.Bottom}
}

func (w *Wall) BottomRight() cp.Vector {
	return cp.Vector{X: w.Right, Y: w.Bottom}
}

type MagicPhysics struct {
	Gravity        float64 `json:"Gravity"`
	Density        float64 `json:"density"`
	Elasticity     float64 `json:"elasticity"`
	Impulse        float64 `json:"impulse"`
	Bias           float64 `json:"bias"`
	Fruit_friction float64 `json:"fruit_friction"`
	Wall_friction  float64 `json:"wall_friction"`
	Damping        float64 `json:"damping"`
}
type MagicNums struct {
	CloudHeight float64 `json:"CloudHeight"`
	CloudSpeed  float64 `json:"CloudSpeed"`
	Wall        Wall    `json:"Wall"`
	LineLength  int32   `json:"LineLength"`
	LineWidth   int32   `json:"LineWidth"`
	LineOffsetX int32   `json:"LineOffsetX"`
	LineOffsetY int32   `json:"LineOffsetY"`

	NextFruitOffsetX    int32 `json:"NextFruitOffsetX"`
	NextFruitOffsetY    int32 `json:"NextFruitOffsetY"`
	CurrentFruitOffsetX int32 `json:"CurrentFruitOffsetX"`
	CurrentFruitOffsetY int32 `json:"CurrentFruitOffsetY"`

	Physics MagicPhysics `json:"Physics"`

	Points     []int32   `json:"Points"`
	FruitRadii []float64 `json:"FruitRadii"`
}

func LoadMagicNumsJson() *MagicNums {

	// Read the file contents
	jsonData, err := os.ReadFile("Magic.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal JSON data into struct
	var magicNums MagicNums
	err = json.Unmarshal(jsonData, &magicNums)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return &magicNums
}
