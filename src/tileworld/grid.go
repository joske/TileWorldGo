package tileworld

import (
	"fmt"
	"math"
	"math/rand"
)

// Grid encapsulates the grid
type Grid struct {
	cols, rows int8
	objects    [][]*GridObject
	agents     []*Agent
	tiles      []*GridObject
	holes      []*GridObject
}

// NewGrid constructor
func NewGrid(cols, rows int8, numAgents, numTiles, numHoles, numObstacles uint8) *Grid {
	g := new(Grid)
	g.cols = cols
	g.rows = rows
	g.objects = make([][]*GridObject, cols)
	for i := 0; i < int(rows); i++ {
		g.objects[i] = make([]*GridObject, rows)
	}
	g.agents = make([]*Agent, numAgents)
	g.tiles = make([]*GridObject, numTiles)
	g.holes = make([]*GridObject, numHoles)
	for i := uint8(0); i < numAgents; i++ {
		l := g.RandomFreeLocation(cols, rows)
		o := NewAgent(l, TypeAgent, i)
		g.SetObject(&o.GridObject, l)
		g.agents[i] = o
	}
	for i := uint8(0); i < numTiles; i++ {
		g.createTile(i)
	}
	for i := uint8(0); i < numHoles; i++ {
		l := g.RandomFreeLocation(cols, rows)
		o := NewGridObject(l, TypeHole, i)
		g.SetObject(o, l)
		g.holes[i] = o
	}
	for i := uint8(0); i < numObstacles; i++ {
		l := g.RandomFreeLocation(cols, rows)
		o := NewGridObject(l, TypeObstacle, i)
		g.SetObject(o, l)
	}
	return g
}

func (g Grid) createTile(i uint8) {
	l := g.RandomFreeLocation(g.cols, g.rows)
	o := NewGridObject(l, TypeTile, i)
	o.score = rand.Intn(6) + 1
	g.SetObject(o, l)
	g.tiles[i] = o
}

func (g Grid) createHole(i uint8) {
	l := g.RandomFreeLocation(g.cols, g.rows)
	o := NewGridObject(l, TypeHole, i)
	g.SetObject(o, l)
	g.holes[i] = o
}

// Object get an object at given coordinates
func (g Grid) Object(l *Location) *GridObject {
	return g.objects[l.col][l.row]
}

// SetObject set an object
func (g Grid) SetObject(o *GridObject, l *Location) {
	g.objects[l.col][l.row] = o
}

// Update update grid
func (g Grid) Update() {
	for _, a := range g.agents {
		g.updateAgent(a)
	}
	g.printGrid()
}

func (g Grid) printGrid() {
	for r := int8(0); r < g.rows; r++ {
		for c := int8(0); c < g.cols; c++ {
			o := g.Object(NewLocation(c, r))
			if o != nil {
				switch o.objectType {
				case TypeAgent:
					a := g.agents[o.num]
					if a.hasTile {
						fmt.Print("Å")
					} else {
						fmt.Print("A")
					}
				case TypeTile:
					fmt.Print("T")
				case TypeHole:
					fmt.Print("H")
				case TypeObstacle:
					fmt.Print("#")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	for _, a := range g.agents {
		fmt.Printf("Agent %d: %d\n", a.num, a.score)
	}
}

func (g Grid) updateAgent(o *Agent) {
	switch o.state {
	case StateIdle:
		g.idle(o)
	case StateToTile:
		g.moveToTile(o)
	case StateToHole:
		g.moveToHole(o)
	}
}

func (g Grid) idle(o *Agent) {
	o.tile = g.getClosestTile(o.location)
	o.state = StateToTile
}

func (g Grid) moveToTile(o *Agent) {
	g.moveToObject(o, o.tile)
}

func (g Grid) moveToHole(o *Agent) {
	g.moveToObject(o, o.hole)
}

func (g Grid) moveToObject(agent *Agent, o *GridObject) {
	if agent.location.Equals(o.location) {
		if agent.state == StateToTile {
			g.PickTile(agent, o)
		} else {
			g.DumpTile(agent, o)
		}
		return
	}
	if g.Object(o.location) != o {
		if agent.state == StateToTile {
			// tile gone
			agent.tile = g.getClosestTile(agent.location)
			agent.path = GetPathAStar(&g, agent.location, agent.tile.location)
		} else {
			// hole gone
			agent.hole = g.getClosestHole(agent.location)
			agent.path = GetPathAStar(&g, agent.location, agent.hole.location)
		}
		return
	}

	if len(agent.path) == 0 {
		agent.path = GetPathAStar(&g, agent.location, o.location)
	}
	if len(agent.path) > 0 {
		nextLocation := &agent.path[0]
		if g.isValidLocation(nextLocation) || nextLocation.Equals(o.location) {
			agent.path = agent.path[1:]
			g.move(&agent.GridObject, nextLocation)
		} else {
			agent.path = nil
		}
	}
}

func (g Grid) DumpTile(agent *Agent, o *GridObject) {
	agent.DumpTile()
	g.createHole(agent.hole.num)
	agent.tile = g.getClosestTile(agent.location)
	agent.state = StateToTile
	agent.hasTile = false
	agent.hole = nil
	agent.path = nil
}

func (g Grid) PickTile(agent *Agent, o *GridObject) {
	agent.PickTile()
	g.createTile(agent.tile.num)
	agent.hole = g.getClosestHole(agent.location)
	agent.state = StateToHole
}

func (g Grid) move(o *GridObject, l *Location) {
	oldLocation := o.Location()
	g.SetObject(nil, oldLocation)
	g.SetObject(o, l)
	o.SetLocation(l)
}

func (g Grid) getClosestTile(l *Location) *GridObject {
	return g.getClosestObject(TypeTile, g.tiles, l)
}

func (g Grid) getClosestHole(l *Location) *GridObject {
	return g.getClosestObject(TypeHole, g.holes, l)
}

func (g Grid) getClosestObject(t ObjectType, a []*GridObject, l *Location) *GridObject {
	closest := math.MaxInt64
	var best *GridObject
	for i := 0; i < len(a); i++ {
		o := a[i]
		dist := o.location.Distance(l)
		if dist < closest {
			closest = dist
			best = o
		}
	}
	return best
}

// RandomFreeLocation get a free location on the grid
func (g Grid) RandomFreeLocation(cols int8, rows int8) *Location {
	c := int8(rand.Intn(int(cols)))
	r := int8(rand.Intn(int(rows)))
	l := NewLocation(c, r)
	for !g.isFreeLocation(l) {
		c = int8(rand.Intn(int(cols)))
		r = int8(rand.Intn(int(rows)))
		l = NewLocation(c, r)
	}
	return l
}

func (g Grid) isFreeLocation(l *Location) bool {
	return g.Object(l) == nil
}

func (g Grid) isValidLocation(l *Location) bool {
	return l.col < g.cols && l.col >= 0 && l.row >= 0 && l.row < g.rows && g.isFreeLocation(l)
}
