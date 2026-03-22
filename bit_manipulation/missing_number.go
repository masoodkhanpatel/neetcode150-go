package bit_manipulation

// MissingNumber returns the missing number in the range [0, n].
// Time Complexity: O(n)
// Space Complexity: O(1)
func MissingNumber(nums []int) int {
	n := len(nums)
	result := n
	for i, v := range nums {
		result ^= i ^ v
	}
	return result
}
