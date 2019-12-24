package game

import "math/rand"

type Field struct {
	Player string
	Width  int
	Height int
	Ships  []Ship
	Shots  []Shot
}

func (f *Field) PlaceShips(maxAttempts int) {
	cells := make([][]int, f.Height)
	for k := range cells {
		cells[k] = make([]int, f.Height)
	}
	rotations := []shipRotation{shipRotationHorizontal, shipRotationVertical}
	occupiedCells := map[Coordinate]bool{}
	for _, t := range []ShipType{ShipTypeQuadruple, ShipTypeTriple, ShipTypeDouble, ShipTypeSingle} {
		for i := 0; i < t.getNumbers(); i++ {
			invalidAttempts := map[Coordinate]bool{}
			rotation := rotations[rand.Intn(len(rotations))]
			for i := 0; i < maxAttempts; i++ {
				startPosition := Coordinate{}
				for {
					startPosition = Coordinate{X: rand.Intn(f.Width), Y: rand.Intn(f.Height)}
					if occupiedCells[startPosition] {
						continue
					}
					if invalidAttempts[startPosition] {
						continue
					}
					break
				}
				cells, ok := getShipCoordinates(startPosition, occupiedCells, t.getSize(), rotation)
				if !ok {
					invalidAttempts[startPosition] = true
					continue
				}
				f.Ships = append(f.Ships, Ship{Coordinates: cells})
			}
		}
	}
}
