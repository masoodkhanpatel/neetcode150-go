package binary_search

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}

	for _, tt := range tests {
		got := searchRotated(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("searchRotated(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
