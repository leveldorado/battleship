package rest

import (
	"github.com/labstack/echo"
	"github.com/leveldorado/battleship/backend/pkg/starting"
)

func CreateServer(g starting.Service) *echo.Echo {
	e := echo.New()
	e.POST("/api/v1/game", createGame(g))
	return e
}
