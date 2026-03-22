package stack

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n:        1,
			expected: []string{"()"},
		},
	}

	for _, tt := range tests {
		actual := generateParenthesis(tt.n)
		sort.Strings(actual)
		sort.Strings(tt.expected)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("generateParenthesis(%d) = %v; expected %v", tt.n, actual, tt.expected)
		}
	}
}
