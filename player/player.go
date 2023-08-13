package player

type Charactertype string

const (
	Speed                    = 10
	PlayerType Charactertype = "player"
	EnemyType  Charactertype = "enemy"
	NPCType    Charactertype = "npc"
)

type Player struct {
	Name     string
	X, Y     float64
	charType Charactertype
}

func NewPlayer(name string, x float64, y float64, charType Charactertype) Player {
	return Player{
		Name:     name,
		X:        x,
		Y:        y,
		charType: charType,
	}
}

func (p *Player) MoveUp() {
	if p.charType == PlayerType {
		p.Y -= Speed
	}
}

func (p *Player) MoveDown() {
	if p.charType == PlayerType {
		p.Y += Speed
	}
}

func (p *Player) MoveLeft() {
	if p.charType == PlayerType {
		p.X -= Speed
	}
}

func (p *Player) MoveRight() {
	if p.charType == PlayerType {
		p.X += Speed
	}
}

func (p *Player) ResetPosition(screenWidth, screenHeight int) {
	halfWidth, halfHeight := float64(screenWidth)/2, float64(screenHeight)/2
	p.X = halfWidth
	p.Y = halfHeight
}

// func (p *Player) ResetPosition(screenWidth, screenHeight, rectSize int) {
// 	p.X = float64(screenWidth/2 - rectSize/2)
// 	p.Y = float64(screenHeight/2 - rectSize/2)
// }
