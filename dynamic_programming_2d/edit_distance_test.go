package dynamic_programming_2d

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		word1, word2 string
		expected     int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "", 0},
		{"a", "", 1},
	}
	for _, tt := range tests {
		if got := MinDistance(tt.word1, tt.word2); got != tt.expected {
			t.Errorf("MinDistance(%q, %q) = %d; want %d", tt.word1, tt.word2, got, tt.expected)
		}
	}
}
