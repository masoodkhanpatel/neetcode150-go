package arrays_hashing

// IsAnagram returns true if t is an anagram of s, and false otherwise.
// Time Complexity: O(n)
// Space Complexity: O(1) (since the alphabet size is constant, 26)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make([]int, 26)

	for i := 0; i < len(s); i++ {
		charCount[s[i]-'a']++
		charCount[t[i]-'a']--
	}

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}
