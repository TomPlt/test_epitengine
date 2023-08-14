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

// Movement functions // TODO: Move to a separate file

func (p *Player) MoveUp() {
	if p.CharType == PlayerType {
		p.Y -= Speed
	}
}

func (p *Player) MoveDown() {
	if p.CharType == PlayerType {
		p.Y += Speed
	}
}

func (p *Player) MoveLeft() {
	if p.CharType == PlayerType {
		p.X -= Speed
	}
}

func (p *Player) MoveRight() {
	if p.CharType == PlayerType {
		p.X += Speed
	}
}

func (p *Player) ResetPosition(screenWidth, screenHeight int) {
	halfWidth, halfHeight := float64(screenWidth)/2, float64(screenHeight)/2
	p.X = halfWidth
	p.Y = halfHeight
}
