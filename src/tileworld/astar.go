package tileworld

import (
	"container/heap"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/utils"
)

func comp(a, b interface{}) int {
	aK := a.(*Node)
	bK := b.(*Node)
	return utils.IntComparator(aK.priority, bK.priority)
}

func GetPathAStar(grid *Grid, from, to *Location) []Direction {
	openList := make(PriorityQueue, 10)
	closedList := treeset.NewWith(comp)
	fromNode := new(Node)
	fromNode.location = from
	fromNode.parent = nil
	fromNode.priority = 0
	heap.Push(&openList, fromNode)
	heap.Init(&openList)
	for openList.Len() > 0 {
		current := heap.Pop(&openList).(*Node)
		if current.location == to {
			// arrived
			return makePathFromNode(current, from)
		}
		closedList.Add(current)
		checkNeighbor(grid, current, openList, closedList, Up, from, to)
		checkNeighbor(grid, current, openList, closedList, Down, from, to)
		checkNeighbor(grid, current, openList, closedList, Right, from, to)
		checkNeighbor(grid, current, openList, closedList, Left, from, to)
	}
	return []Direction{}
}

func checkNeighbor(grid *Grid, current *Node, openList PriorityQueue, closedList *treeset.Set, dir Direction, from, to *Location) {
	nextLoc := current.location.NextLocation(dir)
	if nextLoc == to || grid.isValidLocation(nextLoc) {
		h := nextLoc.Distance(to)
		g := current.location.Distance(from) + 1
		child := &Node{
			parent:   current,
			priority: g + h,
			location: nextLoc,
		}
		if !closedList.Contains(child) {
			heap.Push(&openList, child)
		}
	}
}

func makePathFromNode(end *Node, from *Location) []Direction {
	directions := []Direction{}
	current := end
	parent := end.parent
	for current.location != from {
		d := moveFromParent(current.location, parent.location)
		directions = append(directions, d)
	}
	return directions
}

func moveFromParent(current, parent *Location) Direction {
	if current.col == parent.col {
		if parent.row == current.row-1 {
			return Down
		} else {
			return Up
		}
	} else {
		if parent.col == current.col-1 {
			return Right
		} else {
			return Left
		}
	}
}
