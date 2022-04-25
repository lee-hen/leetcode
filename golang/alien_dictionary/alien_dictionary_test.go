package alien_dictionary

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlienOrder(t *testing.T) {
	words := []string{"wrt","wrf","er","ett","rftt"}
	actual := AlienOrder(words)
	require.Equal(t, "wertf", actual)
	require.Equal(t, "wertf", alienOrder(words))
}

func TestAlienOrder1(t *testing.T) {
	words := []string{"z","x"}
	actual := AlienOrder(words)
	require.Equal(t, "zx", actual)
	require.Equal(t, "zx", alienOrder(words))
}

func TestAlienOrder2(t *testing.T) {
	words := []string{"z","x","z"}
	actual := AlienOrder(words)
	require.Equal(t, "", actual)
	require.Equal(t, "", alienOrder(words))
}

func TestAlienOrder3(t *testing.T) {
	words := []string{"abc","ab"}
	actual := AlienOrder(words)
	require.Equal(t, "", actual)
	require.Equal(t, "", alienOrder(words))
}

func TestAlienOrder4 (t *testing.T) {
	words := []string{"abbb","abc"}
	actual1 := AlienOrder(words)
	if "bca" == actual1 {
		require.Equal(t, "bca", actual1)
	} else if "abc" == actual1 {
		require.Equal(t, "abc", actual1)
	} else {
		require.Equal(t, "bac", actual1)
	}

	require.NotEqual(t, "acb", actual1)
	require.NotEqual(t, "cba", actual1)
	require.NotEqual(t, "cab", actual1)

	actual2 := alienOrder(words)
	if "bca" == actual2 {
		require.Equal(t, "bca", actual2)
	} else if "abc" == actual2 {
		require.Equal(t, "abc", actual2)
	} else {
		require.Equal(t, "bac", actual1)
	}

	require.NotEqual(t, "acb", actual2)
	require.NotEqual(t, "cba", actual2)
	require.NotEqual(t, "cab", actual2)
}

func TestAlienOrder5(t *testing.T) {
	words := []string{"z","z"}
	actual := AlienOrder(words)
	require.Equal(t, "z", actual)
	require.Equal(t, "z", alienOrder(words))
}
