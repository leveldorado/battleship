package starting

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context) (Field, error)
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) Create(ctx context.Context) (Field, error) {
	g := Game{ID: uuid.New().String(), FirstPlayer: uuid.New().String()}
	f := Field{GameID: g.ID, PlayerID: g.FirstPlayer, Width: defaultSideSize, Height: defaultSideSize}
	f.PlaceShips(DefaultMaxAttempts)
	return f, nil
}
