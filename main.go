package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"sourcery.be/tileworld/src/tileworld"
)

// some globals
const (
	COLS = 40
	ROWS = 40
)

func main() {
	rand.Seed(time.Now().UnixNano())
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
