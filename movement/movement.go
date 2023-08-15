package movement

import (
	"github.com/TomPlt/test_epitengine/enemy"
	"github.com/TomPlt/test_epitengine/player"
)

func CollidesWith(e *enemy.Enemy, p *player.Player) (float64, float64, bool) {
	if p.X < e.X+float64(e.EnemyImage.Bounds().Dx()) &&
		p.X+float64(p.PlayerImage.Bounds().Dx()) > e.X &&
		p.Y < e.Y+float64(e.EnemyImage.Bounds().Dy()) &&
		p.Y+float64(p.PlayerImage.Bounds().Dy()) > e.Y {
		return p.X, p.Y, true
	}
	return -1.0, -1.0, false
}

const CollisionThreshold = 2 // This is the delay you want to introduce before the pushback

// Movement functions // TODO: Move to a separate file
func MoveUp(p *player.Player, e *enemy.Enemy) {
	if p.CharType == player.PlayerType && p.Health > 0 {
		p.Y -= player.Speed
		if _, _, collides := CollidesWith(e, p); collides {
			p.CollisionCounter++
			if p.CollisionCounter > CollisionThreshold {
				p.Y += player.Speed
				p.CollisionCounter = 0 // Reset the counter
			}
		} else {
			p.CollisionCounter = 0 // Reset the counter if there's no collision
		}
	}
}
func MoveDown(p *player.Player, e *enemy.Enemy) {
	if p.CharType == player.PlayerType && p.Health > 0 {
		p.Y += player.Speed
		if _, _, collides := CollidesWith(e, p); collides {
			p.CollisionCounter++
			if p.CollisionCounter > CollisionThreshold {
				p.Y -= player.Speed
				p.CollisionCounter = 0 // Reset the counter
			}
		} else {
			p.CollisionCounter = 0 // Reset the counter if there's no collision
		}
	}
}

func MoveLeft(p *player.Player, e *enemy.Enemy) {
	if p.CharType == player.PlayerType && p.Health > 0 {
		p.X -= player.Speed
		if _, _, collides := CollidesWith(e, p); collides {
			p.CollisionCounter++
			if p.CollisionCounter > CollisionThreshold {
				p.X += player.Speed
				p.CollisionCounter = 0 // Reset the counter
			}
		} else {
			p.CollisionCounter = 0 // Reset the counter if there's no collision
		}
	}
}

func MoveRight(p *player.Player, e *enemy.Enemy) {
	if p.CharType == player.PlayerType && p.Health > 0 {
		p.X += player.Speed
		if _, _, collides := CollidesWith(e, p); collides {
			p.CollisionCounter++
			if p.CollisionCounter > CollisionThreshold {
				p.X -= player.Speed
				p.CollisionCounter = 0 // Reset the counter
			}
		} else {
			p.CollisionCounter = 0 // Reset the counter if there's no collision
		}
	}
}

func ResetPosition(p *player.Player, screenWidth int, screenHeight int) {
	halfWidth, halfHeight := float64(screenWidth)/2, float64(screenHeight)/2
	p.X = halfWidth
	p.Y = halfHeight
}

// Movement based on Player Position
func MoveToPlayer(e *enemy.Enemy, p *player.Player) {
	originalX, originalY := e.X, e.Y // store the original position

	if e.X < p.X {
		e.X += 1
	}
	if e.X > p.X {
		e.X -= 1
	}
	if e.Y < p.Y {
		e.Y += 1
	}
	if e.Y > p.Y {
		e.Y -= 1
	}

	// Check for collision
	_, _, collides := CollidesWith(e, p)
	if collides {
		// Reset position to avoid collision
		e.X = originalX
		e.Y = originalY
	}
}
