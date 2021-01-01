package tileworld

import (
	"math"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

const (
	// MAG the magniication
	MAG = float64(20.0)
)

// GridView wrapper around GTK window
type GridView struct {
	grid *Grid
	win  *gtk.Window
	da   *gtk.DrawingArea
}

// GridViewNew construct
func GridViewNew(grid *Grid) *GridView {
	view := new(GridView)
	view.grid = grid
	view.win, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	view.win.SetDefaultSize(int(grid.cols)*int(MAG), int(grid.rows)*int(MAG))
	view.da, _ = gtk.DrawingAreaNew()

	view.da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		view.drawGrid(cr)
	})
	view.win.Add(view.da)
	view.win.SetTitle("TileWorld")
	view.win.Connect("destroy", gtk.MainQuit)
	go onTimeout(grid, view)
	return view
}

func onTimeout(grid *Grid, view *GridView) bool {
	for {
		grid.Update()
		view.win.QueueDrawArea(0, 0, view.win.GetAllocatedWidth(), view.win.GetAllocatedHeight())
		time.Sleep((time.Second))
	}
}

func (v GridView) drawGrid(cr *cairo.Context) {
	for c := uint8(0); c < v.grid.cols; c++ {
		for r := uint8(0); r < v.grid.rows; r++ {
			o := v.grid.Object(NewLocation(c, r))
			if o != nil {
				drawObject(cr, o, float64(c)*MAG, float64(r)*MAG)
			}
		}
	}
}

func drawObject(cr *cairo.Context, o *GridObject, x, y float64) {
	cr.SetSourceRGB(0, 0, 0)
	cr.SetLineWidth(2)
	switch o.objectType {
	case TypeAgent:
		drawAgent(cr, o, x, y)
		break
	case TypeTile:
		drawTile(cr, o, x, y)
		break
	case TypeHole:
		drawHole(cr, o, x, y)
		break
	case TypeObstacle:
		drawObstacle(cr, o, x, y)
		break
	}
}

func drawAgent(cr *cairo.Context, o *GridObject, x, y float64) {
	cr.NewPath()
	cr.SetSourceRGB(setColor(o.num, cr))
	cr.Rectangle(x, y, MAG, MAG)
	cr.Stroke()
}

func drawTile(cr *cairo.Context, o *GridObject, x, y float64) {
	cr.NewPath()
	cr.Arc(x+MAG/2, y+MAG/2, MAG/2, 0, 2*math.Pi)
	cr.Stroke()
}

func drawHole(cr *cairo.Context, o *GridObject, x, y float64) {
	cr.NewPath()
	cr.Arc(x+MAG/2, y+MAG/2, MAG/2, 0, 2*math.Pi)
	cr.Fill()
}

func drawObstacle(cr *cairo.Context, o *GridObject, x, y float64) {
	cr.NewPath()
	cr.Rectangle(x, y, MAG, MAG)
	cr.Fill()
}

// Show show the window
func (v GridView) Show() {
	v.win.ShowAll()
}

func setColor(i uint8, cr *cairo.Context) (float64, float64, float64) {
	switch i {
	case 0:
		return 0, 0, 255
	case 1:
		return 255, 0, 0
	case 2:
		return 0, 255, 0
	case 3:
		return 128, 128, 0
	case 4:
		return 0, 128, 128
	case 5:
		return 128, 0, 128
	}
	return 0, 0, 0
}