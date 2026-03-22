package backtracking

// Partition returns every possible palindrome partitioning of s.
// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)
func Partition(s string) [][]string {
	var res [][]string
	var part []string

	var dfs func(int)
	dfs = func(i int) {
		if i >= len(s) {
			temp := make([]string, len(part))
			copy(temp, part)
			res = append(res, temp)
			return
		}

		for j := i; j < len(s); j++ {
			if isPalindrome(s, i, j) {
				part = append(part, s[i:j+1])
				dfs(j + 1)
				part = part[:len(part)-1]
			}
		}
	}

	dfs(0)
	return res
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}
