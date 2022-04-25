package jump_game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	require.True(t, CanJump(nums))

	nums = []int{3, 2, 1, 0, 4}
	require.False(t, CanJump(nums))

	nums = []int{0}
	require.True(t, CanJump(nums))

	nums = []int{1, 2, 3}
	require.True(t, CanJump(nums))

	nums = []int{1,1,1,0}
	require.True(t, CanJump(nums))
}

func Test_canJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	require.True(t, canJump(nums))

	nums = []int{3, 2, 1, 0, 4}
	require.False(t, canJump(nums))

	nums = []int{0}
	require.True(t, canJump(nums))

	nums = []int{1, 2, 3}
	require.True(t, canJump(nums))

	nums = []int{1,1,1,0}
	require.True(t, canJump(nums))
}
