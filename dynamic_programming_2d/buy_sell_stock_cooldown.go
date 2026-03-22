package dynamic_programming_2d

// MaxProfitCooldown returns max profit with a cooldown day after each sell.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxProfitCooldown(prices []int) int {
	// States: holding, sold (cooldown), idle
	holding, sold, idle := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		prevHolding, prevSold, prevIdle := holding, sold, idle
		holding = max2d(prevHolding, prevIdle-prices[i])
		sold = prevHolding + prices[i]
		idle = max2d(prevIdle, prevSold)
	}
	return max2d(sold, idle)
}

func max2d(a, b int) int {
	if a > b {
		return a
	}
	return b
}
