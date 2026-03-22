package dynamic_programming_1d

// LongestPalindrome returns the longest palindromic substring.
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func LongestPalindrome(s string) string {
	start, maxLen := 0, 1

	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxLen {
				maxLen = r - l + 1
				start = l
			}
			l--
			r++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)   // odd length
		expand(i, i+1) // even length
	}
	return s[start : start+maxLen]
}
