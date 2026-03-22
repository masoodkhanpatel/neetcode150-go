package advanced_graphs

import "testing"

func TestAlienOrder(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected string
	}{
		{"Example 1", []string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		{"Invalid cycle", []string{"z", "x", "z"}, ""},
		{"Single word", []string{"z"}, "z"},
		{"Prefix invalid", []string{"abc", "ab"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AlienOrder(tt.words)
			if got != tt.expected {
				t.Errorf("AlienOrder(%v) = %q; want %q", tt.words, got, tt.expected)
			}
		})
	}
}
