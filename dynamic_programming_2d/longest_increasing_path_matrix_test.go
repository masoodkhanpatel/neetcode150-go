package dynamic_programming_2d

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected int
	}{
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4},
		{[][]int{{1}}, 1},
	}
	for _, tt := range tests {
		if got := LongestIncreasingPath(tt.matrix); got != tt.expected {
			t.Errorf("LongestIncreasingPath(%v) = %d; want %d", tt.matrix, got, tt.expected)
		}
	}
}
