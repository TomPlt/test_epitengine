package player

import (
	"log"

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

func resizeImage(src *ebiten.Image, width, height int) *ebiten.Image {
	target := ebiten.NewImage(width, height)
	sx := float64(width) / float64(src.Bounds().Dx())
	sy := float64(height) / float64(src.Bounds().Dy())
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(sx, sy)
	target.DrawImage(src, opts)
	return target
}

func NewPlayer(name string, x float64, y float64, CharType Charactertype, image string) Player {
	originalImage, _, err := ebitenutil.NewImageFromFile(image)
	if err != nil {
		log.Fatalf("Failed to load the image: %v", err)
	}

	if originalImage == nil {
		log.Fatal("The original image is nil.")
	}

	imageOut := resizeImage(originalImage, 70, 50)

	return Player{
		Name:        name,
		X:           x,
		Y:           y,
		CharType:    CharType,
		PlayerImage: imageOut,
	}
}

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
