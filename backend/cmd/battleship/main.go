package main

import (
	"log"
	"time"

	"github.com/tylerb/graceful"

	"github.com/leveldorado/battleship/backend/pkg/http/rest"
	"github.com/leveldorado/battleship/backend/pkg/starting"
)

func main() {
	g := starting.NewService()
	s := rest.CreateServer(g)
	s.Server.Addr = ":80"
	err := graceful.ListenAndServe(s.Server, 5*time.Second)
	log.Println("server stopped with error:", err)
}
