package sliding_window

// MaxProfit returns the maximum profit you can achieve from this transaction.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > maxP {
				maxP = profit
			}
		} else {
			l = r
		}
		r++
	}
	return maxP
}
