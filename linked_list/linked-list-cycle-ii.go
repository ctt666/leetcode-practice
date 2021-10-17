package linked_list

//找环的入口：快慢指针相遇后，新的指针从起点走，相遇处即环的入口
//a+n*(b+c)+b = 2*(a+b) --> a = (n-1)(b+c) + c
//slow与fast相遇之后p以slow速运行，slow若干圈+c即和p相遇
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			//新建指针
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
