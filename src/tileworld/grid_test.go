package tileworld

import "testing"

func TestGridOneObject(t *testing.T) {
	var grid = NewGrid(10, 10, 1, 0, 0, 0)
	o := grid.agents[0]
	if o == nil {
		t.Error("could not get object")
	}
}
