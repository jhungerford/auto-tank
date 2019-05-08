package tank

type Tank struct {
	left, right Tread
	direction TankDirection
}

func (t *Tank) Direction() TankDirection {
	return t.direction
}

type TankDirection int
const (
	North     = TankDirection(iota)
	NorthEast = TankDirection(iota)
	East      = TankDirection(iota)
	SouthEast = TankDirection(iota)
	South     = TankDirection(iota)
	SouthWest = TankDirection(iota)
	West      = TankDirection(iota)
	NorthWest = TankDirection(iota)
	Stop      = TankDirection(iota)
)

type TreadDirection int
const (
	Forward = TreadDirection(iota)
	Reverse = TreadDirection(iota)
	Off     = TreadDirection(iota)
)


func New() Tank {
	return Tank{
		left: Tread{
			Front: Pins{HighPin: 4, LowPin: 5, SpeedPin: 1},
			Rear: Pins{HighPin: 10, LowPin: 6, SpeedPin: 27},
		},
		right: Tread{
			Front: Pins{HighPin: 0, LowPin: 7, SpeedPin: 23},
			Rear: Pins{HighPin: 22, LowPin: 21, SpeedPin: 24},
		},
	}
}

type Tread interface {

}

type tankTreadDirection struct {
	left, right TreadDirection
}

var tankDirectionMap = map[TankDirection]tankTreadDirection{
	North :     {Forward, Forward},
	NorthEast : {Forward, Off},
	East :      {Forward, Reverse},
	SouthEast : {Off, Reverse},
	South :     {Reverse, Reverse},
	SouthWest : {Reverse, Off},
	West :      {Reverse, Forward},
	NorthWest : {Off, Forward},
	Stop :      {Off, Off},
}

type Pins struct {
	HighPin, LowPin, SpeedPin int
}

type Tread struct {
	Front, Rear Pins
}

type Treads struct {
	Left, Right Tread
}
