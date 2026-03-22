package sliding_window

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "abcabcbb",
			expected: 3,
		},
		{
			name:     "Example 2",
			s:        "bbbbb",
			expected: 1,
		},
		{
			name:     "Example 3",
			s:        "pwwkew",
			expected: 3,
		},
		{
			name:     "Empty",
			s:        "",
			expected: 0,
		},
		{
			name:     "Space",
			s:        " ",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LengthOfLongestSubstring(tt.s)
			if result != tt.expected {
				t.Errorf("LengthOfLongestSubstring(%q) = %d; want %d", tt.s, result, tt.expected)
			}
		})
	}
}
