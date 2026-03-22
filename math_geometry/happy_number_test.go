package math_geometry

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"Happy 19", 19, true},
		{"Not happy 2", 2, false},
		{"Happy 1", 1, true},
		{"Not happy 4", 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHappy(tt.n); got != tt.expected {
				t.Errorf("IsHappy(%d) = %v; want %v", tt.n, got, tt.expected)
			}
		})
	}
}
