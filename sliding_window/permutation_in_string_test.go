package sliding_window

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "Example 1",
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
		{
			name:     "Example 2",
			s1:       "ab",
			s2:       "eidboaoo",
			expected: false,
		},
		{
			name:     "Short s2",
			s1:       "abc",
			s2:       "ab",
			expected: false,
		},
		{
			name:     "Exact match",
			s1:       "abc",
			s2:       "abc",
			expected: true,
		},
		{
			name:     "Permutation at end",
			s1:       "abc",
			s2:       "xyzabc",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckInclusion(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("CheckInclusion(%q, %q) = %v; want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
