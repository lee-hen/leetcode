package design_search_autocomplete_system

import (
	"strings"
)

type Trie struct {
	root *TrieNode
}

func NewTrie(list []string, times []int) *Trie {
	root := new(TrieNode)
	root.children = make(map[byte]*TrieNode)
	for i, word := range list {
		root.Insert(word, times[i])
	}
	return &Trie{
		root: root,
	}
}

func (t *Trie) search(prefix string) []HotSentence {
	node := t.root

	path := new(strings.Builder)
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		node = node.children[c]

		if node == nil {
			return []HotSentence{}
		}
		path.WriteByte(c)
	}

	return node.allPaths(path.String())
}
