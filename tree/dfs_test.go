package tree

import (
	"fmt"
	"testing"
)

func DFS(root *TreeNode, visited map[*TreeNode]bool) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	visited[root] = true
	for _, v := range root.Children() {
		if !visited[v] {
			DFS(v, visited)
		}
	}
}

func TestDFS(t *testing.T) {
	l1 := &TreeNode{
		Val: 4,
	}
	l2 := &TreeNode{
		Val:  2,
		Left: l1,
	}
	r1 := &TreeNode{
		Val: 3,
	}
	root := &TreeNode{
		Val:   1,
		Left:  l2,
		Right: r1,
	}
	visited := map[*TreeNode]bool{}
	DFS(root, visited)
	fmt.Println(visited)
}
