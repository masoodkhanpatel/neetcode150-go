package linked_list

import (
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	// Create nodes
	n1 := &Node{Val: 7}
	n2 := &Node{Val: 13}
	n3 := &Node{Val: 11}
	n4 := &Node{Val: 10}
	n5 := &Node{Val: 1}

	// Link Next
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	// Link Random
	n1.Random = nil
	n2.Random = n1
	n3.Random = n5
	n4.Random = n3
	n5.Random = n1

	res := copyRandomList(n1)

	// Verify copy
	currOld := n1
	currNew := res
	for currOld != nil {
		if currNew == nil {
			t.Fatal("Copied list is shorter than original")
		}
		if currNew.Val != currOld.Val {
			t.Errorf("Value mismatch: expected %d, got %d", currOld.Val, currNew.Val)
		}
		if currNew == currOld {
			t.Error("Node was not deep copied")
		}
		if currOld.Random != nil {
			if currNew.Random == nil || currNew.Random.Val != currOld.Random.Val {
				t.Errorf("Random pointer mismatch at node %d", currOld.Val)
			}
			if currNew.Random == currOld.Random {
				t.Errorf("Random pointer was not deep copied at node %d", currOld.Val)
			}
		} else if currNew.Random != nil {
			t.Errorf("Random pointer should be nil at node %d", currOld.Val)
		}
		currOld = currOld.Next
		currNew = currNew.Next
	}
	if currNew != nil {
		t.Error("Copied list is longer than original")
	}
}
