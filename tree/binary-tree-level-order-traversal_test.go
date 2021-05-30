package tree

import (
	"fmt"
	"testing"
)

func levelOrder_BFS(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	for len(queue) > 0 {
		current := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			current = append(current, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, current)
		queue = queue[length:]
	}
	return result
}

func levelOrder_DFS(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	dfs(root, 0, result)
	return result
}

func dfs(root *TreeNode, level int, result [][]int) {
	if root == nil {
		return
	}
	if len(result) < level+1 {
		result = append(result, []int{})
	}
	result[level] = append(result[level], root.Val)
	dfs(root.Left, level+1, result)
	dfs(root.Right, level+1, result)
}

func TestDfs(t *testing.T) {
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
	fmt.Println(levelOrder_DFS(root))
}
