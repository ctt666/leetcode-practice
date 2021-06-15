package linked_list

//cur cur.Next pre
//退出条件
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, cur, pre = pre, cur.Next, cur
	}
	return pre
}
