package bit_manipulation

// GetSum returns the sum of two integers without using the + operator.
// Time Complexity: O(1)
// Space Complexity: O(1)
func GetSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}
