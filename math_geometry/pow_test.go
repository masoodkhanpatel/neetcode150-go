package math_geometry

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{"2^10", 2.0, 10, 1024.0},
		{"2.1^3", 2.10000, 3, 9.261},
		{"2^-2", 2.0, -2, 0.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MyPow(tt.x, tt.n)
			if math.Abs(got-tt.expected) > 1e-3 {
				t.Errorf("MyPow(%v, %d) = %v; want %v", tt.x, tt.n, got, tt.expected)
			}
		})
	}
}
