package movement

import (
	"github.com/TomPlt/test_epitengine/enemy"
	"github.com/TomPlt/test_epitengine/player"
)

const CollisionDepth = 1

func CollidesWith(e *enemy.Enemy, p *player.Player) bool {
	return p.X < e.X+float64(e.EnemyImage.Bounds().Dx()) &&
		p.X+float64(p.PlayerImage.Bounds().Dx()) > e.X &&
		p.Y < e.Y+float64(e.EnemyImage.Bounds().Dy()) &&
		p.Y+float64(p.PlayerImage.Bounds().Dy()) > e.Y
}

func OverlapDepth(e *enemy.Enemy, p *player.Player) (float64, float64) {
	xOverlap := Min(p.X+float64(p.PlayerImage.Bounds().Dx()), e.X+float64(e.EnemyImage.Bounds().Dx())) - Max(p.X, e.X)
	yOverlap := Min(p.Y+float64(p.PlayerImage.Bounds().Dy()), e.Y+float64(e.EnemyImage.Bounds().Dy())) - Max(p.Y, e.Y)

	return xOverlap, yOverlap
}

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func MovePlayer(p *player.Player, enemies []enemy.Enemy, dir Direction) {
	if p.CharType != player.PlayerType || p.Health <= 0 {
		return
	}

	originalX, originalY := p.X, p.Y

	var moveAmount float64
	switch dir {
	case Up:
		moveAmount = -player.Speed
		p.Y += moveAmount
	case Down:
		moveAmount = player.Speed
		p.Y += moveAmount
	case Left:
		moveAmount = -player.Speed
		p.X += moveAmount
	case Right:
		moveAmount = player.Speed
		p.X += moveAmount
	}

	for _, e := range enemies {
		if CollidesWith(&e, p) {
			xOverlap, yOverlap := OverlapDepth(&e, p)

			switch dir {
			case Up:
				if yOverlap > CollisionDepth {
					p.Y = originalY
				}
			case Down:
				if yOverlap > CollisionDepth {
					p.Y = originalY
				}
			case Left:
				if xOverlap > CollisionDepth {
					p.X = originalX
				}
			case Right:
				if xOverlap > CollisionDepth {
					p.X = originalX
				}
			}
		}
	}
}

func ResetPosition(p *player.Player, screenWidth int, screenHeight int) {
	halfWidth, halfHeight := float64(screenWidth)/2, float64(screenHeight)/2
	p.X = halfWidth
	p.Y = halfHeight
}

func MoveToPlayer(e *enemy.Enemy, p *player.Player) {
	if e.Entype == enemy.Enemy1 {

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

		if CollidesWith(e, p) {
			xOverlap, yOverlap := OverlapDepth(e, p)

			// Adjust enemy's position based on overlap depth, only if it's greater than CollisionDepth.
			if xOverlap > CollisionDepth {
				if e.X < p.X {
					e.X -= xOverlap
				} else {
					e.X += xOverlap
				}
			}

			if yOverlap > CollisionDepth {
				if e.Y < p.Y {
					e.Y -= yOverlap
				} else {
					e.Y += yOverlap
				}
			}
		}
	}
}
