package valid_word_abbr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidWordAbbr_IsUnique(t *testing.T) {
	validWordAbbr := Constructor([]string{"deer", "door", "cake", "card"})
	strings := []string{"dear", "cart", "cane", "make", "cake"}
	expects := []bool{false, true, false, true, true}
	actual := make([]bool, 0)
	for _, word := range strings {
		actual = append(actual, validWordAbbr.IsUnique(word))
	}
	require.ElementsMatch(t, expects, actual)
}

func TestSpecialCase(t *testing.T) {
	validWordAbbr := Constructor([]string{"a", "a"})
	require.True(t, validWordAbbr.IsUnique("a"))
}
