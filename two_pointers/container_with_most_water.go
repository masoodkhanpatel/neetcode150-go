package two_pointers

// MaxArea finds two lines that together with the x-axis form a container, such that the container contains the most water.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0

	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
