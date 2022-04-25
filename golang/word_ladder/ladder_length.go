package word_ladder

import (
	"container/list"
	"math"
)

type PathNode struct {
	word string
	previousNode *PathNode
}

func (p *PathNode) GetWord() string {
	return p.word
}

func NewPathNode(word string, previous *PathNode) *PathNode {
	p := new(PathNode)
	p.word = word
	p.previousNode = previous

	return p
}

func (p *PathNode) Collapse(startsWithRoot bool) *list.List{
	path := list.New()
	node := p
	for node != nil {
		if startsWithRoot {
			path.PushBack(node.word)
		} else {
			path.PushFront(node.word)
		}
		node = node.previousNode
	}
	return path
}

type BFSData struct {
	toVisit *list.List
	visited map[string]*PathNode
}

func NewBFSData(root string) *BFSData{
	bfsData := new(BFSData)
	sourcePath := NewPathNode(root, nil)
	bfsData.toVisit = list.New()
	bfsData.toVisit.PushBack(sourcePath)
	bfsData.visited = make(map[string]*PathNode)
	bfsData.visited[root] = sourcePath
	return bfsData
}

func (bfsData *BFSData) IsFinished() bool {
	return bfsData.toVisit.Len() == 0
}

func LadderLength(beginWord string, endWord string, wordList []string) int {
	for _, word := range wordList {
		if endWord == word {
			return len(TransformBFS(beginWord, endWord, wordList))
		}
	}

	return 0
}

func TransformBFS(start, stop string, words []string) []string {
	wildcardToWordList := getWildcardToWordList(words)
	sourceData := NewBFSData(start)
	destData := NewBFSData(stop)

	for !sourceData.IsFinished() && !destData.IsFinished() {
		bfs := func(wildcardToWordList map[string][]string, primary, secondary *BFSData) string {
			count := primary.toVisit.Len()
			for i := 0; i < count; i++ {
				front := primary.toVisit.Front()
				primary.toVisit.Remove(front)
				pathNode := front.Value.(*PathNode)
				word := pathNode.word

				if _, ok := secondary.visited[word]; ok {
					return word
				}

				words := getValidLinkedWords(word, wildcardToWordList)
				for _, w := range words {
					if _, ok := primary.visited[w]; ok {
						continue
					}

					next := NewPathNode(w, pathNode)
					primary.visited[w] = next
					primary.toVisit.PushBack(next)
				}
			}

			return ""
		}

		/* Search out from source. */
		collision := bfs(wildcardToWordList, sourceData, destData)
		if collision != "" {
			return mergePaths(sourceData, destData, collision)
		}

		/* Search out from destination. */
		collision = bfs(wildcardToWordList, destData, sourceData)
		if collision != "" {
			return mergePaths(sourceData, destData, collision)
		}
	}

	return nil
}

func getWildcardToWordList(words []string) map[string][]string {
	wildcardToWords := make(map[string][]string)

	for _, word := range words {
		linked := getWildcardRoots(word)
		for _, linkedWord := range linked {
			wildcardToWords[linkedWord] = append(wildcardToWords[linkedWord], word)
		}
	}
	return wildcardToWords
}

func getWildcardRoots(word string) []string {
	words := make([]string, 0)
	for i := 0; i < len(word); i++ {
		w := word[:i] + "_" + word[i+1:]
		words = append(words, w)
	}
	return words
}

func getValidLinkedWords(word string, wildcardToWords map[string][]string) []string {
	wildcards := getWildcardRoots(word)
	linkedWords := make([]string, 0)

	for _, wildcard := range wildcards {
		words := wildcardToWords[wildcard]
		for _, linkedWord := range words {
			if linkedWord != word {
				linkedWords = append(linkedWords, linkedWord)
			}
		}
	}
	return linkedWords
}

func mergePaths(bfs1, bfs2 *BFSData, connection string) []string {
	end1 := bfs1.visited[connection] // end1 -> source
	end2 := bfs2.visited[connection] // end2 -> dest
	pathOne := end1.Collapse(false) // forward: source -> connection
	pathTwo := end2.Collapse(true) // reverse: connection -> dest

	pathTwo.Remove(pathTwo.Front())
	pathOne.PushBackList(pathTwo)

	path := make([]string, 0)
	curr := pathOne.Front()
	for curr != nil {
		path = append(path,  curr.Value.(string))
		curr = curr.Next()
	}

	return path
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	return transFormDFS(beginWord, endWord, wordList)
}

func transFormDFS(start, stop string, words []string) int {
	if len(start) != len(stop) {
		return 0
	}

	graph := make(map[string][]string)
	seen := make(map[string]bool)
	words = append(words, start)
	createTransFormGraph(stop, words, graph, seen)

	minPath := math.MaxInt32

	seen = make(map[string]bool)
	var dfs func(start string, path int, graph map[string][]string)

	dfs = func(start string, path int, graph map[string][]string)  {
		seen[start] = true

		if stop == start {
			minPath = Min(minPath, path)
		}

		for _, next := range graph[start] {
			if seen[next] {
				continue
			}
			dfs(next, path+1, graph)
		}

		seen[start] = false
	}

	dfs(start, 1, graph)
	if minPath == math.MaxInt32 {
		return 0
	}

	return minPath
}

func createTransFormGraph(curr string, words []string, graph map[string][]string, seen map[string]bool) {
	if seen[curr] == true {
		return
	}
	seen[curr] = true
	for _, word := range words {
		if couldTransform(curr, word) {
			graph[curr] = append(graph[curr], word)
			createTransFormGraph(word, words, graph, seen)
		}
	}
}

func couldTransform(curr, other string) bool {
	if len(curr) != len(other) {
		if Abs(len(curr)-len(other)) == 1 {
			//if len(curr) < len(other) {
			//	return couldTransform(other, curr)
			//}
			//return strings.Contains(curr, other)
		}
		return false
	}

	var cnt int
	for i := 0; i < len(curr); i++ {
		if curr[i] != other[i] {
			cnt++
		}
	}

	return cnt == 1
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
