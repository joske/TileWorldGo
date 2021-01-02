package tileworld

import (
	"fmt"
	"testing"
)

func TestNextLocation(t *testing.T) {
	var initial = NewLocation(10, 10)
	var next = initial.NextLocation(Up)
	if next.col != 10 {
		t.Error("alles kapot")
	}
	if next.row != 9 {
		t.Error("alles kapot")
	}
	next = initial.NextLocation(Down)
	if next.col != 10 {
		t.Error("alles kapot")
	}
	if next.row != 11 {
		t.Error("alles kapot")
	}
	next = initial.NextLocation(Left)
	if next.col != 9 {
		t.Error("alles kapot")
	}
	if next.row != 10 {
		t.Error("alles kapot")
	}
	next = initial.NextLocation(Right)
	if next.col != 11 {
		t.Error("alles kapot")
	}
	if next.row != 10 {
		t.Error("alles kapot")
	}
}

func TestDistance(t *testing.T) {
	initial := NewLocation(0, 0)
	next := NewLocation(1, 1)
	dist := initial.Distance(next)
	if dist != 2 {
		t.Error(fmt.Sprintf("ai ai ai: %d", dist))
	}
}
