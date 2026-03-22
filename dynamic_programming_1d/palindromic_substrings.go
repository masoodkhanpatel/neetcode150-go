package dynamic_programming_1d

// CountSubstrings counts the number of palindromic substrings.
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func CountSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		// Odd length
		for l, r := i, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			count++
		}
		// Even length
		for l, r := i, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			count++
		}
	}
	return count
}
