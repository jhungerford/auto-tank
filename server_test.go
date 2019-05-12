package main

import (
	"github.com/jhungerford/auto-tank/server"
	"github.com/jhungerford/auto-tank/tank"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Paths in the server are relative to the root directory, so this test is located there too.
func TestIndex(t *testing.T) {
	stubTank := &tank.StubTank{Moves: nil}

	s := server.NewServer(stubTank)

	req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatalf("Error creating request for index: %v", err)
	}

	w := httptest.NewRecorder()
	s.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Wrong status - expected %d, got %d", http.StatusOK, w.Code)
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Error reading index body: %v", err)
	}

	if !strings.Contains(string(body), "html") {
		t.Fatalf("Expected HTML from the index.")
	}
}

func TestMove(t *testing.T) {
	stubTank := &tank.StubTank{Moves: nil}

	s := server.NewServer(stubTank)

	req, err := http.NewRequest(http.MethodPost, "/move", strings.NewReader("stop"))
	if err != nil {
		t.Fatalf("Error creating request for move: %v", err)
	}

	w := httptest.NewRecorder()
	s.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("Wrong status - expected %d, got %d", http.StatusNoContent, w.Code)
	}

	_, err = ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Error reading index body: %v", err)
	}

	if len(stubTank.Moves) != 1 || stubTank.Moves[0] != tank.Stop {
		t.Fatalf("Move endpoint should move the tank - moves: %v", stubTank.Moves)
	}
}
