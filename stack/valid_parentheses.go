package stack

// Valid Parentheses
func isValid(s string) bool {
	stack := []rune{}
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if val, ok := m[char]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == val {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
