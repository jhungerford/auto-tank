// +build mock

package tank

import (
	"fmt"
	"os"
)

type MockTank struct {}

func Init() Tank {
	return MockTank{}
}

func (tank MockTank) Move(direction string) {
	fmt.Fprintf(os.Stdout, "Mock - move %s\n", direction)
}
