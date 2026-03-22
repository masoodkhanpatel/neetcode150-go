package dynamic_programming_2d

import "testing"

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1, s2, s3 string
		expected   bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"a", "b", "ab", true},
	}
	for _, tt := range tests {
		if got := IsInterleave(tt.s1, tt.s2, tt.s3); got != tt.expected {
			t.Errorf("IsInterleave(%q, %q, %q) = %v; want %v", tt.s1, tt.s2, tt.s3, got, tt.expected)
		}
	}
}
