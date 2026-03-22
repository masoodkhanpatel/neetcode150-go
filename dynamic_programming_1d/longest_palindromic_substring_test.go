package dynamic_programming_1d

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"racecar", "racecar"},
	}
	for _, tt := range tests {
		got := LongestPalindrome(tt.s)
		// Check it's actually a palindrome of the right length
		if len(got) != len(tt.expected) {
			t.Errorf("LongestPalindrome(%q) = %q (len %d); want len %d", tt.s, got, len(got), len(tt.expected))
		}
	}
}
