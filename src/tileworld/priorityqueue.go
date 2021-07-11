package tileworld

// An Node is something we manage in a priority queue.
type Node struct {
	location *Location  // The value of the Node; arbitrary.
	path     []Location // array of locations
	priority int        // The priority of the Node in the queue. (f in A* algorithm)
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the Node in the heap.
}

// A PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	Node := x.(*Node)
	Node.index = n
	*pq = append(*pq, Node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	Node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	Node.index = -1 // for safety
	*pq = old[0 : n-1]
	return Node
}
