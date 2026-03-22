package intervals

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name:        "Insert in middle",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "Merge multiple",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "No overlap",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{6, 8},
			expected:    [][]int{{1, 5}, {6, 8}},
		},
		{
			name:        "Empty intervals",
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expected:    [][]int{{5, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InsertInterval(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertInterval(%v, %v) = %v; want %v", tt.intervals, tt.newInterval, result, tt.expected)
			}
		})
	}
}
