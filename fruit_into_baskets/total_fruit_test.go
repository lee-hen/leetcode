package fruit_into_baskets

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTotalFruit(t *testing.T) {
	fruits := []int{1, 1, 1, 4, 1, 1, 2, 2, 2, 4, 4, 1, 3, 4, 6, 7}
	require.Equal(t, 6, TotalFruit(fruits))

	fruits = []int{1, 2, 3, 4, 1, 1, 2, 2, 2, 4, 4, 1, 3, 4, 6, 7}
	require.Equal(t, 5, TotalFruit(fruits))

	fruits = []int{1, 2, 3, 2, 2}
	require.Equal(t, 4, TotalFruit(fruits))

	fruits = []int{0, 1, 2, 2}
	require.Equal(t, 3, TotalFruit(fruits))

	fruits = []int{1, 2, 1}
	require.Equal(t, 3, TotalFruit(fruits))
}

func Test_totalFruit(t *testing.T) {
	fruits := []int{1, 1, 1, 4, 1, 1, 2, 2, 2, 4, 4, 1, 3, 4, 6, 7}
	require.Equal(t, 6, totalFruit(fruits))

	fruits = []int{1, 2, 3, 4, 1, 1, 2, 2, 2, 4, 4, 1, 3, 4, 6, 7}
	require.Equal(t, 5, totalFruit(fruits))

	fruits = []int{1, 2, 3, 2, 2}
	require.Equal(t, 4, totalFruit(fruits))

	fruits = []int{0, 1, 2, 2}
	require.Equal(t, 3, totalFruit(fruits))

	fruits = []int{1, 2, 1}
	require.Equal(t, 3, totalFruit(fruits))
}
