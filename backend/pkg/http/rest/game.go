package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/leveldorado/battleship/backend/pkg/starting"
)

func createGame(g starting.Service) func(echo.Context) error {
	return func(ctx echo.Context) error {
		g, err := g.Create(ctx.Request().Context())
		if err != nil {
			return fmt.Errorf(`failed to create game: [err: %w]`, err)
		}
		return ctx.JSON(http.StatusOK, g)
	}
}
