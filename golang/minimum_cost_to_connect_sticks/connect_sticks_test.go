package minimum_cost_to_connect_sticks

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	require.Equal(t, 14, ConnectSticks([]int{2, 4, 3}))
}

func TestCase2(t *testing.T) {
	require.Equal(t, 30, ConnectSticks([]int{1, 8, 3, 5}))
}

func TestCase3(t *testing.T) {
	require.Equal(t, 0, ConnectSticks([]int{5}))
}

func TestCase4(t *testing.T) {
	require.Equal(t, 73, ConnectSticks([]int{1,8,3,5,6,7}))
}