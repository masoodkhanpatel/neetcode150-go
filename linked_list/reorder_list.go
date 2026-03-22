package linked_list

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. Find middle
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Reverse second half
	second := slow.Next
	slow.Next = nil
	var prev *ListNode
	for second != nil {
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	// 3. Merge two halves
	first, second := head, prev
	for second != nil {
		temp1, temp2 := first.Next, second.Next
		first.Next = second
		second.Next = temp1
		first, second = temp1, temp2
	}
}
