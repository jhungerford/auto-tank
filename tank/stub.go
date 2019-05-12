package tank

// StubTank remembers the movements it's made between resets - useful for testing.
type StubTank struct {
	Moves []TankDirection
}

func (tank *StubTank) Move(direction TankDirection) {
	tank.Moves = append(tank.Moves, direction)
}

func (tank *StubTank) Reset() {
	tank.Moves = nil
}
