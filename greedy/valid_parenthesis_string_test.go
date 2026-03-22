package greedy

import "testing"

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"Valid with star", "()", true},
		{"Star as empty", "(*)", true},
		{"Star as paren", "(*))", true},
		{"Invalid", "((", false},
		{"All stars", "***", true},
		{"Complex valid", "(((******)))", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidString(tt.s); got != tt.expected {
				t.Errorf("CheckValidString(%q) = %v; want %v", tt.s, got, tt.expected)
			}
		})
	}
}
