package maximize_distance_to_closest_person

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxDistToClosest(t *testing.T) {
	seats := []int{1, 0, 0, 0, 1, 0, 1}
	require.Equal(t, 2, MaxDistToClosest(seats))
	require.Equal(t, 2, maxDistToClosest(seats))
	require.Equal(t, 2, maxDistToClosestTwoPointer(seats))

	seats = []int{1, 0, 0, 0}
	require.Equal(t, 3, MaxDistToClosest(seats))
	require.Equal(t, 3, maxDistToClosest(seats))
	require.Equal(t, 3, maxDistToClosestTwoPointer(seats))

	seats = []int{0, 1}
	require.Equal(t, 1, MaxDistToClosest(seats))
	require.Equal(t, 1, maxDistToClosest(seats))
	require.Equal(t, 1, maxDistToClosestTwoPointer(seats))
}
