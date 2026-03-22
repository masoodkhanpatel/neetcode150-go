package math_geometry

// IsHappy returns true if n is a happy number.
// Time Complexity: O(log n)
// Space Complexity: O(1)
func IsHappy(n int) bool {
	slow, fast := n, sumOfSquares(n)
	for fast != 1 && slow != fast {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))
	}
	return fast == 1
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}
