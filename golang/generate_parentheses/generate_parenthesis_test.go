package generate_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	require.Equal(t, generateParenthesis(3), GenerateParenthesis(3))
	require.Equal(t, generateParenthesis(1), generateParenthesis(1))
	require.ElementsMatch(t, generateParenthesis(4), otherSolution(4))
}
