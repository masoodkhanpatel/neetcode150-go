package dynamic_programming_1d

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"abc", 3},
		{"aaa", 6},
		{"a", 1},
	}
	for _, tt := range tests {
		if got := CountSubstrings(tt.s); got != tt.expected {
			t.Errorf("CountSubstrings(%q) = %d; want %d", tt.s, got, tt.expected)
		}
	}
}
