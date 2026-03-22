package graph

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:          "Example 1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      true,
		},
		{
			name:          "Example 2",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanFinish(tt.numCourses, tt.prerequisites)
			if result != tt.expected {
				t.Errorf("CanFinish() = %v; want %v", result, tt.expected)
			}
		})
	}
}
