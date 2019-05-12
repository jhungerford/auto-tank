// +build mock

package tank

import "log"

type MockTank struct {}

func Init() Tank {
	return MockTank{}
}

func (tank MockTank) Move(direction TankDirection) {
	log.Printf("Mock - move %v\n", direction)
}
