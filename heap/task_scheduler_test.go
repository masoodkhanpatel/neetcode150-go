package heap

import "testing"

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []byte
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			tasks:    []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:        2,
			expected: 8,
		},
		{
			name:     "Example 2",
			tasks:    []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:        0,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LeastInterval(tt.tasks, tt.n)
			if result != tt.expected {
				t.Errorf("LeastInterval(%s, %d) = %d; want %d", tt.tasks, tt.n, result, tt.expected)
			}
		})
	}
}
