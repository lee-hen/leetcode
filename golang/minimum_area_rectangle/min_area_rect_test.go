package minimum_area_rectangle

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinAreaRect(t *testing.T) {
	points := [][]int{{1,1},{1,3},{3,1},{3,3},{4,1},{4,3}}
	require.Equal(t, 2, MinAreaRect(points))
}
