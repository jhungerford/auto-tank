package server

import (
	"github.com/jhungerford/auto-tank/tank"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	tank   tank.Tank
	router *http.ServeMux
}

func NewServer(tank tank.Tank) *server {
	s := &server{
		tank:   tank,
		router: http.NewServeMux(),
	}

	s.routes()

	return s
}

func (s *server) routes() {
	s.router.Handle("/", http.FileServer(http.Dir("static")))
	s.router.HandleFunc("/move", s.handleMove())
}

func (s *server) handleMove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading move body: %v", err)
			http.Error(w, "Error reading body", http.StatusBadRequest)
			return
		}

		direction, err := tank.ParseTankDirection(string(body))
		if err != nil {
			log.Printf("Invalid direction from body: %v", err)
			http.Error(w, "Invalid direction", http.StatusBadRequest)
			return
		}

		s.tank.Move(direction)

		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}