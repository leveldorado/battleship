package game

import "context"

type Repository interface {
	Create(ctx context.Context, g Game) error
	UpdateField(ctx context.Context, gameID, playerID string, f Field) error
	UpdateTurnOn(ctx context.Context, gameID, turnOn string) error
	UpdateWin(ctx context.Context, gameID, win string) error
}

type Service interface {
	Create(ctx context.Context) (Game, error)
}

func NewService() Service {
	return &service{}
}

type service struct {
	r Repository
}

func (s *service) Create(ctx context.Context) (Game, error) {
	f := Field{}
	f.PlaceShips(DefaultMaxAttempts)
	g := Game{Fields: []Field{f}}
	return g, nil
}
