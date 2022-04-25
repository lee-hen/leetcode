package missing_ranges

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findMissingRanges(t *testing.T) {
	nums := []int{-3, -2, -1, 0, 1, 3, 50, 75, 101}
	lower, upper := 0, 99
	expected := []string{"2","4->49","51->74","76->99"}
	require.Equal(t, expected, findMissingRanges(nums, lower, upper))
}

func TestFindMissingRanges(t *testing.T) {
	nums := []int{-3, -2, -1, 0, 1, 3, 50, 75}
	lower, upper := 0, 99
	expected := []string{"2","4->49","51->74","76->99"}
	require.Equal(t, expected, FindMissingRanges(nums, lower, upper))
}