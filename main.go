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
	status, err := C.wiringPiSetupGpio()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("wiringPiSetupGpio status: %d\n", status)

	// Start the web server
	fs := http.FileServer(http.Dir("web"))

	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
