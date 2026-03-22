package advanced_graphs

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, 20},
		{[][]int{{3, 12}, {-2, 5}, {-4, 1}}, 18},
		{[][]int{{0, 0}}, 0},
	}
	for _, tt := range tests {
		if got := MinCostConnectPoints(tt.points); got != tt.expected {
			t.Errorf("MinCostConnectPoints(%v) = %d; want %d", tt.points, got, tt.expected)
		}
	}
}
