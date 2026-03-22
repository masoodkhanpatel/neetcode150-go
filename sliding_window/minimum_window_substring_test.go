package sliding_window

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}

	for _, tt := range tests {
		got := minWindow(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("minWindow(%s, %s) = %s; want %s", tt.s, tt.t, got, tt.want)
		}
	}
}
