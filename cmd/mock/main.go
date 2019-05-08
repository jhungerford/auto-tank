package main

import (
	"log"
	"net/http"
	"tank/mock"
)

func main() {
	// Start the web server
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
