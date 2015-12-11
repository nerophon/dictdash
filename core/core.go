/*
*/
package core

import (
		"fmt"
		"io"
		"bufio"
		"strconv"
)

// strings in Go are read-only byte array slices
// indexing [] into them returns the byte
// this will be easy for us because each alphabet letter is one byte
const alphabet string = "abcdefghijklmnopqrstuvwxyz"

// Go uses utf8, so to have small 'a' be zero you need to subtract 97
const decodeOffset byte = 97

// initial node edge slice allocation
const initialEdgeCount uint = 26

// the Edges structure is: [letterIndex][replacementIndex]
// during a scan, the replacementIndex is fully allocated to a size of 26,
// and represents the index of the replacement letter in the alphabet;
// however at the end of the scan this is compressed so as to remove nil entries;
// after this point the replacementIndex bears no relation to the alphabet
// being optimized for iteration rather than lookup
type Node struct {
	Word string
	Edges [][]*Node
}

func NewNode(word string, edgeCount uint) (node *Node) {
	node = new(Node)
	node.Word = word
	node.Edges = make([][]*Node, len(word))
	for k, _ := range node.Edges {
		node.Edges[k] = make([]*Node, edgeCount)
	}
	return
}

// Not to be used in production, only for debugging SMALL dictionaries
func (node *Node) String() string {
	result := "\nWord: " + node.Word
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
	return result + "\n"
}


// panics if input == nil
func ScanAndGraph(input io.Reader) (dict map[int]map[string]*Node, count int, err error) {
	dict, count, err = scan(input)
	graph(dict)
	return
}

// panics if input == nil
func scan(input io.Reader) (dict map[int]map[string]*Node, count int, err error) {
	dict = make(map[int]map[string]*Node)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if addToDictionary(scanner.Text(), dict) {
			count++
		}
	}
	if err = scanner.Err(); err != nil {
		dict = nil
		return
	}
	return
}

// panics if dict == nil
func addToDictionary(word string, dict map[int]map[string]*Node) (added bool) {
	letterCount := len(word)
	if _, ok := dict[letterCount]; !ok {
		dict[letterCount] = make(map[string]*Node)
	}
	if _, ok := dict[letterCount][word]; ok {
		return false;
	}
	dict[letterCount][word] = NewNode(word, initialEdgeCount)
	return true;
}

// panics if dict == nil
func graph(dict map[int]map[string]*Node) {
	// opportunity for concurrency here
	for letterCount, subDict := range dict {
		for word, node := range subDict {
			for letter := 0; letter < letterCount; letter++ {
				for replacement := 0; replacement < 26; replacement++ {
					searchWord := replaceAtIndex(word, alphabet[replacement], letter)
					if searchWord == word {
						continue // self should remain nil
					}
					if node.Edges[letter][replacement] != nil {
						continue // mirror, no need to lookup
					}
					if connectedNode, ok := subDict[searchWord]; ok {
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

// panics if dict == nil
func compress(dict map[int]map[string]*Node) {

}

func Search(src string, dst string) {
	fmt.Println("Sorry, command not yet implemented.\n")
}












