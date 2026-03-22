package bit_manipulation

// ReverseBits reverses the bits of a 32-bit unsigned integer.
// Time Complexity: O(1)
// Space Complexity: O(1)
func ReverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}
