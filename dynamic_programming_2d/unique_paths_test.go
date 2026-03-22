package dynamic_programming_2d

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct{ m, n, expected int }{
		{3, 7, 28}, {3, 2, 3}, {1, 1, 1}, {7, 3, 28},
	}
	for _, tt := range tests {
		if got := UniquePaths(tt.m, tt.n); got != tt.expected {
			t.Errorf("UniquePaths(%d,%d) = %d; want %d", tt.m, tt.n, got, tt.expected)
		}
	}
}
