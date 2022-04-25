package optimal_account_balancing

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	transactions := [][]int{{0, 1, 10}, {2, 0, 5}}
	require.Equal(t, 2, MinTransfers(transactions))
}

func TestCase2(t *testing.T) {
	transactions := [][]int{{0, 1, 10}, {1, 0, 1}, {1, 3, 5}, {2, 0, 1}}
	require.Equal(t, 3, MinTransfers(transactions))
}

func TestCase3(t *testing.T) {
	transactions := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}}
	require.Equal(t, 0, MinTransfers(transactions))
}

func TestCase4(t *testing.T) {
	transactions := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 4}, {3,4,5},{4, 0, 1}}
	require.Equal(t, 2, MinTransfers(transactions))
}

func TestCase5(t *testing.T) {
	transactions := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 4}, {3,4,5}}
	require.Equal(t, 3, MinTransfers(transactions))
}

func TestCase6(t *testing.T) {
	transactions := [][]int{{10, 11, 6}, {12, 13, 7}, {14, 15, 2}, {14, 16, 2}, {14, 17, 2}, {14, 18, 2}}
	require.Equal(t, 6, MinTransfers(transactions))
}

func TestCase7(t *testing.T) {
	transactions := [][]int{{0, 3, 2}, {1, 4, 3}, {2, 3, 2}, {2, 4, 2}}
	require.Equal(t, 3, MinTransfers(transactions))
}

func TestCase8(t *testing.T) {
	transactions := [][]int{{0, 1, 10}, {1, 0, 1}, {1, 2, 5}, {2, 0, 5}}
	require.Equal(t, 1, MinTransfers(transactions))
}

func TestCase9(t *testing.T) {
	transactions := [][]int{{1, 2, 3}, {1, 3, 3}, {6, 4, 1}, {5, 4, 4}}
	require.Equal(t, 4, MinTransfers(transactions))
}

func TestCase10(t *testing.T) {
	transactions := [][]int{{0, 3, 9}, {1, 4, 2}, {2, 5, 5}, {3, 4, 6}, {4, 5, 2}}
	require.Equal(t, 4, MinTransfers(transactions))
}

func TestCase11(t *testing.T) {
	transactions := [][]int{{1, 8, 1}, {1, 13, 21}, {2, 8, 10}, {3, 9, 20}, {4, 10, 61}, {5, 11, 61}, {6, 12, 59}, {7, 13, 60}}
	require.Equal(t, 8, MinTransfers(transactions))
}
