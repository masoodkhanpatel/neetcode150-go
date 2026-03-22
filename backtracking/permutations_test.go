package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Permute(tt.nums)

			// Helper to sort and compare
			sortInner := func(arr [][]int) {
				sort.Slice(arr, func(i, j int) bool {
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
				t.Errorf("Permute(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
