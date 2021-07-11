package tileworld

import (
	"fmt"
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
	col, row int8
}

// NewLocation is a constructor
func NewLocation(col, row int8) *Location {
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

// GetDirection returns the directino the next location is in
func (l Location) GetDirection(next *Location) Direction {
	if l.row == next.row {
		if l.col == next.col+1 {
			return Left
		} else {
			return Right
		}
	} else {
		if l.row == next.row+1 {
			return Up
		} else {
			return Down
		}
	}
}

// Distance gives the manhattan distance between 2 locations
func (l Location) Distance(o *Location) int {
	return abs(int(l.col)-int(o.col)) + abs(int(l.row)-int(o.row))
}

// Equals equality
func (l Location) Equals(o *Location) bool {
	if l.col == o.col && l.row == o.row {
		return true
	}
	return false
}

func (l Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.col, l.row)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
