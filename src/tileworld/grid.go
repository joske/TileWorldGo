package tileworld

import (
	"fmt"
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
		l := g.RandomFreeLocation(cols, rows)
		o := NewGridObject(l, TypeTile, i)
		g.SetObject(o, l)
		g.tiles[i] = o
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
	for c := uint8(0); c < g.cols; c++ {
		for r := uint8(0); r < g.rows; r++ {
			o := g.Object(NewLocation(c, r))
			if o != nil {
				switch o.objectType {
				case TypeAgent:
					fmt.Print("A")
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
	fmt.Printf("UpdateAgent %s\n", o.location)
	d := rand.Intn(4) + 1
	nextLocation := o.Location().NextLocation(Direction(d))
	for !g.isValidLocation(nextLocation) || !g.isFreeLocation(nextLocation) {
		d := rand.Intn(4) + 1
		nextLocation = o.Location().NextLocation(Direction(d))
	}
	g.move(o, nextLocation)
}

func (g Grid) move(o *GridObject, l *Location) {
	oldLocation := o.Location()
	g.SetObject(nil, oldLocation)
	g.SetObject(o, l)
	o.SetLocation(l)
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
	return l.col < g.cols && l.col >= 0 && l.row >= 0 && l.row < g.rows
}
