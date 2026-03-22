package stack

import "testing"

func TestCarFleet(t *testing.T) {
	tests := []struct {
		target   int
		position []int
		speed    []int
		expected int
	}{
		{
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			expected: 3,
		},
		{
			target:   10,
			position: []int{3},
			speed:    []int{3},
			expected: 1,
		},
		{
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		actual := carFleet(tt.target, tt.position, tt.speed)
		if actual != tt.expected {
			t.Errorf("carFleet(%d, %v, %v) = %d; expected %d", tt.target, tt.position, tt.speed, actual, tt.expected)
		}
	}
}
