package two_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []int{0, 1}
	output := twoSum([]int{2, 7, 11, 15}, 9)
	require.ElementsMatch(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []int{1, 2}
	output := twoSum([]int{3, 2, 4}, 6)
	require.ElementsMatch(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := []int{0, 1}
	output := twoSum([]int{3, 3}, 6)
	require.ElementsMatch(t, expected, output)
}
