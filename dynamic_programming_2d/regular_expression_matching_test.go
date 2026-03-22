package dynamic_programming_2d

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s, p     string
		expected bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
	}
	for _, tt := range tests {
		if got := IsMatch(tt.s, tt.p); got != tt.expected {
			t.Errorf("IsMatch(%q, %q) = %v; want %v", tt.s, tt.p, got, tt.expected)
		}
	}
}
