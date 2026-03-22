package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected: [][]int{
				{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6},
			},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected:   [][]int{{1, 2, 2}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombinationSum2(tt.candidates, tt.target)

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
				t.Errorf("CombinationSum2(%v, %d) = %v; want %v", tt.candidates, tt.target, result, tt.expected)
			}
		})
	}
}
