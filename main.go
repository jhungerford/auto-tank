package main

import (
	"github.com/jhungerford/auto-tank/server"
	"github.com/jhungerford/auto-tank/tank"
	"log"
	"net/http"
)

func main() {
	var t = tank.Init()
	var s = server.NewServer(t)

	log.Fatal(http.ListenAndServe(":8080", s))
}
