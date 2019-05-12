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

type PiTank struct {
	left, right piTread
}

func Init() Tank {
	_, err := C.wiringPiSetup()
	if err != nil {
		log.Fatal(err)
	}

	t := PiTank{
		left: piTread{
			front: piPins{C.int(4), C.int(5), C.int(1)},
			rear:  piPins{C.int(10), C.int(6), C.int(27)},
		},
		right: piTread{
			front: piPins{C.int(0), C.int(7), C.int(23)},
			rear:  piPins{C.int(22), C.int(21), C.int(24)},
		},
	}

	t.left.init()
	t.right.init()

	return t
}

type tankTreadDirection struct {
	left, right TreadDirection
}

var tankDirectionMap = map[TankDirection]tankTreadDirection{
	Up:    {Forward, Forward},
	Down:  {Reverse, Reverse},
	Left:  {Reverse, Forward},
	Right: {Forward, Reverse},
	Stop:  {Off, Off},
}

func (tank PiTank) Move(direction TankDirection) {
	log.Printf("Pi - move %v\n", direction)

	treadDirections := tankDirectionMap[direction]
	tank.left.Move(treadDirections.left)
	tank.right.Move(treadDirections.right)
}

type piPins struct {
	high, low, speed C.int
}

type piTread struct {
	front, rear piPins
}

func (tread piTread) init() {
	for _, pins := range []piPins{tread.front, tread.rear} {
		for _, pin := range []C.int{pins.high, pins.low, pins.speed} {
			C.pinMode(pin, C.OUTPUT)
		}
	}
}

func (tread piTread) Move(direction TreadDirection) {
	// TODO: throttle speed - switch from a digital 1 or 0 to a PWM range
	for _, pins := range []piPins{tread.front, tread.rear} {
		switch direction {
		case Forward:
			C.digitalWrite(pins.low, 0)
			C.digitalWrite(pins.high, 1)
			C.digitalWrite(pins.speed, 1)
		case Reverse:
			C.digitalWrite(pins.low, 1)
			C.digitalWrite(pins.high, 0)
			C.digitalWrite(pins.speed, 1)
		case Off:
			C.digitalWrite(pins.speed, 0)
		}
	}
}
