/*
*/
package grapher

import (
		"strconv"
		"github.com/nerophon/dictdash/search"
)


// the Edges structure is: [letterIndex][replacementIndex]
// during a scan, the replacementIndex is fully allocated to a size of 26,
// and represents the index of the replacement letter in the alphabet;
// the Neighbours structure is a compressed version of edges,
// containing all edges in a flat list without gaps
type WordNode struct {
	Word string
	Edges [][]*WordNode
	Neighbours []search.GraphNode
}

func NewWordNode(word string, edgeCount uint) (node *WordNode) {
	node = new(WordNode)
	node.Word = word
	node.Edges = make([][]*WordNode, len(word))
	for k, _ := range node.Edges {
		node.Edges[k] = make([]*WordNode, edgeCount)
	}
	return
}

// for visual inspection of small dictionaries
// not suitable for production or large dictionaries
func (node *WordNode) String() string {
	result := "\nWord: " + node.Word + "\nEdges: "
	if node.Edges == nil {
		result = result + "nil"
	} else {
		for k, v := range node.Edges {
			result = result + "\n\t" + strconv.Itoa(k) + ":["
			for _, vv := range v {
				if vv == nil {
					result = result + "-,"
				} else {
					result = result + vv.Word + ","
				}
			}
			result = result + "]"
		}
	}
	result = result + "\nNeighbours: "
	if node.Neighbours == nil {
		result = result + "nil"
	} else {
		result = result + "["
		for _, v := range node.Neighbours {
			result = result + v.(*WordNode).Word + ","
		}
		result = result + "]"
	}
	result = result + "\n"
	return result
}

// Implementing the search.GraphNode interface in search/aStar.go:

// nodes to which a direct no-hops path exists
func (node *WordNode) GetNeighbors() []search.GraphNode {
	return node.Neighbours
}

// the exact movement cost to a neighbouring node
func (node *WordNode) ActualNeighborCost(to search.GraphNode) int {
	return 1
}

// heuristic that estimates the total travel cost to any node
// assumes equal length words
func (node *WordNode) EstimatedTargetCost(to search.GraphNode) int {
	differences := 0
	for i := 0; i < len(node.Word); i++ {
		if node.Word[i] != to.(*WordNode).Word[i] {
			differences++
		}
	}
	return differences
}














