package peak_index_in_a_mountain_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	require.Equal(t, peakIndexInMountainArray([]int{0, 2, 1, 0}), PeakIndexInMountainArray([]int{0, 2, 1, 0}))
}
