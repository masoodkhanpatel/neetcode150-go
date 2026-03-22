package stack

// Generate Parentheses
func generateParenthesis(n int) []string {
	var res []string
	var backtrack func(open, close int, current string)
	backtrack = func(open, close int, current string) {
		if len(current) == 2*n {
			res = append(res, current)
			return
		}
		if open < n {
			backtrack(open+1, close, current+"(")
		}
		if close < open {
			backtrack(open, close+1, current+")")
		}
	}
	backtrack(0, 0, "")
	return res
}
