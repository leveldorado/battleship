package rest

import (
	"github.com/labstack/echo"
	"github.com/leveldorado/battleship/backend/pkg/game"
)

func CreateServer(g game.Service) *echo.Echo {
	e := echo.New()
	e.POST("/api/v1/game", createGame(g))
	return e
}
