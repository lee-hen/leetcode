package best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	require.Equal(t, 5, MaxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	require.Equal(t, 0, MaxProfit(prices))
}
