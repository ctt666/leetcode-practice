package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(1+maxDepth(root.Left), 1+maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
