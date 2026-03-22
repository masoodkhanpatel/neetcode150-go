package stack

import "strconv"

// Evaluate Reverse Polish Notation
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch t {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		} else {
			val, _ := strconv.Atoi(t)
			stack = append(stack, val)
		}
	}
	return stack[0]
}
