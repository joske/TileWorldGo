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
}

type Agent struct {
	GridObject
	state   State
	hasTile bool
	tile    *GridObject
	hole    *GridObject
	path    []Location
}

// NewGridObject create
func NewGridObject(l *Location, t ObjectType, n uint8) *GridObject {
	o := new(GridObject)
	o.location = l
	o.num = n
	o.score = 0
	o.objectType = t
	return o
}

func NewAgent(l *Location, t ObjectType, n uint8) *Agent {
	o := new(Agent)
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
func (o *Agent) SetTile(t *GridObject) {
	o.tile = t
}

// SetHole assign a tile
func (o *Agent) SetHole(t *GridObject) {
	o.hole = t
}

// PickTile we now have a tile
func (o *Agent) PickTile() {
	o.hasTile = true
}

// DumpTile we now have a tile
func (o *Agent) DumpTile() {
	o.hasTile = false
	o.score += o.tile.score
}
