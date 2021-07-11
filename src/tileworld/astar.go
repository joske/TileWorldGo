package tileworld

import (
	"container/heap"

	"github.com/emirpasic/gods/sets/treeset"
)

func comp(a, b interface{}) int {
	aK := a.(*Node)
	bK := b.(*Node)
	return aK.location.Distance(bK.location)
}

func GetPathAStar(grid *Grid, from, to *Location) []Location {
	openList := make(PriorityQueue, 0)
	closedList := treeset.NewWith(comp)
	fromNode := new(Node)
	fromNode.location = from
	fromNode.path = make([]Location, 0)
	fromNode.priority = 0
	heap.Push(&openList, fromNode)
	heap.Init(&openList)
	for openList.Len() > 0 {
		current := heap.Pop(&openList).(*Node)
		if current.location.Equals(to) {
			// arrived
			return current.path
		}
		closedList.Add(current)
		checkNeighbor(grid, current, &openList, closedList, Up, from, to)
		checkNeighbor(grid, current, &openList, closedList, Down, from, to)
		checkNeighbor(grid, current, &openList, closedList, Right, from, to)
		checkNeighbor(grid, current, &openList, closedList, Left, from, to)
	}
	return []Location{}
}

func checkNeighbor(grid *Grid, current *Node, openList *PriorityQueue, closedList *treeset.Set, dir Direction, from, to *Location) {
	nextLoc := current.location.NextLocation(dir)
	if nextLoc.Equals(to) || grid.isValidLocation(nextLoc) {
		h := nextLoc.Distance(to)
		g := len(current.path) + 1
		child := &Node{
			path:     current.path,
			priority: g + h,
			location: nextLoc,
		}
		child.path = append(child.path, *nextLoc)
		if !closedList.Contains(child) {
			heap.Push(openList, child)
		}
	}
}
