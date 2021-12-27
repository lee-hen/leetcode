package odd_even_jump

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOddEvenJumps(t *testing.T) {
	arr := []int{10,13,12,14,15}
	require.Equal(t, 2, OddEvenJumps(arr))

	arr = []int{2,3,1,1,4}
	require.Equal(t, 3, OddEvenJumps(arr))

	arr = []int{5,1,3,4,2}
	require.Equal(t, 3, OddEvenJumps(arr))

	arr = []int{1, 2, 3, 2, 1, 4, 4, 5}
	require.Equal(t, 6, OddEvenJumps(arr))
}

func Test_oddEvenJumps(t *testing.T) {
	arr := []int{10,13,12,14,15}
	require.Equal(t, 2, oddEvenJumps(arr))

	arr = []int{2,3,1,1,4}
	require.Equal(t, 3, oddEvenJumps(arr))

	arr = []int{5,1,3,4,2}
	require.Equal(t, 3, oddEvenJumps(arr))

	arr = []int{1, 2, 3, 2, 1, 4, 4, 5}
	require.Equal(t, 6, oddEvenJumps(arr))
}
