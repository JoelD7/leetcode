package utils

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	BuildTree takes a level-order slice where nil represents missing nodes

	How to build trees?
	1. *Read level-by-level*: Write down the nodes from top to bottom, reading each row from left to right.

	2. *Record nil for missing children of existing nodes*: If a node exists but is missing a left or right child, you
		must put a nil in the array.

	3. *Ignore children of nil nodes*: If you put a nil in the array, do not write nils for its imaginary children on
		the next row. The queue naturally skips them.

*
*/
func BuildTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			current.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, current.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			current.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}

// Print outputs the tree to the console starting from the root.
func (n *TreeNode) Print() {
	if n == nil {
		fmt.Println("<empty tree>")
		return
	}
	fmt.Println(n.Val)
	printSubTree(n.Left, "", true)
	printSubTree(n.Right, "", false)
}

// printSubTree handles the recursive drawing of branches.
func printSubTree(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	// 1. Draw the current branch
	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── L: ")
	} else {
		fmt.Print("└── R: ")
	}

	// 2. Print the node value
	fmt.Println(node.Val)

	// 3. Calculate the prefix for the next level down
	childPrefix := prefix
	if isLeft {
		childPrefix += "│      " // Left children carry a vertical pipe
	} else {
		childPrefix += "       " // Right children leave a blank space
	}

	// 4. Recurse
	printSubTree(node.Left, childPrefix, true)
	printSubTree(node.Right, childPrefix, false)
}

func Ptr(i int) *int { return &i }
