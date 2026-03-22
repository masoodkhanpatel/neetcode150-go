package sliding_window

// LengthOfLongestSubstring returns the length of the longest substring without repeating characters.
// Time Complexity: O(n)
// Space Complexity: O(min(m, n)) where n is the size of the string and m is the size of the charset/alphabet
func LengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	l := 0
	res := 0

	for r := 0; r < len(s); r++ {
		for charSet[s[r]] {
			delete(charSet, s[l])
			l++
		}
		charSet[s[r]] = true
		if (r - l + 1) > res {
			res = r - l + 1
		}
	}
	return res
}
