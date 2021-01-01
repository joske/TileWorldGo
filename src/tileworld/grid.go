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
}

// NewGrid constructor
func NewGrid(cols, rows uint8) *Grid {
	g := new(Grid)
	g.cols = cols
	g.rows = rows
	g.objects = make([][]*GridObject, cols)
	fmt.Println(g.objects)
	for i := 0; i < int(rows); i++ {
		g.objects[i] = make([]*GridObject, rows)
	}
	fmt.Println(g.objects)
	g.agents = make([]*GridObject, 1)
	l := g.RandomFreeLocation(cols, rows)
	o := NewGridObject(l, TypeAgent, 1)
	g.SetObject(o, l)
	g.agents[0] = o
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
	fmt.Println(g.objects)
}

func (g Grid) updateAgent(o *GridObject) {
	fmt.Printf("UpdateAgent %s\n", o.location)
	d := rand.Intn(4) + 1
	nextLocation := o.Location().NextLocation(Direction(d))
	for !g.isFreeLocation(nextLocation) {
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
