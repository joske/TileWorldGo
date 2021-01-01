package tileworld

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
