package greedy

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []int
	}{
		{"Example 1", "ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"Example 2", "eccbbbbdec", []int{10}},
		{"Single char", "a", []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartitionLabels(tt.s); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("PartitionLabels(%q) = %v; want %v", tt.s, got, tt.expected)
			}
		})
	}
}
