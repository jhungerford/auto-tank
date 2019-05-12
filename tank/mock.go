// +build mock

package tank

import "log"

type MockTank struct {}

func Init() Tank {
	return MockTank{}
}

func (tank MockTank) Move(direction string) {
	log.Printf("Mock - move %s\n", direction)
}
