package math_geometry

// SpiralOrder returns all elements of the matrix in spiral order.
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func SpiralOrder(matrix [][]int) []int {
	result := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}
	return result
}
