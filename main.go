package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/TomPlt/test_epitengine/enemy"
	"github.com/TomPlt/test_epitengine/movement"
	"github.com/TomPlt/test_epitengine/player"
	"github.com/hajimehoshi/ebiten/v2"
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
	players           []player.Player
	enemies           []enemy.Enemy
	collisionDetected bool
	collisionNotified bool
	maxCollisionTime  int
	collisionTimer    int
	isGameOver        bool
	// gopherImage *ebiten.Image
}

// func (g *Game) init() {
// 	// Load the image for gopher during initialization.

// }

func (g *Game) Update() error {
	// Handle arrow keY  inputs
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game closed")
	}
	for i := range g.players {

		// if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// 	g.players[i].ResetPosition(screenWidth, screenHeight)
		// }
		// check if there in an enmy in the same position as the player with bounding box
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			movement.MovePlayer(&g.players[i], g.enemies, movement.Up)
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			movement.MovePlayer(&g.players[i], g.enemies, movement.Down)
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			movement.MovePlayer(&g.players[i], g.enemies, movement.Left)
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			movement.MovePlayer(&g.players[i], g.enemies, movement.Right)
		}

		for j := range g.enemies {
			movement.MoveToPlayer(&g.enemies[j], &g.players[i])
		}

		collides := movement.CollidesWith(&g.enemies[0], &g.players[i])
		if collides && !g.collisionNotified {
			g.collisionDetected = true
			g.players[i].Health -= 10
			g.collisionTimer = 0 // Reset the timer when collision occurs
		} else if !collides {
			g.collisionTimer++ // Increment the timer if no collision
			if g.collisionTimer > g.maxCollisionTime {
				g.collisionDetected = false
				g.collisionNotified = false
			}
		}
		if g.players[i].Name == "Player" && g.players[i].Health <= 0 {
			g.isGameOver = true
		}
	}
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	for counter, p := range g.players {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.X, p.Y)
		screen.DrawImage(p.PlayerImage, op)

		positionText := fmt.Sprintf("%v, X: %v, Y: %v", p.Name, math.Round(p.X), math.Round(p.Y))
		text.Draw(screen, positionText, basicfont.Face7x13, 10, 15*(1+counter), color.White)
		text.Draw(screen, fmt.Sprintf("Health: %v", p.Health), basicfont.Face7x13, 10, 15*(1+counter)+15, color.White)
	}

	if g.collisionDetected && g.collisionTimer <= g.maxCollisionTime {
		msg := "Collision!"
		x := screenWidth/2 - len(msg)*7/2 // Calculate the x position to center the text
		y := screenHeight / 2             // Y position is at the center
		text.Draw(screen, msg, basicfont.Face7x13, x, y, color.White)
		g.collisionNotified = true
	}

	for counter, e := range g.enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.X, e.Y)
		screen.DrawImage(e.EnemyImage, op)
		positionText := fmt.Sprintf("%v X: %v, Y: %v", e.Name, math.Round(e.X), math.Round(e.Y))
		text.Draw(screen, positionText, basicfont.Face7x13, screenWidth-150, 15*(1+counter), color.White)
	}
	if g.isGameOver {
		g.drawGameOver(screen)
	}
}
func (g *Game) drawGameOver(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.NRGBA{0, 0, 0, 200}) // RGBA with 50% transparency
	screen.DrawImage(overlay, &ebiten.DrawImageOptions{})
	text.Draw(screen, "Game Over", basicfont.Face7x13, screenWidth/2-40, screenHeight/2, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}
	// game.init()

	game.players = []player.Player{
		player.NewPlayer("Player", float64(screenWidth)/2, float64(screenHeight)/2, player.PlayerType, "images/gopher.png"),
	}
	game.enemies = []enemy.Enemy{
		enemy.NewEnemy("Enemy1", float64(screenWidth)/6, float64(screenHeight)/2, enemy.Enemy1, "images/croco.png"),
		enemy.NewEnemy("Tree", float64(screenWidth)/9, float64(screenHeight)/3, enemy.Environment, "images/tree.png"),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Gopher Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
