package tank

import "testing"

func TestStubTank(t *testing.T) {
	stub := StubTank{Moves: nil}

	checkMoves(t, nil, stub.Moves)

	stub.Move(Up)
	stub.Move(Down)
	stub.Move(Stop)

	checkMoves(t, []TankDirection{Up, Down, Stop}, stub.Moves)

	stub.Reset()

	checkMoves(t, nil, stub.Moves)
}

func checkMoves(t *testing.T, expected []TankDirection, actual []TankDirection) {
	if len(expected) != len(actual) {
		t.Fatalf("Expected and actual have different sizes.  Expected: %v, Actual: %v", expected, actual)
	}

	for i, value := range expected {
		if actual[i] != value {
			t.Fatalf("Moves do not match.  Expected: %v, Actual: %v", expected, actual)
		}
	}
}