package rotate_image

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Rotate(matrix)
	require.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, matrix)
}
