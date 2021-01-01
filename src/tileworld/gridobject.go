package tileworld

import "fmt"

// ObjectType defines the type
type ObjectType uint8

// Types
const (
	TypeAgent    ObjectType = 1
	TypeTile     ObjectType = 2
	TypeHole     ObjectType = 3
	TypeObstacle ObjectType = 4
)

// GridObject an object on the grid
type GridObject struct {
	location   *Location
	num        uint8
	score      int
	objectType ObjectType
	hasTile    bool
	tile       *GridObject
	hole       *GridObject
}

// NewGridObject create
func NewGridObject(l *Location, t ObjectType, n uint8) *GridObject {
	o := new(GridObject)
	o.location = l
	o.objectType = t
	o.num = n
	o.score = 0
	return o
}

// Location return this object's location
func (o *GridObject) Location() *Location {
	return o.location
}

// SetLocation move this object
func (o *GridObject) SetLocation(l *Location) {
	o.location = l
}

// SetTile assign a tile
func (o *GridObject) SetTile(t *GridObject) {
	o.tile = t
}

// SetHole assign a tile
func (o *GridObject) SetHole(t *GridObject) {
	o.hole = t
}

// PickTile we now have a tile
func (o *GridObject) PickTile() {
	o.hasTile = true
}

// DumpTile we now have a tile
func (o *GridObject) DumpTile() {
	o.hasTile = false
}

func (o *GridObject) String() string {
	var s string
	switch o.objectType {
	case TypeAgent:
		s = "Agent"
		break
	case TypeTile:
		s = "Tile"
		break
	case TypeHole:
		s = "Hole"
		break
	default:
		s = "Obstacle"
		break
	}
	return fmt.Sprintf("%s(%d) @%s", s, o.num, o.location)
}
