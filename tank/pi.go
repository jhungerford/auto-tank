// +build !mock

package tank

import (
	"fmt"
	"os"
)

type PiTank struct {}

func Init() Tank {
	return PiTank{}
}

func (tank PiTank) Move(direction string) {
	fmt.Fprintf(os.Stdout, "Pi - move %s\n", direction)
}
