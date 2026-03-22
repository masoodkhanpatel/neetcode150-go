package arrays_hashing

// TwoSum returns indices of the two numbers such that they add up to target.
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
}
