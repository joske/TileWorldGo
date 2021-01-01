package tileworld

import "testing"

func testNextLocation(t *testing.T) {
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
