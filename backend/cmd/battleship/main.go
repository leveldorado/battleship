package main

import (
	"time"

	"github.com/leveldorado/battleship/backend/pkg/game"
	"github.com/leveldorado/battleship/backend/pkg/http/rest"
	"github.com/tylerb/graceful"
)

func main() {
	g := game.NewService()
	s := rest.CreateServer(g)
	s.Start(":80")
	graceful.ListenAndServe(s.Server, 5*time.Second)
}
