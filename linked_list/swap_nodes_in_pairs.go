package linked_list

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		a, b := tmp.Next, tmp.Next.Next
		a.Next, b.Next, tmp.Next = b.Next, a, b
		tmp = a
	}
	return pre.Next
}
