package sliding_window

// Sliding Window Maximum
// You are given an array of integers nums, 
// there is a sliding window of size k which is moving from the very left of the array to the very right. 
// You can only see the k numbers in the window. 
// Each time the sliding window moves right by one position.

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var q []int // indices

	l, r := 0, 0
	for r < len(nums) {
		// pop smaller values from q
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)

		// remove left value from window
		if l > q[0] {
			q = q[1:]
		}

		if (r + 1) >= k {
			res = append(res, nums[q[0]])
			l++
		}
		r++
	}

	return res
}
