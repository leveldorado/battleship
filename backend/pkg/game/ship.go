package game

type Coordinate struct {
	X int
	Y int
}

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
const DefaultMaxAttempts = 3

type shipRotation string

const (
	shipRotationHorizontal = "horizontal"
	shipRotationVertical   = "vertical"
)

func getShipCoordinates(startPosition Coordinate, occupiedCells map[Coordinate]bool, size int, rotation shipRotation) ([]Coordinate, bool) {
	cells := []Coordinate{startPosition}
	currentCell := startPosition
	for i := 0; i < size; i++ {
		switch rotation {
		case shipRotationVertical:
			currentCell.Y++
		case shipRotationHorizontal:
			currentCell.X++
		}
		if occupiedCells[currentCell] {
			return nil, false
		}
		cells = append(cells, currentCell)
	}
	return cells, true
}
