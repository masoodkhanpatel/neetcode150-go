package bit_manipulation

import "testing"

func TestGetSum(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"1+2", 1, 2, 3},
		{"2+3", 2, 3, 5},
		{"Negative", -1, 1, 0},
		{"Both negative", -3, -2, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSum(tt.a, tt.b); got != tt.expected {
				t.Errorf("GetSum(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
