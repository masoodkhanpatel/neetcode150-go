package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, tt := range tests {
		got := isValid(tt.s)
		if got != tt.want {
			t.Errorf("isValid(%s) = %v; want %v", tt.s, got, tt.want)
		}
	}
}
