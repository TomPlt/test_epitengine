package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/TomPlt/test_epitengine/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 640
	screenHeight = 480
	rectSize     = 20
	speed        = 10
)

type Game struct {
	players []player.Player
}

func (g *Game) Update() error {
	// Handle arrow keY  inputs
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game closed")
	}
	for i := range g.players {

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.players[i].ResetPosition(screenWidth, screenHeight, rectSize)
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) && g.players[i].Y > 0 {
			g.players[i].MoveUp()
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) && g.players[i].Y < screenHeight-rectSize {
			g.players[i].MoveDown()
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.players[i].X > 0 {
			g.players[i].MoveLeft()
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) && g.players[i].X < screenWidth-rectSize {
			g.players[i].MoveRight()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Set the background color to white
	screen.Fill(color.Black)

	// Create a new image of the size of the rectangle
	rectImage := ebiten.NewImage(rectSize, rectSize)
	rectColor := color.RGBA{0, 255, 0, 120} // Red color
	rectImage.Fill(rectColor)

	// Draw the rectangle image on the screen at the desired position
	for _, p := range g.players {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.X, p.Y)
		screen.DrawImage(rectImage, op)
		positionText := fmt.Sprintf("X: %v, Y: %v", p.X, p.Y)
		text.Draw(screen, positionText, basicfont.Face7x13, 15, 20, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	initialPlayers := []player.Player{
		{X: screenWidth/2 - rectSize/2, Y: screenHeight/2 - rectSize/2},
		{X: screenWidth/2 + rectSize, Y: screenHeight/2 + rectSize},
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Move the Rectangle!")
	if err := ebiten.RunGame(&Game{players: initialPlayers}); err != nil {
		log.Fatal(err)
	}
}
