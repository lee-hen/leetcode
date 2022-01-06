package k_closest_points_to_origin

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKClosest(t *testing.T) {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	require.ElementsMatch(t, [][]int{{-2, 2}}, KClosest(points, k))

	points = [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k = 2
	require.ElementsMatch(t, [][]int{{-2, 4}, {3, 3}}, KClosest(points, k))
}

func Test_kClosest(t *testing.T) {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	require.ElementsMatch(t, [][]int{{-2, 2}}, kClosest(points, k))

	points = [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k = 2
	require.ElementsMatch(t, [][]int{{-2, 4}, {3, 3}}, kClosest(points, k))
}
