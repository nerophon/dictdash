/*
For a given dictionary of words, a source word, and a destination word,
dictdash computes the minimum number of single-letter transformations required
to go from source to destination via other words in the dictionary.
*/
package main

import (
		"fmt"
		"strings"
		"bufio"
		"os"
		"github.com/nerophon/dictdash/grapher"
		"github.com/nerophon/dictdash/search"
)


var graph map[int]map[string]*grapher.WordNode

func main() {
	fmt.Print("\nWelcome to Dictionary Dash!\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a command.\n" +
		"Typing \"help\" will show a command list.\n\n")
	commandLoop(reader)
}

func commandLoop(reader *bufio.Reader) {
	fmt.Print("> ")
	text, error := reader.ReadString('\n')
	if error != nil || len(text) == 0 {
		fmt.Println("Sorry, there was an input error:\n%d", error)
		return
	}
	trimmed := strings.Trim(text, "\n ")
	fields := strings.Fields(trimmed)
	numFields := len(fields)
	if numFields <= 0 {
		commandLoop(reader)
		return
	}
	switch fields[0] {
	case "quit":
		fmt.Println("Quitting.\n")
		return
	case "help":
		fmt.Println("\n***Command List***\n")
		fmt.Println("help\t\tshows this command list")
		fmt.Println("scan\t\tscans a whitespace-delimited dictionary at ./dict.txt")
		fmt.Println("scan [path]\tscans a whitespace-delimited dictionary at [path]")
		fmt.Println("search [A] [B]\tsearches for the shortest path from [A] to [B]")
		fmt.Println("quit\t\tquits the application")
		fmt.Println("")
	case "scan":
		if numFields == 1 {
			scan("dict.txt")
		} else if numFields == 2 {
			scan(fields[1])
		} else {
			fmt.Println("Please specify only one path.\n")
		}
	case "search":
		if numFields != 3 {
			fmt.Println("The search command requires exactly two arguments.\n")
		} else {
			searchCmd(fields[1], fields[2])
		}
	default:
		fmt.Println("Sorry, command not understood.\n")
	}
	commandLoop(reader)
}

func scan(path string) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var count int
	graph, count, err = grapher.ScanLinkCompress(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input failed with error:\n", err)
		return
	}
	fmt.Printf("Words scanned: %d\n", count)
	fmt.Printf("Sub-graph count: %d\n", len(graph))
	for k, v := range graph {
		fmt.Printf("Sub-graph[%d] length: %d\n", k, len(v))
	}
	if count < 100 {
		//not suitable for large graphs:
		fmt.Println("FULL GRAPH: ", graph)
	}
	fmt.Printf("\n")
}

func searchCmd(src string, dst string) {
	if graph == nil {
		fmt.Println("No graph to search. Please scan a dictionary before searching.\n\n")
		return
	}
	if len(src) != len(dst) {
		fmt.Println("Source and destination words are of different lengths. This is outside the problem domain.\n\n")
		return
	}
	srcSubGraph, ok := graph[len(src)]
	if !ok {
		fmt.Println("Source word not found in dictionary.\n\n")
		return
	}
	srcNode, ok := srcSubGraph[src]
	if !ok {
		fmt.Println("Source word not found in dictionary.\n\n")
		return
	}
	dstSubGraph, ok := graph[len(dst)]
	if !ok {
		fmt.Println("Destination word not found in dictionary.\n\n")
		return
	}
	dstNode, ok := dstSubGraph[dst]
	if !ok {
		fmt.Println("Destination word not found in dictionary.\n\n")
		return
	}

	path, distance, found := search.Path(search.GraphNode(srcNode), search.GraphNode(dstNode))
	if !found {
		fmt.Printf("No path was found between %s and %s.\n\n", src, dst)
		return
	}
	fmt.Printf("The shortest path between %s and %s is %d transformations long.\n", src, dst, distance)
	fmt.Printf("Full path:\n")
	for k, v := range path {
		fmt.Printf("%d: %s\n", k, v.(*grapher.WordNode).Word)
	}
	fmt.Printf("\n")
}












