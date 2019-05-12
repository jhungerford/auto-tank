// +build !mock

package tank

/*
#cgo LDFLAGS: -lm -lwiringPi
#include <errno.h>
#include <wiringPi.h>
 */
import "C"
import (
	"log"
)

type PiTank struct {}

func Init() Tank {
	_, err := C.wiringPiSetup()
	if err != nil {
		log.Fatal(err)
	}

	return PiTank{}
}

func (tank PiTank) Move(direction string) {
	log.Printf("Pi - move %s\n", direction)
}
