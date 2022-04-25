package reverse_integer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverse(t *testing.T) {
	x := 123
	require.Equal(t, reverse(x), Reverse(x))

	x = -123
	require.Equal(t, reverse(x), Reverse(x))

	x = 120
	require.Equal(t, reverse(x), Reverse(x))
}
