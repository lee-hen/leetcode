package line_reflection

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	require.True(t, IsReflected1([][]int{{1, 1}, {-1, 1}}))
	require.True(t, IsReflected([][]int{{1, 1}, {-1, 1}}))
}

func TestCase2(t *testing.T) {
	require.False(t, IsReflected1([][]int{{0, 0}, {1, 0}, {3, 0}}))
	require.False(t, IsReflected([][]int{{0, 0}, {1, 0}, {3, 0}}))
}

func TestCase3(t *testing.T) {
	require.True(t, IsReflected1([][]int{{1, 1}, {0, 0}, {-1, 1}}))
	require.True(t, IsReflected([][]int{{1, 1}, {0, 0}, {-1, 1}}))
}

func TestCase4(t *testing.T) {
	require.False(t, IsReflected1([][]int{{1, 1}, {-1, -1}}))
	require.False(t, IsReflected([][]int{{1, 1}, {-1, -1}}))
}

func TestCase5(t *testing.T) {
	require.True(t, IsReflected1([][]int{{1,2},{2,2},{1,4},{2,4}}))
	require.True(t, IsReflected([][]int{{1,2},{2,2},{1,4},{2,4}}))
}
