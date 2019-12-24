package game

import (
	"time"
)

type Game struct {
	ID        string
	Fields    []Field
	Players   []string
	TurnOn    string
	Win       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Shot struct {
	Result     ShotResult
	Coordinate Coordinate
}

type ShotResult string

const (
	ShotResultHit  ShotResult = "hit"
	ShotResultMiss ShotResult = "miss"
)
