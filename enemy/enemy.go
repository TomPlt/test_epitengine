package enemy

import (
	"log"

	"github.com/TomPlt/test_epitengine/viz_utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemytype string

const (
	Enemy1 Enemytype = "enemy1"
)

type Enemy struct {
	Name        string
	X, Y        float64
	Entype      Enemytype
	PlayerImage *ebiten.Image
}

func NewEnemy(name string, x float64, y float64, Entype Enemytype, image string) Enemy {
	originalImage, _, err := ebitenutil.NewImageFromFile(image)
	if err != nil {
		log.Fatalf("Failed to load the image: %v", err)
	}

	if originalImage == nil {
		log.Fatal("The original image is nil.")
	}

	imageOut := viz_utils.ResizeImage(originalImage, 70, 50)

	return Enemy{
		Name:        name,
		X:           x,
		Y:           y,
		Entype:      Entype,
		PlayerImage: imageOut,
	}
}

// Movement based on Player Position
func (e *Enemy) MoveToPlayer(playerX float64, playerY float64) {
	if e.X < playerX {
		e.X += 1
	}
	if e.X > playerX {
		e.X -= 1
	}
	if e.Y < playerY {
		e.Y += 1
	}
	if e.Y > playerY {
		e.Y -= 1
	}
}
