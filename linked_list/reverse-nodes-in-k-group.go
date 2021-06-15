package linked_list

//TODO logic
func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	tmp := pre
	for tmp.Next != nil {
		p1, p2 := tmp.Next, tmp
		//构造kgroup 退出条件
		for i := 0; i < k; i++ {
			p2 = p2.Next
			if p2 == nil {
				return pre.Next
			}
		}
		//存储next
		next := p2.Next

		p1, p2 = reverse(p1, p2)
		p2.Next = next
		tmp.Next = p1
		tmp = p2
	}
	return pre.Next
}

//reverse list
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	//退出条件
	for head != tail.Next {
		cur, cur.Next, pre = cur.Next, pre, cur
	}

	return tail, head
}
