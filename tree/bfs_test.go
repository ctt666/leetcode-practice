package tree

import (
	"fmt"
	"testing"
)

func BFS(root *TreeNode) []*TreeNode {
	visited := []*TreeNode{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		visited = append(visited, node)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		queue = queue[1:]
	}
	return visited
}

func TestBFS(t *testing.T) {
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
	fmt.Println(BFS(root))
}
