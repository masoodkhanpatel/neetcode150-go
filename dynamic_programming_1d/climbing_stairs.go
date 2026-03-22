package dynamic_programming_1d

// ClimbStairs returns the number of distinct ways to climb n stairs (1 or 2 steps).
// Time Complexity: O(n)
// Space Complexity: O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
