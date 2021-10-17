package tree

import (
	"fmt"
	"math"
	"testing"
)

//TODO
type Self struct {
	Pre *TreeNode
}

func isValidBST(root *TreeNode) bool {
	s := &Self{}
	return s.helper(root)
}

func (self *Self) helper(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !self.helper(root.Left) {
		return false
	}
	if self.Pre != nil && root.Val <= self.Pre.Val {
		return false
	}
	self.Pre = root
	return self.helper(root.Right)
}

func TestValide(t *testing.T) {
	left := &TreeNode{
		Val: 1,
	}

	root := &TreeNode{
		Val:  1,
		Left: left,
	}
	fmt.Println(isValidBST(root))
}

//根据范围， 开区间
func isValidBST2(root *TreeNode) bool {
	return helper2(root, math.MinInt32, math.MaxInt32)
}

func helper2(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return helper2(root.Left, min, root.Val) && helper2(root.Right, root.Val, max)
}
