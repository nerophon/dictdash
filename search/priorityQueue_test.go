package search

import (
		"testing"
		"container/heap"
)

func TestPushPopPriority(t *testing.T) {
	queue := &PriorityQueue{}
	heap.Init(queue)
	cases := []struct {
		in []*item
	}{
		{[]*item{&item{rank: 1}, &item{rank: 2}, &item{rank: 3},},},
		{[]*item{&item{rank: 1}, &item{rank: 3}, &item{rank: 2},},},
		{[]*item{&item{rank: 2}, &item{rank: 1}, &item{rank: 3},},},
		{[]*item{&item{rank: 2}, &item{rank: 3}, &item{rank: 1},},},
		{[]*item{&item{rank: 3}, &item{rank: 1}, &item{rank: 2},},},
		{[]*item{&item{rank: 3}, &item{rank: 2}, &item{rank: 1},},},
	}
	for _, c := range cases {
		for _, item := range c.in {
			heap.Push(queue, item)
		}
		wantRank := 0
		for queue.Len() > 0 {
			wantRank++
			got := heap.Pop(queue).(*item)
			if wantRank != got.rank {
				t.Errorf("TestPushPopPriority: wantRank=%d, got.Rank=%d", wantRank, got.rank)
			}
		}
	}

	// so what happens if priorities are equal?
	nodeA := &aStarNode{cost: 2, rank: 1}
	nodeB := &aStarNode{cost: 1, rank: 1}
	heap.Push(queue, (*item)(nodeA))
	heap.Push(queue, (*item)(nodeB))
	nodeOut := (*aStarNode)(heap.Pop(queue).(*item))
	if nodeOut.cost != 2 {
		t.Errorf("TestPushPopPriority: equal rank items not popped by add order")
	}
}














