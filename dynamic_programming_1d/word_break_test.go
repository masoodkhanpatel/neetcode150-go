package dynamic_programming_1d

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, tt := range tests {
		if got := WordBreak(tt.s, tt.wordDict); got != tt.expected {
			t.Errorf("WordBreak(%q, %v) = %v; want %v", tt.s, tt.wordDict, got, tt.expected)
		}
	}
}
