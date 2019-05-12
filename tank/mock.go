// +build mock

package tank

import "log"

type MockTank struct {}

func Init() Tank {
	return MockTank{}
}

func (tank MockTank) Move(direction Direction) {
	log.Printf("Mock - move %v\n", direction)
}
