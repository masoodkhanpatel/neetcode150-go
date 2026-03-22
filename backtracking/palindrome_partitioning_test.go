package backtracking

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected [][]string
	}{
		{
			name:     "Example 1",
			s:        "aab",
			expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name:     "Example 2",
			s:        "a",
			expected: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Partition(tt.s)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Partition(%q) = %v; want %v", tt.s, result, tt.expected)
			}
		})
	}
}
