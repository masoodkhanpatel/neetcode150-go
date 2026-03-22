package arrays_hashing

import (
	"sort"
	"strings"
)

// GroupAnagrams groups anagrams together.
// Time Complexity: O(m * n * log(n)) where m is the number of strings and n is the length of the longest string.
// Space Complexity: O(m * n)
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := sortString(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
