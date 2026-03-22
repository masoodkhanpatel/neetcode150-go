package greedy

import "sort"

// IsNStraightHand returns true if hand can be arranged into groups of groupSize consecutive cards.
// Time Complexity: O(n log n)
// Space Complexity: O(n)
func IsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	count := make(map[int]int)
	for _, c := range hand {
		count[c]++
	}

	keys := make([]int, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		if count[k] > 0 {
			freq := count[k]
			for i := 0; i < groupSize; i++ {
				if count[k+i] < freq {
					return false
				}
				count[k+i] -= freq
			}
		}
	}
	return true
}
