package lib

type Trie struct {
	root *TrieNode
}

func NewTrie(list []string) *Trie {
	root := InitTrieNode()

	for _, word := range list {
		root.AddWord(word)
	}

	return &Trie{
		root: root,
	}
}

func (t *Trie) ContainsExactly(prefix string, exact bool) bool {
	lastNode := t.root

	for i := 0; i < len(prefix); i++ {
		lastNode = lastNode.GetChild(prefix[i])

		if lastNode == nil {
			return false
		}
	}
	return !exact || lastNode.Terminates()
}

func (t *Trie) Contains(prefix string) bool {
	return t.ContainsExactly(prefix, false)
}

func (t *Trie) GetRoot() *TrieNode {
	return t.root
}
