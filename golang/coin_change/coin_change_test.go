package coin_change

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	require.Equal(t, 3, CoinChange(coins, amount))

	coins = []int{2}
	amount = 3
	require.Equal(t, -1, CoinChange(coins, amount))
}
