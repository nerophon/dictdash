/*
Based on the example at:
https://golang.org/pkg/container/heap/
*/
package search

// An item is something we manage in a priority queue.
// type item struct {
// 	value interface{}
// 	rank int
// 	index int
// }
// For this usage we'll substitute the aStarNode (see aStar.go):
type item aStarNode

// a PriorityQueue implements heap.Interface and holds items.
type PriorityQueue []*item

func (pq PriorityQueue) Len() int { return len(pq) }

// we want Pop() to give us the highest priority item first
// we use rank, so a lower number means a higher priority
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].rank < pq[j].rank
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// modifies the priority of an item in the queue
// not needed by A*
// func (pq *PriorityQueue) update(item *item, rank int) {
// 	item.rank = rank
// 	heap.Fix(pq, item.index)
// }







