package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	isSame := isSameTree(p.Left, q.Left)
	if !isSame {
		return false
	}

	isSame = isSameTree(p.Right, q.Right)
	if !isSame {
		return false
	}

	return true
}
