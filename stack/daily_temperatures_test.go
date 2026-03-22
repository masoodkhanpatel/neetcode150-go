package stack

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		temperatures []int
		expected     []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
	}

	for _, tt := range tests {
		actual := dailyTemperatures(tt.temperatures)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("dailyTemperatures(%v) = %v; expected %v", tt.temperatures, actual, tt.expected)
		}
	}
}
