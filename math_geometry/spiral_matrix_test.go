package math_geometry

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"3x4", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{"1x1", [][]int{{1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SpiralOrder(%v) = %v; want %v", tt.matrix, got, tt.expected)
			}
		})
	}
}
