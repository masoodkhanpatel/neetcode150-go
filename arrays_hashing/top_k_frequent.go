package arrays_hashing

// TopKFrequent returns the k most frequent elements.
// Time Complexity: O(n)
// Space Complexity: O(n)
func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	freq := make([][]int, len(nums)+1)
	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	var res []int
	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
