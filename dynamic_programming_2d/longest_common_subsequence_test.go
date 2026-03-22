package dynamic_programming_2d

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		t1, t2   string
		expected int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, tt := range tests {
		if got := LongestCommonSubsequence(tt.t1, tt.t2); got != tt.expected {
			t.Errorf("LCS(%q, %q) = %d; want %d", tt.t1, tt.t2, got, tt.expected)
		}
	}
}
