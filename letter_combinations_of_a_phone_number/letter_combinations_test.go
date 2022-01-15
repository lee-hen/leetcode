package letter_combinations_of_a_phone_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	digits := "234"
	require.Equal(t, letterCombinations(digits), LetterCombinations(digits))

	digits = ""
	require.Equal(t, letterCombinations(digits), LetterCombinations(digits))

	digits = "2"
	require.Equal(t, letterCombinations(digits), LetterCombinations(digits))
}
