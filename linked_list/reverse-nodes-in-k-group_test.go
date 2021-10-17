package linked_list

import (
	"fmt"
	"testing"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	tmp := pre
	for tmp.Next != nil {
		p1, p2 := tmp.Next, tmp
		for i := 0; i < k; i++ {
			p2 = p2.Next
			if p2 == nil {
				return pre.Next
			}
		}

		next := p2.Next
		p1, p2 = reverse(p1, p2)
		p2.Next = next
		tmp.Next = p1
		tmp = p2
	}
	return pre.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	for pre != tail {
		cur, cur.Next, pre = cur.Next, pre, cur
	}

	return tail, head
}

func TestReverseInKGroup(t *testing.T) {
	p5 := &ListNode{5, nil}
	p4 := &ListNode{4, p5}
	p3 := &ListNode{3, p4}
	p2 := &ListNode{2, p3}
	p1 := &ListNode{1, p2}
	head := reverseKGroup(p1, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
