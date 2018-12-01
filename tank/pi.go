//+build !windows

package tank

/*
#cgo LDFLAGS: -lm -lwiringPi
#include <errno.h>
#include <wiringPi.h>
*/
import "C"

func (t Tank) Init() error {
	_, err := C.wiringPiSetup()
	t.Left.init()
	t.Right.init()
	return err
}

func (t Tread) Init() {
	for _, pins := range []Pins{t.Front, t.Rear} {
		for _, pin := range []int{pins.HighPin, pins.LowPin, pins.SpeedPin} {
			C.pinMode(C.int(pin), C.OUTPUT)
		}
	}
}

func (t Tread) Move(dir TreadDirection) {
	for _, pins := range []Pins{t.Front, t.Rear} {
		switch dir {
		case Forward:
			C.digitalWrite(C.int(pins.LowPin), 0)
			C.digitalWrite(C.int(pins.HighPin), 1)
			C.digitalWrite(C.int(pins.SpeedPin), 1)
		case Reverse:
			C.digitalWrite(C.int(pins.LowPin), 1)
			C.digitalWrite(C.int(pins.HighPin), 0)
			C.digitalWrite(C.int(pins.SpeedPin), 1)
		case Off:
			C.digitalWrite(C.int(pins.SpeedPin), 0)
		}
	}
}
