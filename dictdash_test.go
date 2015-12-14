
package main

import (
		"testing"
		"os"
		"github.com/nerophon/dictdash/grapher"
		"github.com/nerophon/dictdash/search"
)

var benchGraph grapher.WordGraph
var benchFound bool

func TestSanity(t *testing.T) {
	if (1 != 1) {
		t.Errorf("TestSanity(), cosmic rays changing bits in yo RAM!")
	}
}

func BenchmarkOpen(b *testing.B) {
    for n := 0; n < b.N; n++ {
		file, _ := os.Open("dict.txt")
		file.Close()
    }
}

func BenchmarkScan(b *testing.B) {
    // run the Scan b.N times
    for n := 0; n < b.N; n++ {
		file, _ := os.Open("dict.txt")
		graphRes, _, _ := grapher.ScanLinkCompress(file)
		benchGraph = graphRes // to prevent compiler skip
		file.Close()
    }
}

func BenchmarkSearch(b *testing.B) {
	// setup
	file, _ := os.Open("dict.txt")
	defer file.Close()
	graphRes, _, _ := grapher.ScanLinkCompress(file)
	benchGraph = graphRes // to prevent compiler skip

	srcNode, _ := graph[6]["bounce"]
	dstNode, _ := graph[6]["lather"]
	src := search.GraphNode(srcNode)
	dst := search.GraphNode(dstNode)
	b.ResetTimer()

    // run the Search b.N times
    for n := 0; n < b.N; n++ {
		_, _, found := search.Path(src, dst)
		benchFound = found // to prevent compiler skip
    }
}












