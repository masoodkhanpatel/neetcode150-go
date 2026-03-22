package intervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Overlapping",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Touch at boundary",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "Single interval",
			intervals: [][]int{{1, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "All overlap",
			intervals: [][]int{{1, 10}, {2, 6}, {3, 8}},
			expected:  [][]int{{1, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeIntervals(tt.intervals)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeIntervals(%v) = %v; want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}
