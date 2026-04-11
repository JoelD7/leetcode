package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderPos := make(map[int]int)
	var head int

	for i := 0; i < len(inorder); i++ {
		inorderPos[inorder[i]] = i
	}

	var dfs func(start, end int) *TreeNode
	dfs = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}

		root := &TreeNode{
			Val: preorder[head],
		}
		head++

		mid := inorderPos[root.Val]

		root.Left = dfs(start, mid-1)
		root.Right = dfs(mid+1, end)

		return root
	}

	return dfs(0, len(preorder)-1)
}
