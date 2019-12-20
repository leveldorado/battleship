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

type ShipType string

const (
	ShipTypeSingle    ShipType = "single"
	ShipTypeDouble    ShipType = "double"
	ShipTypeTriple    ShipType = "triple"
	ShipTypeQuadruple ShipType = "quadruple"
)

func (s ShipType) getSize() int {
	switch s {
	case ShipTypeSingle:
		return 1
	case ShipTypeDouble:
		return 2
	case ShipTypeTriple:
		return 3
	case ShipTypeQuadruple:
		return 4
	}
	return 0
}

func (s ShipType) getNumbers() int {
	switch s {
	case ShipTypeSingle:
		return 4
	case ShipTypeDouble:
		return 3
	case ShipTypeTriple:
		return 2
	case ShipTypeQuadruple:
		return 1
	}
	return 0
}

const defaultSideSize = 10

func (f *Field) PlaceShips() {
	cellsNumber := f.Width * f.Height
	cells := make([][]int, f.Height)
	for k := range cells {
		cells[k] = make([]int, f.Height)
	}
	occupiedCells := map[Coordinate]bool{}
	for _, t := range []ShipType{ShipTypeQuadruple, ShipTypeTriple, ShipTypeDouble, ShipTypeSingle} {
		invalidAttempts := map[Coordinate]bool{}
	}
}

func (f *Field) isPositionAvailable() bool {

}
