package arrays_hashing

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Example 1: Anagram",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "Example 2: Not an anagram",
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			name:     "Different lengths",
			s:        "ab",
			t:        "a",
			expected: false,
		},
		{
			name:     "Empty strings",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "Same characters different counts",
			s:        "aacc",
			t:        "ccac",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("IsAnagram(%q, %q) = %v; want %v", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
