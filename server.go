package main

import (
	"github.com/jhungerford/auto-tank/tank"
	"net/http"
)

// GET / - serve files from the static directory
// GET /tank/direction - returns the current direction
// PUT /tank/direction - changes the current direction to the direction in the body.

type Server struct {
	tank *tank.Tank
}

func (server *Server) getDirectionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write(server.tank.Direction())
}

func (server *Server) putDirectionHandler(w http.ResponseWriter, r *http.Request) {
	tank.TankDirection()
}


