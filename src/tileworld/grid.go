package tileworld

import (
	"fmt"
	"math"
	"math/rand"
)

// Grid encapsulates the grid
type Grid struct {
	cols, rows uint8
	objects    [][]*GridObject
	agents     []*GridObject
	tiles      []*GridObject
	holes      []*GridObject
}

// NewGrid constructor
func NewGrid(cols, rows, numAgents, numTiles, numHoles, numObstacles uint8) *Grid {
	g := new(Grid)
	g.cols = cols
	g.rows = rows
	g.objects = make([][]*GridObject, cols)
	for i := 0; i < int(rows); i++ {
		g.objects[i] = make([]*GridObject, rows)
	}
	g.agents = make([]*GridObject, numAgents)
	g.tiles = make([]*GridObject, numTiles)
	g.holes = make([]*GridObject, numHoles)
	for i := uint8(0); i < numAgents; i++ {
		l := g.RandomFreeLocation(cols, rows)
		o := NewGridObject(l, TypeAgent, i)
		g.SetObject(o, l)
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
	fmt.Println("Update")
	for _, a := range g.agents {
		g.updateAgent(a)
	}
	g.printGrid()
}

func (g Grid) printGrid() {
	for r := uint8(0); r < g.rows; r++ {
		for c := uint8(0); c < g.cols; c++ {
			o := g.Object(NewLocation(c, r))
			if o != nil {
				switch o.objectType {
				case TypeAgent:
					if o.hasTile {
						fmt.Print("Å")
					} else {
						fmt.Print("A")
					}
					break
				case TypeTile:
					fmt.Print("T")
					break
				case TypeHole:
					fmt.Print("H")
					break
				case TypeObstacle:
					fmt.Print("#")
					break
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g Grid) updateAgent(o *GridObject) {
	fmt.Printf("UpdateAgent %s\n", o)
	switch o.state {
	case StateIdle:
		g.idle(o)
		break
	case StateToTile:
		g.moveToTile(o)
		break
	case StateToHole:
		g.moveToHole(o)
		break
	}
}

func (g Grid) idle(o *GridObject) {
	o.hasTile = false
	o.tile = g.getClosestTile(o.location)
	o.hole = nil
	o.state = StateToTile
}

func (g Grid) moveToTile(o *GridObject) {
	if o.location.Equals(o.tile.location) {
		o.PickTile()
		g.createTile(o.tile.num)
		o.hole = g.getClosestHole(o.location)
		o.state = StateToHole
		return
	}
	if g.Object(o.tile.location) != o.tile {
		// tile gone
		o.state = StateIdle
		return
	}
	potentialTile := g.getClosestTile(o.location)
	if potentialTile != o.tile {
		o.tile = potentialTile
		o.path = GetPathAStar(&g, o.location, o.tile.location)
	}
	if len(o.path) == 0 {
		o.path = GetPathAStar(&g, o.location, o.tile.location)
		printPath(o.path)
	} else {
		dir := o.path[0]
		o.path = o.path[1:]
		nextLocation := o.location.NextLocation(dir)
		g.move(o, nextLocation)
		fmt.Printf(" -> %s\n", nextLocation)
	}
}
func printPath(path []Direction) {
	fmt.Printf("path:")
	for _, d := range path {
		fmt.Printf("%d ", d)
	}
	fmt.Print("\n")
}
func (g Grid) moveToHole(o *GridObject) {
	if o.location.Equals(o.hole.location) {
		o.DumpTile()
		g.createHole(o.hole.num)
		o.tile = g.getClosestTile(o.location)
		o.state = StateToTile
		o.hasTile = false
		o.hole = nil
		return
	}
	if g.Object(o.hole.location) != o.hole {
		// hole gone
		o.state = StateIdle
		return
	}
	potentialHole := g.getClosestHole(o.location)
	if potentialHole != o.hole {
		o.hole = potentialHole
		o.path = GetPathAStar(&g, o.location, o.hole.location)
	}
	if len(o.path) == 0 {
		o.path = GetPathAStar(&g, o.location, o.hole.location)
		printPath(o.path)
	} else {
		dir := o.path[0]
		o.path = o.path[1:]
		nextLocation := o.location.NextLocation(dir)
		g.move(o, nextLocation)
		fmt.Printf(" -> %s\n", nextLocation)
	}
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
func (g Grid) RandomFreeLocation(cols uint8, rows uint8) *Location {
	c := uint8(rand.Intn(int(cols)))
	r := uint8(rand.Intn(int(rows)))
	l := NewLocation(c, r)
	for !g.isFreeLocation(l) {
		c = uint8(rand.Intn(int(cols)))
		r = uint8(rand.Intn(int(rows)))
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
