package grapher

import (
		"testing"
		"reflect"
)

func TestNodeToString(t *testing.T) {
	cases := []struct {
		in *WordNode
		want string
	}{
		{&WordNode{"", nil, nil,}, "\nWord: \nEdges: nil\nNeighbours: nil\n"},
		{&WordNode{"a", [][]*WordNode{[]*WordNode{},}, nil,}, "\nWord: a\nEdges: \n\t0:[]\nNeighbours: nil\n"},
		{&WordNode{"hi", [][]*WordNode{[]*WordNode{nil,nil,nil,}, []*WordNode{nil,nil,nil,},}, nil, }, "\nWord: hi\nEdges: \n\t0:[-,-,-,]\n\t1:[-,-,-,]\nNeighbours: nil\n"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("NodeToString(), want=%v, got=%v", c.want, got)
		}
	}	
}

func TestNewWordNode(t *testing.T) {
	//t.SkipNow()
	cases := []struct {
		inWord string
		inEdgeCount uint
		want *WordNode
	}{
		{"", 5, &WordNode{"", [][]*WordNode{}, nil,},},
		{"hi", 3, &WordNode{"hi", [][]*WordNode{[]*WordNode{nil,nil,nil,}, []*WordNode{nil,nil,nil,},}, nil,},},
		{"a", 0, &WordNode{"a", [][]*WordNode{[]*WordNode{},}, nil,},},
	}
	for _, c := range cases {
		got := NewWordNode(c.inWord, c.inEdgeCount)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("NewWordNode(%s, %d), expected=%v, actual=%v", c.inWord, c.inEdgeCount, c.want, got)
		}
	}
}

func TestEstimatedTargetCost(t *testing.T) {
	cases := []struct {
		from *WordNode
		to *WordNode
		want int
	}{
		{NewWordNode("", 0), NewWordNode("", 0), 0,},
		{NewWordNode("a", 0), NewWordNode("a", 0), 0,},
		{NewWordNode("a", 0), NewWordNode("b", 0), 1,},
		{NewWordNode("messenger", 0), NewWordNode("engineers", 0), 9,},
		{NewWordNode("ball", 0), NewWordNode("call", 0), 1,},
		{NewWordNode("ball", 0), NewWordNode("pale", 0), 2,},
	}
	for _, c := range cases {
		got := c.from.EstimatedTargetCost(c.to)
		if got != c.want {
			t.Errorf("TestEstimatedTargetCost(), from=%s, to=%s, want=%d, got=%d", c.from.Word, c.to.Word, c.want, got)
		}
	}	
}















