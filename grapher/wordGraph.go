/*
*/
package grapher

import (
		"io"
		"bufio"
		"container/list"
		"github.com/nerophon/dictdash/search"
)

// strings in Go are read-only byte array slices
// indexing [] into them returns the byte
// this string is ideal because each alphabet letter is one byte
const alphabet string = "abcdefghijklmnopqrstuvwxyz"

// Go uses utf8, so to have small 'a' be zero you need to subtract 97
const decodeOffset byte = 97

// initial node edge slice allocation
const initialEdgeCount uint = 26

// because the long form is a mouthfull
type WordGraph map[int]map[string]*WordNode


// panics if input == nil
func ScanLinkCompress(input io.Reader) (graph WordGraph, count int, err error) {
	graph, count, err = scan(input)
	link(graph)
	compress(graph)
	return
}

// panics if input == nil
func scan(input io.Reader) (graph WordGraph, count int, err error) {
	graph = make(WordGraph)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if addToGraph(scanner.Text(), graph) {
			count++
		}
	}
	if err = scanner.Err(); err != nil {
		graph = nil
		return
	}
	return
}

// panics if graph == nil
func addToGraph(word string, graph WordGraph) (added bool) {
	letterCount := len(word)
	if _, ok := graph[letterCount]; !ok {
		graph[letterCount] = make(map[string]*WordNode)
	}
	if _, ok := graph[letterCount][word]; ok {
		return false;
	}
	graph[letterCount][word] = NewWordNode(word, initialEdgeCount)
	return true;
}

// panics if graph == nil
func link(graph WordGraph) {
	// opportunity for concurrency here
	for letterCount, subGraph := range graph {
		for word, node := range subGraph {
			for letter := 0; letter < letterCount; letter++ {
				for replacement := 0; replacement < 26; replacement++ {
					searchWord := replaceAtIndex(word, alphabet[replacement], letter)
					if searchWord == word {
						continue // self should remain nil
					}
					if node.Edges[letter][replacement] != nil {
						continue // mirror, no need to lookup
					}
					if connectedNode, ok := subGraph[searchWord]; ok {
						original := word[letter]
						originalZeroIndexed := original - decodeOffset
						node.Edges[letter][replacement] = connectedNode
						connectedNode.Edges[letter][originalZeroIndexed] = node
					}
				}
			}
		}
	}
	return
}

// panics if i > len(in)
func replaceAtIndex(in string, b byte, i int) string {
    out := []byte(in)
    out[i] = b
    return string(out)
}

// panics if graph == nil
func compress(graph WordGraph) {
	// opportunity for concurrency here
	for letterCount, subGraph := range graph {
		for _, node := range subGraph {
			list := list.New()
			for letter := 0; letter < letterCount; letter++ {
				for _, edge := range node.Edges[letter] {
					if edge != nil {
						list.PushBack(edge)
					}
				}
			}
			node.Neighbours = listToNodeSlice(list)
			node.Edges = nil // TODO check if memory leak
		}
	}
	return
}

// panics if l is null or contains non-*WordNode values
func listToNodeSlice(l *list.List) (s []search.GraphNode) {
	s = make([]search.GraphNode, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		s[i] = e.Value.(search.GraphNode)
		i++
	}
	return
}












