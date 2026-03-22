package binary_search

// Find Minimum in Rotated Sorted Array
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	res := nums[0]

	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < res {
				res = nums[l]
			}
			break
		}

		m := l + (r-l)/2
		if nums[m] < res {
			res = nums[m]
		}

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}
