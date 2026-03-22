package math_geometry

// PlusOne adds one to the number represented as a digit array.
// Time Complexity: O(n)
// Space Complexity: O(n)
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
