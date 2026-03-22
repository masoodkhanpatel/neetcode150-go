package intervals

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Two rooms needed",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			expected:  2,
		},
		{
			name:      "One room",
			intervals: [][]int{{7, 10}, {2, 4}},
			expected:  1,
		},
		{
			name:      "All overlap",
			intervals: [][]int{{1, 5}, {2, 6}, {3, 7}},
			expected:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinMeetingRooms(tt.intervals)
			if result != tt.expected {
				t.Errorf("MinMeetingRooms(%v) = %d; want %d", tt.intervals, result, tt.expected)
			}
		})
	}
}
