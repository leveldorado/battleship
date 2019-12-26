package main

import (
	"time"

	"github.com/leveldorado/battleship/backend/pkg/http/rest"
	"github.com/leveldorado/battleship/backend/pkg/starting"
	"github.com/tylerb/graceful"
)

func main() {
	g := starting.NewService()
	s := rest.CreateServer(g)
	s.Start(":80")
	graceful.ListenAndServe(s.Server, 5*time.Second)
}
