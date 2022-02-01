package design_search_autocomplete_system

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_autocompleteSystem(t *testing.T) {
	a := constructor([]string{"i love you", "island", "iroman", "i love leetcode"}, []int{5, 3, 2, 2})

	require.Equal(t, []string{"i love you","island","i love leetcode"}, a.Input('i'))
	require.Equal(t, []string{"i love you","i love leetcode"}, a.Input(' '))
	require.Equal(t, []string{}, a.Input('a'))
	require.Equal(t, []string{}, a.Input('#'))
	require.Equal(t, []string{"i love you","island","i love leetcode"}, a.Input('i'))
	require.Equal(t, []string{"i love you","i love leetcode","i a"}, a.Input(' '))
	require.Equal(t, []string{"i a"}, a.Input('a'))
	require.Equal(t, []string{}, a.Input('#'))
	require.Equal(t, []string{"i love you", "island", "i a"}, a.Input('i'))
	require.Equal(t, []string{"i love you","i a","i love leetcode"}, a.Input(' '))
	require.Equal(t, []string{"i a"}, a.Input('a'))
	require.Equal(t, []string{}, a.Input('#'))
}

func TestAutocompleteSystem(t *testing.T) {
	a := Constructor([]string{"abc", "abbc", "a"}, []int{3, 3, 3})

	require.Equal(t, []string{}, a.Input('b'))
	require.Equal(t, []string{}, a.Input('c'))
	require.Equal(t, []string{}, a.Input('#'))
	require.Equal(t, []string{"bc"}, a.Input('b'))
	require.Equal(t, []string{"bc"}, a.Input('c'))
	require.Equal(t, []string{}, a.Input('#'))
	require.Equal(t, []string{"a", "abbc", "abc"}, a.Input('a'))
	require.Equal(t, []string{"abbc", "abc"}, a.Input('b'))
	require.Equal(t, []string{"abc"}, a.Input('c'))
	require.Equal(t, []string{}, a.Input('#'))
	require.Equal(t, []string{"abc", "a", "abbc"}, a.Input('a'))
	require.Equal(t, []string{"abc", "abbc"}, a.Input('b'))
	require.Equal(t, []string{"abc"}, a.Input('c'))
	require.Equal(t, []string{}, a.Input('#'))
}
