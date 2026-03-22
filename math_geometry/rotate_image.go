package math_geometry

// Rotate rotates an N×N matrix 90 degrees clockwise in-place.
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func Rotate(matrix [][]int) {
	n := len(matrix)
	// Transpose
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// Reverse each row
	for i := 0; i < n; i++ {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
		}
	}
}
