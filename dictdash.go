/*
For a given dictionary of words, a source word, and a destination word,
dictdash outputs the minimum number of single-letter transformations required
to go from source to destination via other words in the dictionary.
If there is no path, an error is output.
*/
package main

import (
		"fmt"
		"bufio"
		"os"
)

func main() {
	fmt.Print("Welcome to Dictionary Dash!\n")
	reader := bufio.NewReader(os.Stdin)
	commandLoop(reader)
}

func commandLoop(reader *bufio.Reader) {
	fmt.Print("Please enter a command.\n" +
		"Typing \"help\" will show a command list.\n" +
		"> ")
	text, error := reader.ReadString('\n')
	if error != nil	{
		fmt.Println("Sorry, there was an input error:\n%d", error)
		return
	}
	switch text {
	case "help\n":
		fmt.Println("\n***Command List***\n")
		fmt.Println("help\t\tshows this command list")
		fmt.Println("quit\t\tquits the application")
		fmt.Println("")
	case "quit\n":
		fmt.Println("Quitting.\n")
		return
	default:
		fmt.Println("Sorry, command not understood.\n")
	}
	commandLoop(reader)
}



