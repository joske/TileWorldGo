package tileworld

import (
	"fmt"
)

// ObjectType defines the type
type ObjectType uint8

// Types
const (
	TypeAgent    ObjectType = 1
	TypeTile     ObjectType = 2
	TypeHole     ObjectType = 3
	TypeObstacle ObjectType = 4
)

// State state of the agent
type State uint8

// State state of the agent
const (
	StateIdle   State = 0
	StateToTile State = 1
	StateToHole State = 2
)

// GridObject an object on the grid
type GridObject struct {
	location   *Location
	num        uint8
	score      int
	objectType ObjectType
	state      State
	hasTile    bool
	tile       *GridObject
	hole       *GridObject
	path       []Direction
}

// NewGridObject create
func NewGridObject(l *Location, t ObjectType, n uint8) *GridObject {
	o := new(GridObject)
	o.location = l
	o.objectType = t
	o.num = n
	o.score = 0
	o.state = StateIdle
	o.hasTile = false
	o.tile = nil
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
	fmt.Printf("%s - pickTile\n", o)
	o.hasTile = true
}

// DumpTile we now have a tile
func (o *GridObject) DumpTile() {
	fmt.Printf("%s - DumpTile\n", o)
	o.hasTile = false
	o.score += o.tile.score
}

func (o *GridObject) String() string {
	var s string
	switch o.objectType {
	case TypeAgent:
		return fmt.Sprintf("Agent(%d) @%s in state %d, hasTile=%t, tile=%s, hole=%s", o.num, o.location, o.state, o.hasTile, o.tile, o.hole)
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
