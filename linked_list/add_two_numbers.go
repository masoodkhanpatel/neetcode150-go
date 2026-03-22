package linked_list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		val := v1 + v2 + carry
		carry = val / 10
		val = val % 10
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	return dummy.Next
}
