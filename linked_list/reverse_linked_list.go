package linked_list

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, cur, pre = pre, cur.Next, cur
	}
	return pre
}
