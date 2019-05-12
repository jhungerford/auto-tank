// +build mock

package tank

import "log"

// MockTank logs tank movements - it's useful for developing auto-tank's web page.
type MockTank struct {}

func Init() *MockTank {
	return &MockTank{}
}

func (tank *MockTank) Move(direction TankDirection) {
	log.Printf("Mock - move %v\n", direction)
}
