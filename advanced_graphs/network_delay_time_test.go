package advanced_graphs

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		times    [][]int
		n, k     int
		expected int
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{[][]int{{1, 2, 1}}, 2, 1, 1},
		{[][]int{{1, 2, 1}}, 2, 2, -1},
	}
	for _, tt := range tests {
		if got := NetworkDelayTime(tt.times, tt.n, tt.k); got != tt.expected {
			t.Errorf("NetworkDelayTime(%v, %d, %d) = %d; want %d", tt.times, tt.n, tt.k, got, tt.expected)
		}
	}
}
