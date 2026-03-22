package backtracking

// LetterCombinations returns all possible letter combinations that the number could represent.
// Time Complexity: O(n * 4^n)
// Space Complexity: O(n)
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var res []string
	digitToChar := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack func(int, string)
	backtrack = func(i int, currentStr string) {
		if len(currentStr) == len(digits) {
			res = append(res, currentStr)
			return
		}

		for _, c := range digitToChar[digits[i]] {
			backtrack(i+1, currentStr+string(c))
		}
	}

	backtrack(0, "")
	return res
}
