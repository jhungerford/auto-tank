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

type Direction int

const (
	Forward = Direction(iota)
	Reverse = Direction(iota)
	Left    = Direction(iota)
	Right   = Direction(iota)
	Stop    = Direction(iota)
)

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

func (t Tread) move(dir Direction) {
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
		case Stop:
			C.digitalWrite(pins.SpeedPin, 0)
		default:
			panic("Invalid tread direction")
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

func (t Tank) move(direction Direction) {
	switch direction {
	case Forward:
		t.Left.move(Forward)
		t.Right.move(Forward)
	case Reverse:
		t.Left.move(Reverse)
		t.Right.move(Reverse)
	case Left:
		t.Left.move(Reverse)
		t.Right.move(Forward)
	case Right:
		t.Left.move(Forward)
		t.Right.move(Reverse)
	case Stop:
		t.Left.move(Stop)
		t.Right.move(Stop)
	}
}

2
func main() {
	_, err := C.wiringPiSetup()
	if err != nil {
		log.Fatal(err)
	}

	tank := Tank{
		Tread{
			Pins{4, 5, 1},
			Pins{10, 7, 27},
		},
		Tread{
			Pins{7, 0, 23},
			Pins{2, 3, 24},
		},
	}

	tank.init()
	tank.move(Forward)

	// Start the web server
	fs := http.FileServer(http.Dir("web"))

	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
