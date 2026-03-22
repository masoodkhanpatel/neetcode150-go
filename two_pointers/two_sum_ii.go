package two_pointers

// TwoSumII returns the indices of the two numbers such that they add up to target.
// Time Complexity: O(n)
// Space Complexity: O(1)
func TwoSumII(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]

		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
