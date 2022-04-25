package lib

type TrieNode struct {
	children   map[byte]*TrieNode
	terminates bool
	character  byte
}

func InitTrieNode() *TrieNode {
	node := new(TrieNode)
	node.children = make(map[byte]*TrieNode)
	return node
}

func NewTrieNode(character byte) *TrieNode {
	node := InitTrieNode()
	node.character = character
	return node
}

func (node *TrieNode) GetChar() byte {
	return node.character
}

func (node *TrieNode) AddWord(word string) {
	if len(word) == 0 {
		return
	}
	firstChar := word[0]
	child := node.GetChild(firstChar)
	if child == nil {
		child = NewTrieNode(firstChar)
		node.children[firstChar] = child
	}

	if len(word) > 1 {
		child.AddWord(word[1:])
	} else {
		child.SetTerminates(true)
	}
}

func (node *TrieNode) GetChild(c byte) *TrieNode {
	return node.children[c]
}

func (node *TrieNode) Terminates() bool {
	return node.terminates
}

func (node *TrieNode) SetTerminates(t bool) {
	node.terminates = t
}
