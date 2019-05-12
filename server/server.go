package server

import (
	"github.com/jhungerford/auto-tank/tank"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	tank   *tank.Tank
	router *http.ServeMux
}

func Init(tank *tank.Tank) *server {
	s := &server{
		tank:   tank,
		router: http.NewServeMux(),
	}

	s.routes()

	return s
}

func (s *server) routes() {
	fs := http.FileServer(http.Dir("static"))

	s.router.Handle("/", fs)
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

		(*s.tank).Move(string(body))

		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}