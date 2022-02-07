package course_schedule2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindOrder(t *testing.T) {
	require.Equal(t, []int{0, 1, 2, 3}, FindOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
