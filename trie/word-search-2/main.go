package main

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	root := buildTrieNode(words)

	var dfs func(node *TrieNode, i, j int)
	dfs = func(node *TrieNode, i, j int) {
		c := board[i][j]

		if c == '#' || node.children[c-'a'] == nil {
			return
		}

		node = node.children[c-'a']
		if node.word != nil {
			res = append(res, *node.word)
			node.word = nil
		}

		board[i][j] = '#'

		if i < len(board)-1 {
			dfs(node, i+1, j)
		}
		if i > 0 {
			dfs(node, i-1, j)
		}
		if j < len(board[i])-1 {
			dfs(node, i, j+1)
		}
		if j > 0 {
			dfs(node, i, j-1)
		}

		board[i][j] = c
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(root, i, j)
		}
	}

	return res
}

type TrieNode struct {
	word     *string
	children []*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
	}
}

func buildTrieNode(words []string) *TrieNode {
	root := newTrieNode()
	var i int32

	for _, word := range words {
		node := root
		for _, c := range word {
			i = c - 'a'
			if node.children[i] == nil {
				node.children[i] = newTrieNode()
			}
			node = node.children[i]
		}
		w := word
		node.word = &w
	}

	return root
}
