package arrays_hashing

// ContainsDuplicate returns true if any value appears at least twice in the array.
// Time Complexity: O(n)
// Space Complexity: O(n)
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}
