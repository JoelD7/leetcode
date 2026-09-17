package main

type Node struct {
	key      string
	children []*Node
	isWord   bool
}
type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewNode(""),
	}
}

func NewNode(key string) *Node {
	return &Node{
		key:      key,
		children: make([]*Node, 26),
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	var index, char int32

	for i := 0; i < len(word); i++ {
		char = int32(word[i])
		index = char - 'a'

		if node.children[index] == nil {
			node.children[index] = NewNode(string(char))
		}
		node = node.children[index]
	}
	node.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.root.search(word)
}

// Searches a word for the trie starting at n
func (n *Node) search(word string) bool {
	node := n
	var index, char int32

	for i := 0; i < len(word); i++ {
		char = int32(word[i])
		index = char - 'a'

		if char == '.' {
			for _, child := range node.children {
				if child != nil && child.search(word[i+1:]) {
					return true
				}
			}
			return false
		}

		if node.children[index] == nil {
			return false
		}

		node = node.children[index]
	}

	return node != nil && node.isWord
}
