package math_geometry

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		num1     string
		num2     string
		expected string
	}{
		{"2*3", "2", "3", "6"},
		{"123*456", "123", "456", "56088"},
		{"Zero", "0", "123", "0"},
		{"Single digits", "9", "9", "81"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.num1, tt.num2); got != tt.expected {
				t.Errorf("Multiply(%q, %q) = %q; want %q", tt.num1, tt.num2, got, tt.expected)
			}
		})
	}
}
