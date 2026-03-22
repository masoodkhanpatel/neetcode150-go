package arrays_hashing

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:  "Example 1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name:     "Empty string",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Single character",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupAnagrams(tt.input)

			// Sort inner slices and outer slice for comparison because map iteration order is random
			for _, group := range result {
				sort.Strings(group)
			}
			sort.Slice(result, func(i, j int) bool {
				if len(result[i]) != len(result[j]) {
					return len(result[i]) < len(result[j])
				}
				// If lengths are equal, sort by the first element of the group
				if len(result[i]) > 0 && len(result[j]) > 0 {
					return result[i][0] < result[j][0]
				}
				return true
			})

			for _, group := range tt.expected {
				sort.Strings(group)
			}
			sort.Slice(tt.expected, func(i, j int) bool {
				if len(tt.expected[i]) != len(tt.expected[j]) {
					return len(tt.expected[i]) < len(tt.expected[j])
				}
				if len(tt.expected[i]) > 0 && len(tt.expected[j]) > 0 {
					return tt.expected[i][0] < tt.expected[j][0]
				}
				return true
			})

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GroupAnagrams(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
