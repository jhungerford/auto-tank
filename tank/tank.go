package tank

import (
	"errors"
	"fmt"
	"strings"
)

type Direction int
const (
	Up = iota
	Down = iota
	Left = iota
	Right = iota
)

func (direction Direction) String() string {
	switch direction {
	case Up: return "up"
	case Down: return "down"
	case Left: return "left"
	case Right: return "right"
	default: return "unknown"
	}
}

func ParseDirection(str string) (Direction, error) {
	switch strings.ToLower(str) {
	case "up": return Up, nil
	case "down": return Down, nil
	case "left": return Left, nil
	case "right": return Right, nil
	default:
		return -1, errors.New(fmt.Sprintf("'%s' is not a valid direction", str))
	}
}

type Tank interface {
	Move(direction Direction)
}
