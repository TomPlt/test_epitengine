package player

type Player struct {
	name string
	X, Y float64
}

func NewPlayer(name string, x float64, y float64) Player {
	return Player{
		name: name,
		X:    x,
		Y:    y,
	}
}

const (
	Speed = 10
)

func (p *Player) MoveUp() {
	p.Y -= Speed
}

func (p *Player) MoveDown() {
	p.Y += Speed
}

func (p *Player) MoveLeft() {
	p.X -= Speed
}

func (p *Player) MoveRight() {
	p.X += Speed
}

// func (p *Player) ResetPosition(screenWidth, screenHeight, rectSize int) {
// 	p.X = float64(screenWidth/2 - rectSize/2)
// 	p.Y = float64(screenHeight/2 - rectSize/2)
// }
