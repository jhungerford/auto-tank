package tank

import (
	"errors"
	"fmt"
	"strings"
)

type Tread interface {
	Move(direction TreadDirection)
}

type TreadDirection int
const (
	Forward = iota
	Reverse = iota
	Off = iota
)

func (direction TreadDirection) String() string {
	switch direction {
	case Forward: return "forward"
	case Reverse: return "reverse"
	case Off: return "off"
	default: return "unknown"
	}
}

func ParseTreadDirection(str string) (TreadDirection, error) {
	switch strings.ToLower(str) {
	case "forward": return Forward, nil
	case "reverse": return Reverse, nil
	case "off": return Off, nil
	default:
		return -1, errors.New(fmt.Sprintf("'%s' is not a valid direction", str))
	}
}
