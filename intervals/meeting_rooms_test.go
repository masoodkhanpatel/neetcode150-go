package intervals

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  bool
	}{
		{
			name:      "Can attend",
			intervals: [][]int{{0, 30}, {35, 50}, {60, 70}},
			expected:  true,
		},
		{
			name:      "Cannot attend overlap",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			expected:  false,
		},
		{
			name:      "Single meeting",
			intervals: [][]int{{0, 5}},
			expected:  true,
		},
		{
			name:      "Touch boundary ok",
			intervals: [][]int{{0, 5}, {5, 10}},
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanAttendMeetings(tt.intervals)
			if result != tt.expected {
				t.Errorf("CanAttendMeetings(%v) = %v; want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}
