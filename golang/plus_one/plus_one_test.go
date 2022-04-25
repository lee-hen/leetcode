package plus_one

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits := []int{1, 2, 3}
	require.Equal(t, []int{1, 2, 4}, PlusOne(digits))

	digits = []int{4, 3, 2, 1}
	require.Equal(t, []int{4, 3, 2, 2}, PlusOne(digits))

	digits = []int{9}
	require.Equal(t, []int{1, 0}, PlusOne(digits))
}

func Test_plusOne(t *testing.T) {
	digits := []int{1, 2, 3}
	require.Equal(t, []int{1, 2, 4}, plusOne(digits))

	digits = []int{4, 3, 2, 1}
	require.Equal(t, []int{4, 3, 2, 2}, plusOne(digits))

	digits = []int{9}
	require.Equal(t, []int{1, 0}, plusOne(digits))
}
