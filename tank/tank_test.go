package tank

import (
	"strings"
	"testing"
)

var nameToTankDirection = map[string]TankDirection {
	"up": Up,
	"down": Down,
	"left": Left,
	"right": Right,
	"stop": Stop,
}

func TestParseTankDirection(t *testing.T) {
	for name, direction := range nameToTankDirection {
		checkParseTank(t, name, direction)
		checkParseTank(t, strings.ToUpper(name), direction)
	}

	_, err := ParseTankDirection("Invalid")
	if err == nil {
		t.Fatalf("No error from invalid tank direction")
	}
}

func checkParseTank(t *testing.T, name string, expected TankDirection) {
	parsedDirection, err := ParseTankDirection(name)
	if err != nil {
		t.Fatalf("Error parsing %s - %v", name, err)
	}

	if parsedDirection != expected {
		t.Fatalf("Error parsing %s - expected %v, got %v", name, expected, parsedDirection)
	}
}

func TestTankDirectionString(t *testing.T) {
	for name, direction := range nameToTankDirection {
		actual := direction.String()

		if name != actual {
			t.Fatalf("Invalid TankDirection string - expected %s, got %s", name, actual)
		}
	}
}
