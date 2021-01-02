package tileworld

import (
	"github.com/emirpasic/gods/maps/treemap"
)

// GetPath returns a path from the start to the destination
func GetPath(grid *Grid, from, to *Location) []Direction {
	list := []*Location{from}
	queue := treemap.NewWithIntComparator()
	queue.Put(0, list)
	for !queue.Empty() {
		lowest, value := queue.Min()
		path := value.([]*Location)
		queue.Remove(lowest)
		lastLocation := path[len(path)-1]
		if lastLocation.Equals(to) {
			// arrived
			return makePath(path)
		}
		generateNext(grid, to, path, queue, Up)
		generateNext(grid, to, path, queue, Down)
		generateNext(grid, to, path, queue, Left)
		generateNext(grid, to, path, queue, Right)
	}
	return []Direction{}
}

func generateNext(g *Grid, to *Location, path []*Location, queue *treemap.Map, dir Direction) {
	last := path[len(path)-1]
	nextLocation := last.NextLocation(dir)
	if nextLocation.Equals(to) || g.isValidLocation(nextLocation) {
		newPath := []*Location{}
		newPath = append(newPath, path...)
		if !hasLoop(newPath, nextLocation) {
			newPath = append(newPath, nextLocation)
			cost := len(newPath) + to.Distance(nextLocation)
			queue.Put(cost, newPath)
		}
	}
}

func hasLoop(path []*Location, nextLocation *Location) bool {
	for _, l := range path {
		if l.Equals(nextLocation) {
			return true
		}
	}
	return false
}

func makePath(locationList []*Location) []Direction {
	path := []Direction{}
	last := locationList[0]
	fromOne := locationList[1:]
	for _, l := range fromOne {
		dir := last.GetDirection(l)
		path = append(path, dir)
		last = l
	}
	return path
}
