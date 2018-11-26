package main

/*
#cgo LDFLAGS: -lm -lwiringPi
#include <errno.h>
#include <wiringPi.h>
*/
import "C"
import (
	"log"
	"net/http"
)

type TreadDirection int
const (
	Forward = TreadDirection(iota)
	Reverse = TreadDirection(iota)
	Off     = TreadDirection(iota)
)

type TankDirection int
const (
	North     = TankDirection(iota)
	NorthEast = TankDirection(iota)
	East      = TankDirection(iota)
	SouthEast = TankDirection(iota)
	South     = TankDirection(iota)
	SouthWest = TankDirection(iota)
	West      = TankDirection(iota)
	NorthWest = TankDirection(iota)
	Stop      = TankDirection(iota)
)

type tankTreadDirection struct {
	left, right TreadDirection
}

var tankDirectionMap = map[TankDirection]tankTreadDirection{
	North : tankTreadDirection{Forward, Forward},
	NorthEast : tankTreadDirection{Forward, Off},
	East : tankTreadDirection{Forward, Reverse},
	SouthEast : tankTreadDirection{Off, Reverse},
	South : tankTreadDirection{Reverse, Reverse},
	SouthWest : tankTreadDirection{Reverse, Off},
	West : tankTreadDirection{Reverse, Forward},
	NorthWest : tankTreadDirection{Off, Forward},
	Stop : tankTreadDirection{Off, Off},
}

type Pins struct {
	HighPin, LowPin, SpeedPin int
}

type Tread struct {
	Front, Rear Pins
}

func (t Tread) init() {
	for _, pins := range []Pins{t.Front, t.Rear} {
		for _, pin := range []int{pins.HighPin, pins.LowPin, pins.SpeedPin} {
			C.pinMode(pin, C.OUTPUT)
		}
	}
}

func (t Tread) move(dir TreadDirection) {
	for _, pins := range []Pins{t.Front, t.Rear} {
		switch dir {
		case Forward:
			C.digitalWrite(pins.LowPin, 0)
			C.digitalWrite(pins.HighPin, 1)
			C.digitalWrite(pins.SpeedPin, 1)
		case Reverse:
			C.digitalWrite(pins.LowPin, 1)
			C.digitalWrite(pins.HighPin, 0)
			C.digitalWrite(pins.SpeedPin, 1)
		case Off:
			C.digitalWrite(pins.SpeedPin, 0)
		}
	}
}

type Tank struct {
	Left, Right Tread
}

func (t Tank) init() {
	t.Left.init()
	t.Right.init()
}

func (t Tank) move(direction TankDirection) {
	treadDirs := tankDirectionMap[direction]
	t.Left.move(treadDirs.left)
	t.Right.move(treadDirs.right)
}

func main() {
	_, err := C.wiringPiSetup()
	if err != nil {
		log.Fatal(err)
	}

	tank := Tank{
		Tread{
			Pins{7, 0, 23},
			Pins{2, 3, 24},
		},
		Tread{
			Pins{5, 4, 1},
			Pins{7, 10, 27},
		},
	}

	tank.init()
	tank.move(North)

	// Start the web server
	fs := http.FileServer(http.Dir("web"))

	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
