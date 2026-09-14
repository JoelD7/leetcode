package main

type Trie struct {
	root *Node
}

type Node struct {
	key      string
	children []*Node
	isWord   bool
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			children: make([]*Node, 26),
		},
	}
}

func NewNode(key string) *Node {
	return &Node{
		key:      key,
		children: make([]*Node, 26),
	}
}

func (this *Trie) Insert(word string) {
	node := this.root

	var index int32
	for i, c := range word {
		index = c - 'a'
		if node.children[index] == nil {
			node.children[index] = NewNode(string(c))
		}

		if i == len(word)-1 {
			node.children[index].isWord = true
		}

		node = node.children[index]
	}
}

func (this *Trie) Search(word string) bool {
	node := this.root

	var index int32
	for _, c := range word {
		index = c - 'a'
		if node.children[index] == nil {
			return false
		}

		node = node.children[index]
	}

	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root

	var index int32
	for _, c := range prefix {
		index = c - 'a'
		if node.children[index] == nil {
			return false
		}

		node = node.children[index]
	}

	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
