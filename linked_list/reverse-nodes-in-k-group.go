package linked_list

func reverseKGroup(head *ListNode, k int) *ListNode {
	//用于记录链表
	pre := &ListNode{
		Next: head,
	}
	//用于遍历
	tmp := pre
	//遍历推出条件:1.next=nil;2.k的大小
	for tmp.Next != nil {
		//确定段k边界
		p1, p2 := tmp.Next, tmp
		for i := 0; i < k; i++ {
			p2 = p2.Next
			if p2 == nil {
				return pre.Next
			}
		}

		//在reverse之前缓存next位置（否则反转之后找不到）
		next := p2.Next
		//段内实现reverse
		p1, p2 = reverse(p1, p2)
		//reverse之后操作：1.p2->next;2.tmp->p1;3.tmp=p2
		p2.Next = next
		tmp.Next = p1
		tmp = p2
	}
	return pre.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	//初始化前继节点
	var pre *ListNode
	cur := head
	//边界条件：不能用head != tail.next,因为tail节点next指向已修改
	for pre != tail {
		cur, cur.Next, pre = cur.Next, pre, cur
	}

	return tail, head
}
