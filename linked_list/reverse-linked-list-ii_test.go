package linked_list

import (
	"fmt"
	"testing"
)

//明确链表反转过程改变的节点，以及新链表的连接边界
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	dummy := pre
	cur := head
	pLeft, pRight := &ListNode{}, &ListNode{}
	for cur != nil {
		if left == 1 {
			pLeft = cur
			right = right - left
			break
		}
		pre = pre.Next
		cur = cur.Next
		left--
	}

	for cur != nil {
		if right == 1 {
			pRight = cur
			break
		}
		cur = cur.Next
		right--
	}
	next := cur.Next

	if left > 1 || right > 1 {
		return head
	}
	pLeft, pRight = reverse2(pLeft, pRight)
	pre.Next = pLeft
	pRight.Next = next
	return dummy.Next
}

func reverse2(left, right *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := left
	for pre != right {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	return right, left
}

func TestReverseBetween(t *testing.T) {
	p5 := &ListNode{5, nil}
	p4 := &ListNode{4, p5}
	p3 := &ListNode{3, p4}
	p2 := &ListNode{2, p3}
	p1 := &ListNode{1, p2}
	head := reverseBetween(p1, 2, 4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
