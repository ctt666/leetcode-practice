package tree

//二叉搜索树--两节点最近公共祖先
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor2(root.Right, p, q)
	}
	return root
}
