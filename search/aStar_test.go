package search

import (
		"testing"
		"reflect"
)

type TestNode struct {
	Name string
	Neighbours []GraphNode
}
func NewTestNode(name string) (node *TestNode) {
	node = new(TestNode)
	node.Name = name
	return
}
func (node *TestNode) String() string {
	return node.Name
}
func (node *TestNode) GetNeighbors() []GraphNode {
	return node.Neighbours
}
func (node *TestNode) ActualNeighborCost(to GraphNode) int {
	return 1
}
func (node *TestNode) EstimatedTargetCost(to GraphNode) int {
	differences := 0
	for i := 0; i < len(node.Name); i++ {
		if node.Name[i] != to.(*TestNode).Name[i] {
			differences++
		}
	}
	return differences
}


func TestPath(t *testing.T) {
	nodeA := GraphNode(NewTestNode("game"))
	nodeB := GraphNode(NewTestNode("lame"))
	nodeC := GraphNode(NewTestNode("tame"))
	nodeD := GraphNode(NewTestNode("lake"))
	nodeE := GraphNode(NewTestNode("tape"))
	nodeF := GraphNode(NewTestNode("gape"))
	nodeG := GraphNode(NewTestNode("test"))
	nodeA.(*TestNode).Neighbours = []GraphNode{nodeB,nodeC,nodeF}
	nodeB.(*TestNode).Neighbours = []GraphNode{nodeA,nodeC,nodeD}
	nodeC.(*TestNode).Neighbours = []GraphNode{nodeA,nodeB,nodeE}
	nodeD.(*TestNode).Neighbours = []GraphNode{nodeB}
	nodeE.(*TestNode).Neighbours = []GraphNode{nodeC,nodeF}
	nodeF.(*TestNode).Neighbours = []GraphNode{nodeA,nodeE}
	nodeG.(*TestNode).Neighbours = []GraphNode{}
	cases := []struct {
		from, to GraphNode
		pathWant []GraphNode
		distWant int
		foundWant bool
	}{
		// test no path
		{nodeA, nodeG, nil, 0, false},
		// test zero jump path
		{nodeA, nodeA, []GraphNode{nodeA}, 0, true},
		// test one jump path
		{nodeA, nodeB, []GraphNode{nodeA,nodeB}, 1, true},
		// test two jump path
		{nodeA, nodeD, []GraphNode{nodeA,nodeB,nodeD}, 2, true},
		// test two equal cost paths
		// we expect returned path to follow the add order (as per priority queue test),
		// which is A -> C rather than A -> F,
		// because C is added to the Neighbours slice before F (above)
		{nodeA, nodeE, []GraphNode{nodeA,nodeC,nodeE}, 2, true},
	}
	for _, c := range cases {
		pathGot, distGot, foundGot := Path(c.from, c.to)
		if !reflect.DeepEqual(c.pathWant, pathGot) {
			t.Errorf("TestPath: pathWant=%v, pathGot=%v", c.pathWant, pathGot)
		} 
		if c.distWant != distGot {
			t.Errorf("TestPath: distWant=%d, distGot=%d", c.distWant, distGot)
		} 
		if c.foundWant != foundGot {
			t.Errorf("TestPath: foundWant=%b, foundGot=%b", c.foundWant, foundGot)
		} 
	}	
}














