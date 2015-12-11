/*
For a given dictionary of words, a source word, and a destination word,
dictdash outputs the minimum number of single-letter transformations required
to go from source to destination via other words in the dictionary.
If there is no path, an error is output.
*/
package main

import (
		"fmt"
		"strings"
		"bufio"
		"os"
		"github.com/nerophon/dictdash/core"
)


var dictionary map[int]map[string]*core.Node

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
		fmt.Println("scan\t\tscans a space-delimited dictionary at ./dict.txt")
		fmt.Println("scan [path]\tscans a space-delimited dictionary at [path]")
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
			search(fields[1], fields[2])
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
	dictionary, count, err = core.ScanAndGraph(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input failed with error:\n", err)
		return
	}
	fmt.Printf("Words scanned: %d\n", count)
	fmt.Printf("Sub-Dictionary count: %d\n", len(dictionary))
	for k, v := range dictionary {
		fmt.Printf("Dictionary[%d] length: %d\n", k, len(v))
	}
	fmt.Println("Dictionary: ", dictionary)
	fmt.Printf("\n")
}

func search(src string, dst string) {
	fmt.Println("Sorry, command not yet implemented.\n")
}












