package game

import "context"

type Repository interface {
	Create(ctx context.Context, g Game) error
	UpdateField(ctx context.Context, gameID, playerID string, f Field) error
	UpdateTurnOn(ctx context.Context, gameID, turnOn string) error
	UpdateWin(ctx context.Context, gameID, win string) error
}

type Service interface {
}

type service struct {
	r Repository
}
