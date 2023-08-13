package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/TomPlt/test_epitengine/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 640
	screenHeight = 480
	// rectSize     = 20
	speed = 10
)

type Game struct {
	players     []player.Player
	gopherImage *ebiten.Image
}

func (g *Game) init() {
	// Load the image for gopher during initialization.
	var err error
	originalImage, _, err := ebitenutil.NewImageFromFile("images/gopher.png")
	g.gopherImage = resizeImage(originalImage, 70, 50)
	if err != nil {
		log.Fatalf("Failed to load the gopher image: %v", err)
	}
}

func (g *Game) Update() error {
	// Handle arrow keY  inputs
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game closed")
	}
	for i := range g.players {

		// if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// 	g.players[i].ResetPosition(screenWidth, screenHeight)
		// }

		if ebiten.IsKeyPressed(ebiten.KeyUp) && g.players[i].Y > 0 {
			g.players[i].MoveUp()
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) && g.players[i].Y < float64(screenHeight-g.gopherImage.Bounds().Dy()) {
			g.players[i].MoveDown()
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.players[i].X > 0 {
			g.players[i].MoveLeft()
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) && g.players[i].X < float64(screenWidth-g.gopherImage.Bounds().Dx()) {
			g.players[i].MoveRight()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Set the background color to white

	// Create a new image of the size of the rectangle
	// rectImage := ebiten.NewImage(rectSize, rectSize)
	// rectColor := color.RGBA{0, 255, 0, 120} // Red color
	// rectImage.Fill(rectColor)

	// Draw the rectangle image on the screen at the desired position
	for _, p := range g.players {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.X, p.Y)
		screen.DrawImage(g.gopherImage, op)
	}
	for counter, i := range g.players {
		positionText := fmt.Sprintf("X: %v, Y: %v", i.X, i.Y)
		text.Draw(screen, positionText, basicfont.Face7x13, 10, 15*(1+counter), color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func resizeImage(src *ebiten.Image, width, height int) *ebiten.Image {
	// Create a new blank image with the desired width and height
	target := ebiten.NewImage(width, height)

	// Define the scale factors
	sx := float64(width) / float64(src.Bounds().Dx())
	sy := float64(height) / float64(src.Bounds().Dy())

	// Set up the image options
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(sx, sy)

	// Draw the original image onto the target image with the scaling transformation
	target.DrawImage(src, opts)

	return target
}

func main() {
	game := &Game{}
	game.init()

	game.players = []player.Player{
		player.NewPlayer("Player", float64(screenWidth)/2, float64(screenHeight)/2, player.PlayerType),
		player.NewPlayer("Merchant", float64(screenWidth)/2, float64(screenHeight)/6, player.NPCType), // I replaced NPCType with OtherType because in the previous example, I did not introduce NPCType.
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Gopher Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
