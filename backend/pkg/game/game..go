package game

import "time"

type Game struct {
	ID        string
	Fields    []Field
	Players   []string
	TurnOn    string
	Win       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Field struct {
	Player string
	Width  int
	Height int
	Ships  []Ship
	Shots  []Shot
}

type Shot struct {
	Result     ShotResult
	Coordinate Coordinate
}

type Coordinate struct {
	X int
	Y int
}

type ShotResult string

const (
	ShotResultHit  ShotResult = "hit"
	ShotResultMiss ShotResult = "miss"
)

type Ship struct {
	Coordinates []Coordinate
}

func (f *Field) PlaceShips() {

}
