package math_geometry

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{"Example 1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"Example 2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"Carry over", []int{9}, []int{1, 0}},
		{"All nines", []int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusOne(tt.digits); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("PlusOne(%v) = %v; want %v", tt.digits, got, tt.expected)
			}
		})
	}
}
