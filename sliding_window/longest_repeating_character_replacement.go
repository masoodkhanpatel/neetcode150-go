package sliding_window

// CharacterReplacement returns the length of the longest substring containing the same letter you can get after replacing at most k characters.
// Time Complexity: O(n)
// Space Complexity: O(1) (since the alphabet size is constant, 26)
func CharacterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res := 0
	l := 0
	maxf := 0

	for r := 0; r < len(s); r++ {
		count[s[r]]++
		if count[s[r]] > maxf {
			maxf = count[s[r]]
		}

		if (r-l+1)-maxf > k {
			count[s[l]]--
			l++
		}

		if (r - l + 1) > res {
			res = r - l + 1
		}
	}
	return res
}
