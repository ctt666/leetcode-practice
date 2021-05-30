package tree

type Self1 struct {
	Path []int
}

//前序遍历
func (s *Self1) PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	s.Path = append(s.Path, root.Val)
	s.PreOrder(root.Left)
	s.PreOrder(root.Right)
}

//中序遍历
func (s *Self1) InOrder(root *TreeNode) {
	if root == nil {
		return
	}

	s.InOrder(root.Left)
	s.Path = append(s.Path, root.Val)
	s.InOrder(root.Right)
}

//后序遍历
func (s *Self1) PostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	s.PostOrder(root.Left)
	s.PostOrder(root.Right)
	s.Path = append(s.Path, root.Val)
}
