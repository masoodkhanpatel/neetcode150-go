package greedy

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"No solution", []int{2, 3, 4}, []int{3, 4, 3}, -1},
		{"Single station", []int{5}, []int{4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanCompleteCircuit(tt.gas, tt.cost); got != tt.expected {
				t.Errorf("CanCompleteCircuit(%v, %v) = %d; want %d", tt.gas, tt.cost, got, tt.expected)
			}
		})
	}
}
