package jewels_and_stones

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	jewels := "aA"
	stones := "aAAbbbb"
	require.Equal(t, 3, NumJewelsInStones(jewels, stones))

	jewels = "z"
	stones = "ZZ"
	require.Equal(t, 0, NumJewelsInStones(jewels, stones))
}
