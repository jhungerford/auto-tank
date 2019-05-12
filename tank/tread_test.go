package tank

import (
	"strings"
	"testing"
)

var nameToTreadDirection = map[string]TreadDirection {
	"forward": Forward,
	"reverse": Reverse,
	"off": Off,
}

func TestParseTreadDirection(t *testing.T) {
	for name, direction := range nameToTreadDirection {
		checkParseTread(t, name, direction)
		checkParseTread(t, strings.ToUpper(name), direction)
	}

	_, err := ParseTreadDirection("Invalid")
	if err == nil {
		t.Fatalf("No error from invalid tread direction")
	}
}

func checkParseTread(t *testing.T, name string, expected TreadDirection) {
	parsedDirection, err := ParseTreadDirection(name)
	if err != nil {
		t.Fatalf("Error parsing %s - %v", name, err)
	}

	if parsedDirection != expected {
		t.Fatalf("Error parsing %s - expected %v, got %v", name, expected, parsedDirection)
	}
}

func TestTreadDirectionString(t *testing.T) {
	for name, direction := range nameToTreadDirection {
		actual := direction.String()

		if name != actual {
			t.Fatalf("Invalid TreadDirection string - expected %s, got %s", name, actual)
		}
	}
}