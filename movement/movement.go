package movement

import (
	"github.com/TomPlt/test_epitengine/enemy"
	"github.com/TomPlt/test_epitengine/player"
)

func CollidesWith(e *enemy.Enemy, p *player.Player) (float64, float64) {
	if p.X < e.X+float64(e.EnemyImage.Bounds().Dx()) && p.X+float64(p.PlayerImage.Bounds().Dx()) > e.X &&
		p.Y < e.Y+float64(e.EnemyImage.Bounds().Dy()) &&
		p.Y+float64(p.PlayerImage.Bounds().Dy()) > e.Y {
		return p.X, p.Y
	}
	return -1.0, -1.0
}

// Movement functions // TODO: Move to a separate file
func MoveUp(p *player.Player) {
	if p.CharType == player.PlayerType {
		p.Y -= player.Speed
	}
}

func MoveDown(p *player.Player) {
	if p.CharType == player.PlayerType {
		p.Y += player.Speed
	}
}

func MoveLeft(p *player.Player) {
	if p.CharType == player.PlayerType {
		p.X -= player.Speed
	}
}

func MoveRight(p *player.Player) {
	if p.CharType == player.PlayerType {
		p.X += player.Speed
	}
}

func ResetPosition(p *player.Player, screenWidth int, screenHeight int) {
	halfWidth, halfHeight := float64(screenWidth)/2, float64(screenHeight)/2
	p.X = halfWidth
	p.Y = halfHeight
}

// Movement based on Player Position
func MoveToPlayer(e *enemy.Enemy, playerX float64, playerY float64) {
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
