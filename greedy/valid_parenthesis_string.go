package greedy

// CheckValidString returns true if the string is valid with '*' acting as '(', ')' or empty.
// Time Complexity: O(n)
// Space Complexity: O(1)
func CheckValidString(s string) bool {
	minOpen, maxOpen := 0, 0
	for _, c := range s {
		switch c {
		case '(':
			minOpen++
			maxOpen++
		case ')':
			minOpen--
			maxOpen--
		case '*':
			minOpen--
			maxOpen++
		}
		if maxOpen < 0 {
			return false
		}
		if minOpen < 0 {
			minOpen = 0
		}
	}
	return minOpen == 0
}
