package tank

import (
	"github.com/jhungerford/auto-tank/tank"
	"log"
	"net/http"
)

func main() {
	// Test drive the tank.
	t := tank.New()

	err := t.Init()
	if err != nil {
		log.Fatal(err)
	}

	t.Move(tank.North)

	// Start the web server
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}