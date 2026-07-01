package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func newNode(val int) *Node {
	return &Node{
		Val:       val,
		Neighbors: make([]*Node, 0),
	}
}

func (n *Node) addNeighbor(node *Node) {
	n.Neighbors = append(n.Neighbors, node)
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	table := make(map[int]*Node)

	var cloneNode func(node *Node) *Node
	cloneNode = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		var clone *Node
		var ok bool

		clone, ok = table[node.Val]
		if ok {
			return clone
		}

		clone = newNode(node.Val)
		table[node.Val] = clone

		for _, nb := range node.Neighbors {
			cloneNb := cloneNode(nb)
			clone.addNeighbor(cloneNb)
		}

		return clone
	}

	return cloneNode(node)
}
