package arrays_hashing

// LongestConsecutive returns the length of the longest consecutive elements sequence.
// Time Complexity: O(n)
// Space Complexity: O(n)
func LongestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest := 0
	for num := range set {
		// Check if it's the start of a sequence
		if _, ok := set[num-1]; !ok {
			length := 1
			for {
				if _, ok := set[num+length]; ok {
					length++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
