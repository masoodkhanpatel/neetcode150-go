package graph

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		expected []int
	}{
		{
			name:     "Example 1",
			edges:    [][]int{{1, 2}, {1, 3}, {2, 3}},
			expected: []int{2, 3},
		},
		{
			name:     "Example 2",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			expected: []int{1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindRedundantConnection(tt.edges)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FindRedundantConnection() = %v; want %v", result, tt.expected)
			}
		})
	}
}
