package tileworld

import (
	"fmt"
	"testing"
)

func TestAstar(t *testing.T) {
	from := NewLocation(0, 0)
	to := NewLocation(1, 0)
	var grid = NewGrid(10, 10, 0, 0, 0, 0)
	path := GetPathAStar(grid, from, to)
	if len(path) != 1 {
		t.Error(fmt.Sprintf("path length: %d", len(path)))
	}
	if !path[0].Equals(to) {
		t.Error(fmt.Sprintf("path does not lead to goal : %s", path[0]))
	}
}

func TestAstar2(t *testing.T) {
	from := NewLocation(0, 0)
	to := NewLocation(1, 1)
	var grid = NewGrid(10, 10, 0, 0, 0, 0)
	path := GetPathAStar(grid, from, to)
	if len(path) != 2 {
		t.Error(fmt.Sprintf("path length: %d", len(path)))
	}
	if !path[1].Equals(to) {
		t.Error(fmt.Sprintf("path does not lead to goal : %s", path[0]))
	}
}

func TestAstar3(t *testing.T) {
	from := NewLocation(0, 0)
	to := NewLocation(2, 2)
	var grid = NewGrid(10, 10, 0, 0, 0, 0)
	path := GetPathAStar(grid, from, to)
	fmt.Printf("path:%s", path)
	if len(path) != 4 {
		t.Error(fmt.Sprintf("path length: %d", len(path)))
	}
	if !path[3].Equals(to) {
		t.Error(fmt.Sprintf("path does not lead to goal : %s", path[0]))
	}
}
