package backtracking

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]string
	}{
		{
			name: "Example 1",
			n:    4,
			expected: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name:     "Example 2",
			n:        1,
			expected: [][]string{{"Q"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SolveNQueens(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SolveNQueens(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}
