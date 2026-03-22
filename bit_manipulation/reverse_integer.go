package bit_manipulation

import "math"

// Reverse reverses the digits of a 32-bit signed integer; returns 0 on overflow.
// Time Complexity: O(log x)
// Space Complexity: O(1)
func Reverse(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		result = result*10 + digit
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
	}
	return result
}
