package greedy

import "testing"

func TestMergeTriplets(t *testing.T) {
	tests := []struct {
		name     string
		triplets [][]int
		target   []int
		expected bool
	}{
		{"Example 1", [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}, true},
		{"Example 2", [][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}, false},
		{"Example 3", [][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}}, []int{2, 5, 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTriplets(tt.triplets, tt.target); got != tt.expected {
				t.Errorf("MergeTriplets(%v, %v) = %v; want %v", tt.triplets, tt.target, got, tt.expected)
			}
		})
	}
}
