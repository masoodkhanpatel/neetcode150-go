package heap

import (
	"reflect"
	"sort"
	"testing"
)

func TestKClosest(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1",
			points:   [][]int{{1, 3}, {-2, 2}},
			k:        1,
			expected: [][]int{{-2, 2}},
		},
		{
			name:     "Example 2",
			points:   [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:        2,
			expected: [][]int{{3, 3}, {-2, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KClosest(tt.points, tt.k)
			// Sort result and expected for comparison
			sort.Slice(result, func(i, j int) bool {
				if result[i][0] != result[j][0] {
					return result[i][0] < result[j][0]
				}
				return result[i][1] < result[j][1]
			})
			sort.Slice(tt.expected, func(i, j int) bool {
				if tt.expected[i][0] != tt.expected[j][0] {
					return tt.expected[i][0] < tt.expected[j][0]
				}
				return tt.expected[i][1] < tt.expected[j][1]
			})
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KClosest(%v, %d) = %v; want %v", tt.points, tt.k, result, tt.expected)
			}
		})
	}
}
