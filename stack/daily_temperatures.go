package stack

// Daily Temperatures
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{} // stores indices

	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = i - idx
		}
		stack = append(stack, i)
	}

	return res
}
