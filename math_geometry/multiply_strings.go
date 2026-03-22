package math_geometry

import "strconv"

// Multiply returns the product of two non-negative integers given as strings.
// Time Complexity: O(m*n)
// Space Complexity: O(m+n)
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	pos := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1, p2 := i+j, i+j+1
			sum := mul + pos[p2]
			pos[p2] = sum % 10
			pos[p1] += sum / 10
		}
	}

	result := ""
	for _, v := range pos {
		if !(result == "" && v == 0) {
			result += strconv.Itoa(v)
		}
	}
	return result
}
