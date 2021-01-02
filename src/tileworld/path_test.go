package tileworld

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	from := NewLocation(0, 0)
	to := NewLocation(1, 0)
	var grid = NewGrid(10, 10, 0, 0, 0, 0)
	path := GetPath(grid, from, to)
	if len(path) != 1 {
		t.Error(fmt.Sprintf("path length: %d", len(path)))
	}
}
