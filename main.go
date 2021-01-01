package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"sourcery.be/tileworld/src/tileworld"
)

// some globals
const (
	COLS = 10
	ROWS = 10
)

func main() {
	fmt.Println("start")
	gtk.Init(nil)
	fmt.Println("gtk inited")
	grid := tileworld.NewGrid(COLS, ROWS, 1, 1, 1, 1)
	fmt.Println("grid created")
	view := tileworld.GridViewNew(grid)
	fmt.Println("view created")
	view.Show()
	fmt.Println("view shown")
	gtk.Main()
}
