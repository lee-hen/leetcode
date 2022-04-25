package container_with_most_water

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{1,8,6,2,5,4,8,3,7}
	require.Equal(t, 49, MaxArea(height))

	height = []int{1, 1}
	require.Equal(t, 1, MaxArea(height))

	height = []int{1, 1, 1, 1, 2, 2, 1, 1, 1}
	require.Equal(t, 8, MaxArea(height))
}

func Test_maxArea(t *testing.T) {
	height := []int{1,8,6,2,5,4,8,3,7}
	require.Equal(t, 49, maxArea(height))

	height = []int{1, 1}
	require.Equal(t, 1, maxArea(height))

	height = []int{1, 1, 1, 1, 2, 2, 1, 1, 1}
	require.Equal(t, 8, maxArea(height))
}
