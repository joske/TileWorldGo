package tileworld

import "testing"

func testGridOneObject(t *testing.T) {
	var grid = NewGrid(10, 10)
	l := NewLocation(1, 1)
	grid.SetObject(NewGridObject(l, TypeAgent, 1), l)
	o := grid.Object(l)
	if o == nil {
		t.Error("could not get object")
	}
	if o.location.col != 1 {
		t.Error("bad column")
	}
	if o.location.row != 1 {
		t.Error("bad row")
	}
}
