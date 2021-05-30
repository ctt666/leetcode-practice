package linked_list

//a+n*(b+c)+b = 2*(a+b) --> a = (n-1)(b+c) + c
//slow与fast相遇之后p以slow速运行，slow若干圈+c即和p相遇
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil && slow.Next != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
