package best_time_to_buy_and_sell_stock

import "math"

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxSoFar := math.MinInt32
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		maxSoFar = max(maxSoFar, -prices[i-1])
		maxProfit = max(maxProfit, maxSoFar+prices[i])
	}

	return maxProfit
}

func max(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
}

