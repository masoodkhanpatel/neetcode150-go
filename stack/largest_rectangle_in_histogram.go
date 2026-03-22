package stack

// Largest Rectangle in Histogram
func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []int{} // stores indices

	for i, h := range heights {
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			if height*width > maxArea {
				maxArea = height * width
			}
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		width := len(heights)
		if len(stack) > 0 {
			width = len(heights) - stack[len(stack)-1] - 1
		}
		if height*width > maxArea {
			maxArea = height * width
		}
	}

	return maxArea
}
