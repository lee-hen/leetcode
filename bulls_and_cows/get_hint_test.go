package bulls_and_cows

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetHint(t *testing.T) {
	secret := "1807"
	guess := "7810"
	require.Equal(t, getHint(secret, guess), GetHint(secret, guess))

	secret = "1123"
	guess = "0111"
	require.Equal(t, getHint(secret, guess), GetHint(secret, guess))
}
