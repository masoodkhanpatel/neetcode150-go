package two_pointers

import (
	"unicode"
)

// IsPalindrome returns true if the string is a palindrome.
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		leftRune := rune(s[l])
		rightRune := rune(s[r])

		if !isAlphanumeric(leftRune) {
			l++
			continue
		}
		if !isAlphanumeric(rightRune) {
			r--
			continue
		}

		if unicode.ToLower(leftRune) != unicode.ToLower(rightRune) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
