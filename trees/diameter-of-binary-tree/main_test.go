package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		//[1,2,3,4,5]
		root := BuildTree([]*int{Ptr(1), Ptr(2), Ptr(3), Ptr(4), Ptr(5)})
		assert.Equal(t, 3, diameterOfBinaryTree(root))
	})

	t.Run("Test case 2", func(t *testing.T) {
		//[3,1,null,null,2]
		root := BuildTree([]*int{Ptr(3), Ptr(1), nil, nil, Ptr(2)})
		root.Print()
		assert.Equal(t, 2, diameterOfBinaryTree(root))
	})
}

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

func Ptr(i int) *int { return &i }

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
