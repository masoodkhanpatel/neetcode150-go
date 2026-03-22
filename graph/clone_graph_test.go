package graph

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	cloned := CloneGraph(n1)

	if cloned == n1 {
		t.Errorf("Cloned graph should be a deep copy, not the same reference")
	}

	if cloned.Val != n1.Val {
		t.Errorf("Cloned node value mismatch: got %d, want %d", cloned.Val, n1.Val)
	}

	// Basic check for neighbors
	if len(cloned.Neighbors) != len(n1.Neighbors) {
		t.Errorf("Cloned node neighbor count mismatch")
	}

	// Verify deep copy of neighbor
	if cloned.Neighbors[0] == n2 {
		t.Errorf("Neighbor should be a deep copy")
	}
}
