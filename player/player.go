package player

import (
	"log"

	"github.com/TomPlt/test_epitengine/viz_utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Charactertype string

const (
	Speed                    = 10
	PlayerType Charactertype = "player"
	NPCType    Charactertype = "npc"
)

type Player struct {
	Name        string
	X, Y        float64
	CharType    Charactertype
	PlayerImage *ebiten.Image
}

func NewPlayer(name string, x float64, y float64, CharType Charactertype, image string) Player {
	originalImage, _, err := ebitenutil.NewImageFromFile(image)
	if err != nil {
		log.Fatalf("Failed to load the image: %v", err)
	}

	if originalImage == nil {
		log.Fatal("The original image is nil.")
	}

	imageOut := viz_utils.ResizeImage(originalImage, 70, 50)

	return Player{
		Name:        name,
		X:           x,
		Y:           y,
		CharType:    CharType,
		PlayerImage: imageOut,
	}
}