package bit_manipulation

// SingleNumber returns the element that appears only once; all others appear twice.
// Time Complexity: O(n)
// Space Complexity: O(1)
func SingleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}
