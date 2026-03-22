package math_geometry

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			"Example 1",
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			"Example 2",
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetZeroes(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.expected) {
				t.Errorf("SetZeroes() = %v; want %v", tt.matrix, tt.expected)
			}
		})
	}
}
