package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"sourcery.be/tileworld/src/tileworld"
)

// some globals
const (
	COLS = 40
	ROWS = 40
)

func main() {
	fmt.Println("start")
	gtk.Init(nil)
	fmt.Println("gtk inited")
	grid := tileworld.NewGrid(COLS, ROWS, 6, 20, 20, 20)
	fmt.Println("grid created")
	view := tileworld.GridViewNew(grid)
	fmt.Println("view created")
	view.Show()
	fmt.Println("view shown")
	gtk.Main()
}
