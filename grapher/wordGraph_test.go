package grapher

import (
		"testing"
		"strings"
		"reflect"
		"container/list"
		"github.com/nerophon/dictdash/search"
)

func TestAddToGraph(t *testing.T) {
	//t.SkipNow()
	cases := []struct {
		inWord string
		expAdded bool
		inGraph WordGraph
		expGraph WordGraph
	}{
		//test empty input
		{"", true,
			WordGraph{},
			WordGraph{
			0: map[string]*WordNode{
				"": NewWordNode("", initialEdgeCount),
				},
			},
		},
		//test short word
		{"short", true,
			WordGraph{},
			WordGraph{
			5: map[string]*WordNode{
				"short": NewWordNode("short", initialEdgeCount),
				},
			},
		},
		//test short word with existing different word
		{"short", true,
			WordGraph{
			4: map[string]*WordNode{
				"tall": NewWordNode("tall", initialEdgeCount),
				},
			},
			WordGraph{
			4: map[string]*WordNode{
				"tall": NewWordNode("tall", initialEdgeCount),
				},
			5: map[string]*WordNode{
				"short": NewWordNode("short", initialEdgeCount),
				},
			},
		},
		//test short word with existing identical word
		{"short", false,
			WordGraph{
			5: map[string]*WordNode{
				"short": NewWordNode("short", initialEdgeCount),
				},
			},
			WordGraph{
			5: map[string]*WordNode{
				"short": NewWordNode("short", initialEdgeCount),
				},
			},
		},
	}
	for _, c := range cases {
		added := addToGraph(c.inWord, c.inGraph)
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
		expGraph WordGraph
	}{
		//test empty input
		{"", nil, 0, WordGraph{},
		},
		//test basic input
		{"how are you", nil, 3, WordGraph{
			3: map[string]*WordNode{
				"how": NewWordNode("how", initialEdgeCount),
				"are": NewWordNode("are", initialEdgeCount),
				"you": NewWordNode("you", initialEdgeCount),
				},
			},
		},
		//test capitalized input
		{"how are YOU", nil, 3, WordGraph{
			3: map[string]*WordNode{
				"how": NewWordNode("how", initialEdgeCount),
				"are": NewWordNode("are", initialEdgeCount),
				"YOU": NewWordNode("YOU", initialEdgeCount),
				},
			},
		},
		//test multiple word lengths and repeated words
		{"thou art more lovely and more temperate", nil, 6, WordGraph{
			3: map[string]*WordNode{
				"art": NewWordNode("art", initialEdgeCount),
				"and": NewWordNode("and", initialEdgeCount),
				},
			4: map[string]*WordNode{
				"thou": NewWordNode("thou", initialEdgeCount),
				"more": NewWordNode("more", initialEdgeCount),
				},
			6: map[string]*WordNode{
				"lovely": NewWordNode("lovely", initialEdgeCount),
				},
			9: map[string]*WordNode{
				"temperate": NewWordNode("temperate", initialEdgeCount),
				},
			},
		},
	}
	for _, c := range cases {
		input := strings.NewReader(c.in)
		graph, count, err := scan(input)
		if err != c.expErr {
			t.Errorf("Scan(%v), expected=%v, actual=%v", c.in, c.expErr, err)
		}
		if count != c.expCount {
			t.Errorf("Scan(%d), expected=%d, actual=%d", c.in, c.expCount, count)
		}
		if !reflect.DeepEqual(graph, c.expGraph) {
			t.Errorf("Scan(%v), expected=%v, actual=%v", c.in, c.expGraph, graph)
		}
	}
}

func TestReplaceAtIndex(t *testing.T) {
	//t.SkipNow()
	cases := []struct {
		in string
		b byte
		i int
		want string
	}{
		{"cold", alphabet[1], 0, "bold"},
		//{"mess", alphabet[24], 5, "messy"}, //panic!
	}
	for _, c := range cases {
		got := replaceAtIndex(c.in, c.b, c.i)
		if got != c.want {
			t.Errorf("replaceAtIndex(%s, %s, %d), want=%s, got=%s", c.in, []byte{c.b}, c.i, c.want, got)
		}
	}	
}

func TestLink(t *testing.T) {
	//t.SkipNow()
	cases := []struct {
		in WordGraph
		want WordGraph
	}{
		//test empty input
		{
			WordGraph{},
			WordGraph{},
		},
		//test basic input
		{
			WordGraph{
				3: map[string]*WordNode{
					"hit": NewWordNode("hit", initialEdgeCount),
					"hat": NewWordNode("hat", initialEdgeCount),
					"cat": NewWordNode("cat", initialEdgeCount),
				},
			},
			WordGraph{
				3: map[string]*WordNode{
					"hit": NewWordNode("hit", initialEdgeCount),
					"hat": NewWordNode("hat", initialEdgeCount),
					"cat": NewWordNode("cat", initialEdgeCount),
				},
			},
		},
		//test multiple letter counts
		{
			WordGraph{
				3: map[string]*WordNode{
					"hit": NewWordNode("hit", initialEdgeCount),
					"hat": NewWordNode("hat", initialEdgeCount),
					"cat": NewWordNode("cat", initialEdgeCount),
				},
				4: map[string]*WordNode{
					"hate": NewWordNode("hate", initialEdgeCount),
					"cart": NewWordNode("cart", initialEdgeCount),
					"part": NewWordNode("part", initialEdgeCount),
				},
			},
			WordGraph{
				3: map[string]*WordNode{
					"hit": NewWordNode("hit", initialEdgeCount),
					"hat": NewWordNode("hat", initialEdgeCount),
					"cat": NewWordNode("cat", initialEdgeCount),
				},
				4: map[string]*WordNode{
					"hate": NewWordNode("hate", initialEdgeCount),
					"cart": NewWordNode("cart", initialEdgeCount),
					"part": NewWordNode("part", initialEdgeCount),
				},
			},
		},
	}
	// can't use a literal to assign a pointer conveniently,
	// so let's just manually add the expected associations:
	cases[1].want[3]["hit"].Edges[1][0] = cases[1].want[3]["hat"]
	cases[1].want[3]["hat"].Edges[1][8] = cases[1].want[3]["hit"]
	cases[1].want[3]["hat"].Edges[0][2] = cases[1].want[3]["cat"]
	cases[1].want[3]["cat"].Edges[0][7] = cases[1].want[3]["hat"]

	cases[2].want[3]["hit"].Edges[1][0] = cases[2].want[3]["hat"]
	cases[2].want[3]["hat"].Edges[1][8] = cases[2].want[3]["hit"]
	cases[2].want[3]["hat"].Edges[0][2] = cases[2].want[3]["cat"]
	cases[2].want[3]["cat"].Edges[0][7] = cases[2].want[3]["hat"]
	cases[2].want[4]["part"].Edges[0][2] = cases[2].want[4]["cart"]
	cases[2].want[4]["cart"].Edges[0][15] = cases[2].want[4]["part"]

	for _, c := range cases {
		link(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("graph(), want=%v, got=%v", c.want, c.in)
		}
		//t.Logf("graph(), want=%v, got=%v", c.want, c.in)
	}
}

func TestListToNodeSlice(t *testing.T) {
	//t.SkipNow()
	nodeA := search.GraphNode(NewWordNode("aardvark", 0))
	nodeB := search.GraphNode(NewWordNode("bonobo", 0))
	nodeC := search.GraphNode(NewWordNode("chipmunk", 0))
	cases := []struct {
		in *list.List
		want []search.GraphNode
	}{
		{list.New(), []search.GraphNode{},},
		{list.New(), []search.GraphNode{nodeA, nodeB, nodeC},},
		//{list.New(), []search.GraphNode{nodeC, nodeB, nodeA},}, order matters
	}
	cases[1].in.PushBack(nodeA)
	cases[1].in.PushBack(nodeB)
	cases[1].in.PushBack(nodeC)
	// cases[2].in.PushBack(nodeA)
	// cases[2].in.PushBack(nodeB)
	// cases[2].in.PushBack(nodeC)
	for _, c := range cases {
		got := listToNodeSlice(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("listToNodeSlice(%v), want=%v, got=%v", c.in, c.want, got)
		}
	}
}


func TestCompress(t *testing.T) {
	//t.SkipNow()
	cases := []struct {
		in WordGraph
		want WordGraph
	}{
		//test empty input
		{
			WordGraph{},
			WordGraph{},
		},
		//test sparse input
		{
			WordGraph{
				3: map[string]*WordNode{
					"hit": NewWordNode("hit", initialEdgeCount),
					"hat": NewWordNode("hat", initialEdgeCount),
					"hot": NewWordNode("hot", initialEdgeCount),
				},
				4: map[string]*WordNode{
					"hate": NewWordNode("hate", initialEdgeCount),
					"cart": NewWordNode("cart", initialEdgeCount),
					"part": NewWordNode("part", initialEdgeCount),
				},
			},
			WordGraph{
				3: map[string]*WordNode{
					"hit": NewWordNode("hit", 0),
					"hat": NewWordNode("hat", 0),
					"hot": NewWordNode("hot", 0),
				},
				4: map[string]*WordNode{
					"hate": NewWordNode("hate", 0),
					"cart": NewWordNode("cart", 0),
					"part": NewWordNode("part", 0),
				},
			},
		},
	}
	// can't use a literal to assign a pointer conveniently,
	// so let's just manually add the associations.

	// here we add the initial edges, as if link() had been run
	cases[1].in[3]["hit"].Edges[1][0] = cases[1].in[3]["hat"]
	cases[1].in[3]["hit"].Edges[1][14] = cases[1].in[3]["hot"]
	cases[1].in[3]["hat"].Edges[1][8] = cases[1].in[3]["hit"]
	cases[1].in[3]["hat"].Edges[1][14] = cases[1].in[3]["hot"]
	cases[1].in[3]["hot"].Edges[1][0] = cases[1].in[3]["hat"]
	cases[1].in[3]["hot"].Edges[1][8] = cases[1].in[3]["hit"]
	cases[1].in[4]["part"].Edges[0][2] = cases[1].in[4]["cart"]
	cases[1].in[4]["cart"].Edges[0][15] = cases[1].in[4]["part"]

	// next we nil the sparse edges
	cases[1].want[3]["hit"].Edges = nil
	cases[1].want[3]["hat"].Edges = nil
	cases[1].want[3]["hot"].Edges = nil
	cases[1].want[4]["hate"].Edges = nil
	cases[1].want[4]["part"].Edges = nil
	cases[1].want[4]["cart"].Edges = nil

	// now we initialize the compressed slices to the expected size
	cases[1].want[3]["hit"].Neighbours = make([]search.GraphNode, 2)
	cases[1].want[3]["hat"].Neighbours = make([]search.GraphNode, 2)
	cases[1].want[3]["hot"].Neighbours = make([]search.GraphNode, 2)
	cases[1].want[4]["hate"].Neighbours = make([]search.GraphNode, 0)
	cases[1].want[4]["part"].Neighbours = make([]search.GraphNode, 1)
	cases[1].want[4]["cart"].Neighbours = make([]search.GraphNode, 1)
	
	// finally we link the nodes in the expected way
	cases[1].want[3]["hit"].Neighbours[0] = search.GraphNode(cases[1].want[3]["hat"])
	cases[1].want[3]["hit"].Neighbours[1] = search.GraphNode(cases[1].want[3]["hot"])
	cases[1].want[3]["hat"].Neighbours[0] = search.GraphNode(cases[1].want[3]["hit"])
	cases[1].want[3]["hat"].Neighbours[1] = search.GraphNode(cases[1].want[3]["hot"])
	cases[1].want[3]["hot"].Neighbours[0] = search.GraphNode(cases[1].want[3]["hat"])
	cases[1].want[3]["hot"].Neighbours[1] = search.GraphNode(cases[1].want[3]["hit"])
	cases[1].want[4]["part"].Neighbours[0] = search.GraphNode(cases[1].want[4]["cart"])
	cases[1].want[4]["cart"].Neighbours[0] = search.GraphNode(cases[1].want[4]["part"])

	for _, c := range cases {
		compress(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("compress(), want=%v, got=%v", c.want, c.in)
		}
		//t.Logf("compress(), want=%v, got=%v", c.want, c.in)
	}
}














