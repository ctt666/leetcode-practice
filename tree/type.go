package tree

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n TreeNode) String() string {
	return fmt.Sprintf("val = %d", n.Val)
}

func (n TreeNode) Children() (r []*TreeNode) {
	if n.Left != nil {
		r = append(r, n.Left)
	}
	if n.Right != nil {
		r = append(r, n.Right)
	}
	return r
}
