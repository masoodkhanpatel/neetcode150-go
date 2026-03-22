package two_pointers

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Example 1",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Example 2",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    " ",
			expected: true,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "Non-alphanumeric only",
			input:    ".,",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
