package greedy

import "testing"

func TestIsNStraightHand(t *testing.T) {
	tests := []struct {
		name      string
		hand      []int
		groupSize int
		expected  bool
	}{
		{"Valid groups", []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true},
		{"Invalid groups", []int{1, 2, 3, 4, 5}, 4, false},
		{"Size 1 groups", []int{1, 1, 2, 2}, 1, true},
		{"Cannot group", []int{1, 2, 3, 4, 5, 6}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNStraightHand(tt.hand, tt.groupSize); got != tt.expected {
				t.Errorf("IsNStraightHand(%v, %d) = %v; want %v", tt.hand, tt.groupSize, got, tt.expected)
			}
		})
	}
}
