package trapping_rain_water

import (
	"github.com/stretchr/testify/require"
	"testing"
)


func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	require.Equal(t, 6, Trap(height))

	height = []int{4, 2, 0, 3, 2, 5}
	require.Equal(t, 9, Trap(height))

	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 3, 4, 5, 6, 0, 1, 2}
	require.Equal(t, 14, Trap(height))
}


func Test_trap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	require.Equal(t, 6, trap(height))

	height = []int{4, 2, 0, 3, 2, 5}
	require.Equal(t, 9, trap(height))

	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 3, 4, 5, 6, 0, 1, 2}
	require.Equal(t, 14, trap(height))
}
