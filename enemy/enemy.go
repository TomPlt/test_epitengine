package enemy

import (
	"log"

	"github.com/TomPlt/test_epitengine/viz_utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemytype string

const (
	Enemy1      Enemytype = "enemy1"
	Environment Enemytype = "environment"
)

type Enemy struct {
	Name       string
	X, Y       float64
	Entype     Enemytype
	EnemyImage *ebiten.Image
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
		Name:       name,
		X:          x,
		Y:          y,
		Entype:     Entype,
		EnemyImage: imageOut,
	}
}
