package coin_change

import (
	"math"
)

func CoinChange(coins []int, amount int) int {
	numOfCoins := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		numOfCoins[i] = math.MaxInt32
	}

	for _, coin := range coins {
		for a := 1; a <= amount; a++ {
			if a == coin {
				numOfCoins[a] = 1
			} else if coin < a {
				numOfCoins[a] = Min(numOfCoins[a], numOfCoins[a-coin]+1)
			}
		}
	}

	if numOfCoins[len(numOfCoins)-1] == math.MaxInt32 {
		return -1
	}

	return numOfCoins[len(numOfCoins)-1]
}

func Min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}