package tank

type Tank interface {
	Move(direction string)
}

func MoveTank(tank Tank, direction string) {
	tank.Move(direction)
}