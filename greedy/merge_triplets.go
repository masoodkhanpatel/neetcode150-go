package greedy

// MergeTriplets returns true if we can choose some triplets and merge (take max per position) to get target.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MergeTriplets(triplets [][]int, target []int) bool {
	result := [3]int{}
	for _, t := range triplets {
		// Skip triplets that exceed target in any position
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue
		}
		for i := 0; i < 3; i++ {
			if t[i] > result[i] {
				result[i] = t[i]
			}
		}
	}
	return result[0] == target[0] && result[1] == target[1] && result[2] == target[2]
}
