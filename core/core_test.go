package core

import (
		"testing"
		"strings"
		"reflect"
)

func TestNodeToString(t *testing.T) {
	cases := []struct {
		in *Node
		want string
	}{
		{&Node{"", [][]*Node{},}, "\nWord: \n"},
		{&Node{"a", [][]*Node{[]*Node{},},}, "\nWord: a\n\t0:[]\n"},
		{&Node{"hi", [][]*Node{[]*Node{nil,nil,nil,}, []*Node{nil,nil,nil,},},}, "\nWord: hi\n\t0:[-,-,-,]\n\t1:[-,-,-,]\n"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("NodeToString(), want=%v, got=%v", c.want, got)
		}
	}	
}

func TestNewNode(t *testing.T) {
	//t.SkipNow()
	cases := []struct {
		inWord string
		inEdgeCount uint
		want *Node
	}{
		{"", 5, &Node{"", [][]*Node{},},},
		{"hi", 3, &Node{"hi", [][]*Node{[]*Node{nil,nil,nil,}, []*Node{nil,nil,nil,},},},},
		{"a", 0, &Node{"a", [][]*Node{[]*Node{},},},},
	}
	for _, c := range cases {
		got := NewNode(c.inWord, c.inEdgeCount)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("NewNode(%s, %d), expected=%v, actual=%v", c.inWord, c.inEdgeCount, c.want, got)
		}
	}
}

func TestAddToDictionary(t *testing.T) {
	cases := []struct {
		inWord string
		expAdded bool
		inDict map[int]map[string]*Node
		expDict map[int]map[string]*Node
	}{
		//test empty input
		{"", true,
			map[int]map[string]*Node{},
			map[int]map[string]*Node{
			0: map[string]*Node{
				"": NewNode("", initialEdgeCount),
				},
			},
		},
		//test short word
		{"short", true,
			map[int]map[string]*Node{},
			map[int]map[string]*Node{
			5: map[string]*Node{
				"short": NewNode("short", initialEdgeCount),
				},
			},
		},
		//test short word with existing different word
		{"short", true,
			map[int]map[string]*Node{
			4: map[string]*Node{
				"tall": NewNode("tall", initialEdgeCount),
				},
			},
			map[int]map[string]*Node{
			4: map[string]*Node{
				"tall": NewNode("tall", initialEdgeCount),
				},
			5: map[string]*Node{
				"short": NewNode("short", initialEdgeCount),
				},
			},
		},
		//test short word with existing identical word
		{"short", false,
			map[int]map[string]*Node{
			5: map[string]*Node{
				"short": NewNode("short", initialEdgeCount),
				},
			},
			map[int]map[string]*Node{
			5: map[string]*Node{
				"short": NewNode("short", initialEdgeCount),
				},
			},
		},
	}
	for _, c := range cases {
		added := addToDictionary(c.inWord, c.inDict)
		if added != c.expAdded {
			t.Errorf("addToDictionary(%s), expAdded=%b, added=%b", c.inWord, c.expAdded, added)
		}
	}	
}

func TestScan(t *testing.T) {
	//t.SkipNow()
	cases := []struct {
		in string
		expErr error
		expCount int
		expDict map[int]map[string]*Node
	}{
		//test empty input
		{"", nil, 0, map[int]map[string]*Node{},
		},
		//test basic input
		{"how are you", nil, 3, map[int]map[string]*Node{
			3: map[string]*Node{
				"how": NewNode("how", initialEdgeCount),
				"are": NewNode("are", initialEdgeCount),
				"you": NewNode("you", initialEdgeCount),
				},
			},
		},
		//test capitalized input
		{"how are YOU", nil, 3, map[int]map[string]*Node{
			3: map[string]*Node{
				"how": NewNode("how", initialEdgeCount),
				"are": NewNode("are", initialEdgeCount),
				"YOU": NewNode("YOU", initialEdgeCount),
				},
			},
		},
		//test multiple word lengths and repeated words
		{"thou art more lovely and more temperate", nil, 6, map[int]map[string]*Node{
			3: map[string]*Node{
				"art": NewNode("art", initialEdgeCount),
				"and": NewNode("and", initialEdgeCount),
				},
			4: map[string]*Node{
				"thou": NewNode("thou", initialEdgeCount),
				"more": NewNode("more", initialEdgeCount),
				},
			6: map[string]*Node{
				"lovely": NewNode("lovely", initialEdgeCount),
				},
			9: map[string]*Node{
				"temperate": NewNode("temperate", initialEdgeCount),
				},
			},
		},
	}
	for _, c := range cases {
		input := strings.NewReader(c.in)
		dict, count, err := scan(input)
		if err != c.expErr {
			t.Errorf("Scan(%v), expected=%v, actual=%v", c.in, c.expErr, err)
		}
		if count != c.expCount {
			t.Errorf("Scan(%d), expected=%d, actual=%d", c.in, c.expCount, count)
		}
		if !reflect.DeepEqual(dict, c.expDict) {
			t.Errorf("Scan(%v), expected=%v, actual=%v", c.in, c.expDict, dict)
		}
	}
}

func TestReplaceAtIndex(t *testing.T) {
	cases := []struct {
		in string
		b byte
		i int
		want string
	}{
		{"cold", alphabet[1], 0, "bold"},
		//{"mold", alphabet[24], 5, "moldy"}, //panic!
	}
	for _, c := range cases {
		got := replaceAtIndex(c.in, c.b, c.i)
		if got != c.want {
			t.Errorf("replaceAtIndex(%s, %s, %d), want=%s, got=%s", c.in, []byte{c.b}, c.i, c.want, got)
		}
	}	
}

func TestGraph(t *testing.T) {
	cases := []struct {
		in map[int]map[string]*Node
		want map[int]map[string]*Node
	}{
		//test basic input
		{map[int]map[string]*Node{
			3: map[string]*Node{
				"hit": NewNode("hit", initialEdgeCount),
				"hat": NewNode("hat", initialEdgeCount),
				"cat": NewNode("cat", initialEdgeCount),
			},
		},
		map[int]map[string]*Node{
			3: map[string]*Node{
				"hit": NewNode("hit", initialEdgeCount),
				"hat": NewNode("hat", initialEdgeCount),
				"cat": NewNode("cat", initialEdgeCount),
			},
		},},
	}
	cases[0].want[3]["hit"].Edges[1][0] = cases[0].want[3]["hat"]
	cases[0].want[3]["hat"].Edges[1][8] = cases[0].want[3]["hit"]
	cases[0].want[3]["hat"].Edges[0][2] = cases[0].want[3]["cat"]
	cases[0].want[3]["cat"].Edges[0][7] = cases[0].want[3]["hat"]
	for _, c := range cases {
		graph(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("graph(), want=%v, got=%v", c.want, c.in)
		}
	}
}




