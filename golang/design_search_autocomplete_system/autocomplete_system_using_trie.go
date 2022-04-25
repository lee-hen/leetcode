package design_search_autocomplete_system

import (
	"sort"
	"strings"
)

type AutocompleteSystem struct {
	trie *Trie
	current *strings.Builder
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	return AutocompleteSystem{
		trie: NewTrie(sentences, times),
		current: new(strings.Builder),
	}
}

func (a *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		a.trie.root.Insert(a.current.String(), 1)
		a.current.Reset()
		return []string{}
	}

	a.current.WriteByte(c)
	hotSentences := a.trie.search(a.current.String())
	SortHotSentences(hotSentences)
	sentences := make([]string, 0)
	for i := 0; i < 3 && i < len(hotSentences); i++ {
		sentences = append(sentences, hotSentences[i].word)
	}
	return sentences
}

func SortHotSentences(hotSentences []HotSentence) {
	sort.Slice(hotSentences, func(i, j int) bool {
		if hotSentences[i].hot == hotSentences[j].hot {
			return hotSentences[i].word < hotSentences[j].word
		}
		return hotSentences[i].hot > hotSentences[j].hot
	})
}
