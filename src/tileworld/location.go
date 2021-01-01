package tileworld

import (
	"fmt"
	"math"
)

// Direction is a direction
type Direction uint8

// directions
const (
	Up    Direction = 1
	Down  Direction = 2
	Left  Direction = 3
	Right Direction = 4
)

// Location is a location on the grid: col = x, row = y
type Location struct {
	col, row uint8
}

// NewLocation is a constructor
func NewLocation(col, row uint8) *Location {
	l := new(Location)
	l.col = col
	l.row = row
	return l
}

// NextLocation creates a new Location in the given direction
func (l Location) NextLocation(dir Direction) *Location {
	switch dir {
	case Up:
		return NewLocation(l.col, l.row-1)
	case Down:
		return NewLocation(l.col, l.row+1)
	case Left:
		return NewLocation(l.col-1, l.row)
	case Right:
		return NewLocation(l.col+1, l.row)
	}
	return nil
}

// Distance gives the manhattan distance between 2 locations
func (l Location) Distance(o Location) int {
	return int(math.Abs(float64(l.col-o.col)) + math.Abs(float64(l.row-o.row)))
}

func (l Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.col, l.row)
}
