package design_search_autocomplete_system

type TrieNode struct {
	children   map[byte]*TrieNode
	terminates bool
	time int
}

func NewTrieNode() *TrieNode {
	node := new(TrieNode)
	node.children = make(map[byte]*TrieNode)
	return node
}

func (node *TrieNode) Insert(word string, time int) {
	currNode := node
	for i := range word {
		c := word[i]
		child := currNode.children[c]
		if child == nil {
			child = NewTrieNode()
			currNode.children[c] = child
		}
		currNode = child
	}
	currNode.SetTerminates(true)
	currNode.time += time
}

func (node *TrieNode) Terminates() bool {
	return node.terminates
}

func (node *TrieNode) SetTerminates(t bool) {
	node.terminates = t
}

func (node *TrieNode) allPaths(path string) []HotSentence {
	paths := make([]HotSentence, 0)
	if node.time > 0 {
		paths = append(paths, HotSentence{path, node.time})
	}
	for c := range node.children {
		paths = append(paths, node.children[c].allPaths(path+string(c))...)
	}
	return paths
}
