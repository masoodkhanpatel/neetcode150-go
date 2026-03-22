package heap

import "testing"

func TestKthLargest(t *testing.T) {
	kl := Constructor(3, []int{4, 5, 8, 2})
	tests := []struct {
		val      int
		expected int
	}{
		{3, 4},
		{5, 5},
		{10, 5},
		{9, 8},
		{4, 8},
	}

	for _, tt := range tests {
		res := kl.Add(tt.val)
		if res != tt.expected {
			t.Errorf("kl.Add(%d) = %d; want %d", tt.val, res, tt.expected)
		}
	}
}
