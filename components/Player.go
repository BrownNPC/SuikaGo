package components

type Player struct {
	CanDropFruit        bool
	FramesSinceLastDrop int
	MoveSpeed           float64
	Friction            float64
	MaxSpeed            float64
}

func (c Player) ID() int {
	return PlayerComponentId
}

func NewPlayer(MoveSpeed float64, MaxSpeed float64, Friction float64) *Player {
	return &Player{
		MoveSpeed: MoveSpeed,
		MaxSpeed:  MaxSpeed,
		Friction:  Friction,
	}
}
