package arrays_hashing

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Example 3: Order doesn't matter",
			nums:     []int{1, 2},
			k:        2,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TopKFrequent(tt.nums, tt.k)
			sort.Ints(result)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TopKFrequent(%v, %d) = %v; want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
