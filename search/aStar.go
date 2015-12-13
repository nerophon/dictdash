/*
Based on:
https://github.com/beefsack/go-astar/blob/master/astar.go
MIT Licence, to be reproduced in full if published
*/
package search

import "container/heap"

// the A* pathfinding algorithm

// GraphNode is an interface which allows A* searching on arbitrary objects
type GraphNode interface {

	// GraphNodes to which a direct no-hops path exists
	GetNeighbors() []GraphNode

	// the exact movement cost to a neighbouring GraphNode
	ActualNeighborCost(to GraphNode) int

	// heuristic that estimates the total travel cost to any GraphNode
	EstimatedTargetCost(to GraphNode) int
}

// a wrapper to store A* data for a GraphNode.
type aStarNode struct {
	graphNode GraphNode
	parent *aStarNode
	open bool
	closed bool
	cost int // actual cost of reaching this node from the start
	rank int // lower is better (i.e. 1 is better than 2)
	index int // in the priority queue
}

// a collection of aStarNodes keyed by GraphNodes
// needed for quick reference
type nodeMap map[GraphNode]*aStarNode

// get gets the GraphNode wrapped in an aStarNode,
// instantiating if required.
func (nm nodeMap) get(gn GraphNode) *aStarNode {
	node, ok := nm[gn]
	if !ok {
		node = &aStarNode{
			graphNode: gn,
		}
		nm[gn] = node
	}
	return node
}

// this search is a simple special case of single-pair shortest-path,
// insomuch as all the edges have cost = 1;
// the implementation here is A*;
// where the heuristic is the theoretic min transforms,
// ignoring whether or not the word is in the dictionary;
// this heuristic is admissable because it cannot overestimate
// 
// panics if graph is nil
// func Search(src string, dst string, graph WordGraph) {
// 	fmt.Println("Sorry, command not yet implemented.\n")
// }

// calculates the shortest path between two GraphNodes;
// if no path is found, found will be false
func Path(from, to GraphNode) (path []GraphNode, distance int, found bool) {
	nodeMap := nodeMap{}
	queue := &PriorityQueue{}
	heap.Init(queue)
	fromNode := nodeMap.get(from)
	fromNode.open = true
	heap.Push(queue, (*item)(fromNode))
	for {
		if queue.Len() == 0 {
			// there's no path, return found false
			return
		}
		// priority queue does its magic:
		current := (*aStarNode)(heap.Pop(queue).(*item))
		current.open = false
		current.closed = true

		if current == nodeMap.get(to) {
			// found a path to the goal!
			reversePath := []GraphNode{}
			curr := current
			for curr != nil {
				reversePath = append(reversePath, curr.graphNode)
				curr = curr.parent
			}
			// reverse path to get from --> to
			path = make([]GraphNode, len(reversePath))
			j := 0
			for i := len(reversePath)-1; i >= 0; i-- {
				path[j] = reversePath[i]
				j++
			}
			return path, current.cost, true
		}


		for _, neighbor := range current.graphNode.GetNeighbors() {
			cost := current.cost + current.graphNode.ActualNeighborCost(neighbor)
			neighborNode := nodeMap.get(neighbor)
			if cost < neighborNode.cost {
				// this branch is for handling nodes
				// that have already been assessed in previous iterations
				// where the new cost is cheaper
				if neighborNode.open {
					heap.Remove(queue, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.EstimatedTargetCost(to)
				neighborNode.parent = current
				heap.Push(queue, (*item)(neighborNode))
			}
		}
	}
}







