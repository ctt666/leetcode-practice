package linked_list

//b->a->d->c
func swapPairs(head *ListNode) *ListNode {
	//pre用来返回swap后的链表
	pre := &ListNode{}
	pre.Next = head
	//tmp用于操作swap
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		a, b := tmp.Next, tmp.Next.Next
		//更新a.Next b.Next tmp.Next tmp
		a.Next, b.Next, tmp.Next = b.Next, a, b
		tmp = a
	}
	return pre.Next
}
