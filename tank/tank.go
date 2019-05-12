package tank

import (
	"errors"
	"fmt"
	"strings"
)

type Tank interface {
	Move(direction TankDirection)
}

type TankDirection int
const (
	Up = iota
	Down = iota
	Left = iota
	Right = iota
	Stop = iota
)

func (direction TankDirection) String() string {
	switch direction {
	case Up: return "up"
	case Down: return "down"
	case Left: return "left"
	case Right: return "right"
	case Stop: return "stop"
	default: return "unknown"
	}
}

func ParseTankDirection(str string) (TankDirection, error) {
	switch strings.ToLower(str) {
	case "up": return Up, nil
	case "down": return Down, nil
	case "left": return Left, nil
	case "right": return Right, nil
	case "stop": return Stop, nil
	default:
		return -1, errors.New(fmt.Sprintf("'%s' is not a valid direction", str))
	}
}
