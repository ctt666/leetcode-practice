package test

type Node struct {
	Val  int
	Next *Node
}

func hasCycle(root *Node) int {
	if root == nil {
		return 0
	}

	p1, p2 := root, root

	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			break
		}
	}

	if p2 == nil || p2.Next == nil {
		return 0
	}

	count := 1
	for p2 != p1.Next {
		p1 = p1.Next
		count++
	}
	return count
}
