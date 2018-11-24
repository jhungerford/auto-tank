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

func main() {
	status, err := C.wiringPiSetup()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("wiringPiSetupGpio status: %d\n", status)

	// left side of the tank
	C.pinMode(4, C.OUTPUT)
	C.pinMode(5, C.OUTPUT)
	C.pinMode(10, C.OUTPUT)
	C.pinMode(7, C.OUTPUT)
	C.pinMode(1, C.OUTPUT)
	C.pinMode(27, C.OUTPUT)

	C.digitalWrite(4, 1)
	C.digitalWrite(5, 0)
	C.digitalWrite(10, 1)
	C.digitalWrite(7, 0)
	C.digitalWrite(1, 1)
	C.digitalWrite(27, 1)

	// right side
	C.pinMode(7, C.OUTPUT)
	C.pinMode(0, C.OUTPUT)
	C.pinMode(2, C.OUTPUT)
	C.pinMode(3, C.OUTPUT)
	C.pinMode(23, C.OUTPUT)
	C.pinMode(24, C.OUTPUT)

	C.digitalWrite(7, 0)
	C.digitalWrite(0, 1)
	C.digitalWrite(2, 0)
	C.digitalWrite(3, 1)
	C.digitalWrite(23, 1)
	C.digitalWrite(24, 1)

	// Start the web server
	fs := http.FileServer(http.Dir("web"))

	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
