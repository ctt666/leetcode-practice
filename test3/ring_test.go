package test3

import (
	"fmt"
	"testing"
)

//链表

type Node struct {
	Next  *Node
	Value int
}

//2,4,4,7,8,9
func insert(p *Node, k int) {
	if p == nil {
		return
	}

	m := make(map[*Node]bool)
	pre := p
	p = p.Next

	for !m[p] {
		m[p] = true
		if p.Value >= k && pre.Value <= k {
			break
		}
		if (p.Value >= k || pre.Value <= k) && p.Value < pre.Value {
			break
		}
		p = p.Next
		pre = pre.Next
	}

	pre.Next = &Node{
		Next:  p,
		Value: k,
	}
}

func TestRing(t *testing.T) {
	p := &Node{
		Value: 4,
	}
	p.Next = p
	insert(p, 4)
	insert(p, 4)
	insert(p, 2)
	//insert(p, 8)
	//insert(p, 9)
	for i := 0; i < 6; i++ {
		fmt.Println(p.Value)
		p = p.Next
	}
}
