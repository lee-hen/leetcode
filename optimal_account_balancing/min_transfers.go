package optimal_account_balancing

import (
	"math"
)

func MinTransfers(transactions [][]int) int {
	accountBalances := make(map[int]int)

	for _, transaction := range transactions {
		accountBalances[transaction[0]] -= transaction[2]
		accountBalances[transaction[1]] += transaction[2]
	}

	balances := make([]int, 0)
	for _, balance := range accountBalances {
		balances = append(balances, balance)
	}

	return dfs(0, balances)
}

func dfs(idx int, balances []int) int {
	if idx == len(balances) {
		return 0
	}

	if balances[idx] == 0 {
		return dfs(idx+1, balances)
	}

	transfers := math.MaxInt32
	currDebt := balances[idx]
	for i := idx+1; i < len(balances); i++ {
		if balances[i] * currDebt >= 0 {
			continue
		}

		balances[i] += currDebt
		transfers = min(transfers, 1 + dfs(idx+1, balances))
		balances[i] -= currDebt
	}

	return transfers
}

func min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
