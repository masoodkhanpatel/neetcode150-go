package linked_list

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	// Case 1: Cycle exists
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	if !hasCycle(n1) {
		t.Error("expected true, got false")
	}

	// Case 2: No cycle
	n5 := &ListNode{Val: 1}
	n6 := &ListNode{Val: 2}
	n5.Next = n6
	if hasCycle(n5) {
		t.Error("expected false, got true")
	}

	// Case 3: Empty list
	if hasCycle(nil) {
		t.Error("expected false, got true")
	}
}
