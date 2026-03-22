package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombinationSum(tt.candidates, tt.target)

			// Helper to sort and compare
			sortInner := func(arr [][]int) {
				for i := range arr {
					sort.Ints(arr[i])
				}
				sort.Slice(arr, func(i, j int) bool {
					if len(arr[i]) != len(arr[j]) {
						return len(arr[i]) < len(arr[j])
					}
					for k := 0; k < len(arr[i]); k++ {
						if arr[i][k] != arr[j][k] {
							return arr[i][k] < arr[j][k]
						}
					}
					return false
				})
			}

			sortInner(result)
			sortInner(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CombinationSum(%v, %d) = %v; want %v", tt.candidates, tt.target, result, tt.expected)
			}
		})
	}
}
