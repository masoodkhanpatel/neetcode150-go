package bit_manipulation

// HammingWeight returns the number of set bits in n.
// Time Complexity: O(log n)
// Space Complexity: O(1)
func HammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}
