package greedy

// PartitionLabels partitions the string so each letter appears in at most one part.
// Returns the sizes of each part.
// Time Complexity: O(n)
// Space Complexity: O(1)
func PartitionLabels(s string) []int {
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i
	}

	result := []int{}
	start, end := 0, 0
	for i, c := range s {
		if last[c-'a'] > end {
			end = last[c-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	return result
}
