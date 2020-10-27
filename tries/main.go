package main

import "fmt"

// AlphabetSize represents the amount of possible characters in the trie node 
const AlphabetSize = 26

// Node represents each node in a trie. A location for each character in a word. Node is comprised of a slice a characters and a boolean to deteremine if word is completed.
type Node struct {
	children [AlphabetSize]*Node
	isEnd bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie creates a new Trie
func InitTrie() *Trie {
	result := &Trie{root:&Node{}}
	return result
}

// Insert method interates over the word argument and inserts each character into the trie until the word is complete. 
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a';
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search method takes in a word argument and RETURNS true if the word exists and false if not.
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a';
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}

func main() {
	myTrie := InitTrie()

	wordData := []string {
		"donte",
		"binary",
		"stacks",
		"queue",
		"linked",
		"bible",
		"donuts",
	}

	for _, v := range wordData {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("link"))
}