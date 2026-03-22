package backtracking

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "Example 1",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Example 2",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "Example 3",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LetterCombinations(tt.digits)
			if !reflect.DeepEqual(result, tt.expected) && !(len(result) == 0 && len(tt.expected) == 0) {
				t.Errorf("LetterCombinations(%q) = %v; want %v", tt.digits, result, tt.expected)
			}
		})
	}
}
