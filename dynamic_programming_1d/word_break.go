package dynamic_programming_1d

// WordBreak returns true if s can be segmented into words from wordDict.
// Time Complexity: O(n^2 * m) where m is average word length
// Space Complexity: O(n)
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
