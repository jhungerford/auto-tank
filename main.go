package main

/*
#cgo LDFLAGS: -lm -lwiringPi
#include <errno.h>
#include <wiringPi.h>
*/
import "C"
import (
	"log"
	"os"
	"fmt"
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
	North :     {Forward, Forward},
	NorthEast : {Forward, Off},
	East :      {Forward, Reverse},
	SouthEast : {Off, Reverse},
	South :     {Reverse, Reverse},
	SouthWest : {Reverse, Off},
	West :      {Reverse, Forward},
	NorthWest : {Off, Forward},
	Stop :      {Off, Off},
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
			C.pinMode(C.int(pin), C.OUTPUT)
		}
	}
}

func (t Tread) move(dir TreadDirection) {
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
			Pins{4, 5, 1},
			Pins{10, 6, 27},
		},
		Tread{
			Pins{0, 7, 23},
			Pins{22, 21, 24},
		},
	}

	tank.init()
	// tank.move(North)

	// Start the web server
	fs := http.FileServer(http.Dir("web"))
	h := http.NewServeMux()
	h.HandleFunc("/api/move", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Move")
	})

	h.Handle("/", http.StripPrefix("/", fs))


	// fs := http.FileServer(http.Dir("web"))

	// http.Handle("/", fs)
	// http.HandleFunc("/api/move", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Move")
	// })
	fmt.Fprintf(os.Stdout, "Running...")
	log.Fatal(http.ListenAndServe(":8080", h))
}
