package backtracking

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	tests := []struct {
		name     string
		word     string
		expected bool
	}{
		{"Example 1", "ABCCED", true},
		{"Example 2", "SEE", true},
		{"Example 3", "ABCB", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Exist(board, tt.word)
			if result != tt.expected {
				t.Errorf("Exist(%q) = %v; want %v", tt.word, result, tt.expected)
			}
		})
	}
}
