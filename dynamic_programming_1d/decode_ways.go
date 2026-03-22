package dynamic_programming_1d

// NumDecodings returns the number of ways to decode the digit string.
// Time Complexity: O(n)
// Space Complexity: O(1)
func NumDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	prev2, prev1 := 1, 1
	for i := 1; i < len(s); i++ {
		cur := 0
		if s[i] != '0' {
			cur += prev1
		}
		two := (int(s[i-1]-'0'))*10 + int(s[i]-'0')
		if two >= 10 && two <= 26 {
			cur += prev2
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}
