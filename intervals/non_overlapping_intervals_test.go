package intervals

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Remove one",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			expected:  1,
		},
		{
			name:      "Remove two",
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			expected:  2,
		},
		{
			name:      "No removal",
			intervals: [][]int{{1, 2}, {2, 3}},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EraseOverlapIntervals(tt.intervals)
			if result != tt.expected {
				t.Errorf("EraseOverlapIntervals(%v) = %d; want %d", tt.intervals, result, tt.expected)
			}
		})
	}
}
