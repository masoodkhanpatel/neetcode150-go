package math_geometry

// MyPow computes x raised to the power n using fast exponentiation.
// Time Complexity: O(log n)
// Space Complexity: O(1)
func MyPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result
}
