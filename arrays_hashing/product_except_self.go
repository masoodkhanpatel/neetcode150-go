package arrays_hashing

// ProductExceptSelf returns an array such that answer[i] is equal to the product of all the elements of nums except nums[i].
// Time Complexity: O(n)
// Space Complexity: O(1) (excluding output array)
func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	// Calculate prefix products
	prefix := 1
	for i := 0; i < length; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// Calculate suffix products and multiply with prefix
	suffix := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}
