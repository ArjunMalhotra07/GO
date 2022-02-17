package main

import "fmt"

//Number if possible characters in the trie
const AlphabetSize = 26

//Node
type Node struct {
	children [26]*Node
	isEnd    bool
}

//Trie Pointer pointing to the root Node
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//insert will take in a word and add it to the trie

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' //Index Number 14 of children array will represent letter 'O'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//search

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' //Index Number 14 of children array will represent letter 'O'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd

}
func main() {
	f := fmt.Println
	testTrie := InitTrie()
	f(testTrie.root)
	f(testTrie.Search("geralt"))

	toAdd := []string{
		"aragorn", "geralt", "rivia", "yennefer", "vengerberg",
	}

	for _, v := range toAdd {
		testTrie.Insert(v)
	}
	f(testTrie.root)
	f(testTrie.Search("geralt"))
}
