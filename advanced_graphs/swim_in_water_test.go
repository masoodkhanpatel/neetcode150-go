package advanced_graphs

import "testing"

func TestSwimInWater(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{0, 2}, {1, 3}}, 3},
		{[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, 16},
	}
	for _, tt := range tests {
		if got := SwimInWater(tt.grid); got != tt.expected {
			t.Errorf("SwimInWater(%v) = %d; want %d", tt.grid, got, tt.expected)
		}
	}
}
