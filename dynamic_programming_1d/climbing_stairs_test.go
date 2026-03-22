package dynamic_programming_1d

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct{ n, expected int }{
		{1, 1}, {2, 2}, {3, 3}, {4, 5}, {10, 89},
	}
	for _, tt := range tests {
		if got := ClimbStairs(tt.n); got != tt.expected {
			t.Errorf("ClimbStairs(%d) = %d; want %d", tt.n, got, tt.expected)
		}
	}
}
