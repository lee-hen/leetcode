package shortest_distance2

import "math"

type WordDistance struct {
    wordsIndices map[string][]int
}


func Constructor(wordsDict []string) WordDistance {
    wordsIndices := make(map[string][]int)
    
    for i, word := range wordsDict {
        wordsIndices[word] = append(wordsIndices[word], i)
    }
    
    
    return WordDistance{
        wordsIndices: wordsIndices,
    }
}


// 1 5 9 i
// 6 7   j

// 1


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    word1Indices, word2Indices := this.wordsIndices[word1], this.wordsIndices[word2]
    if len(word1Indices) < len(word2Indices) {
        word1Indices, word2Indices = word2Indices, word1Indices
    }
    
    
    var i, j int
    var shortDistance = math.MaxInt32
    for i < len(word1Indices) {
        if j >= len(word2Indices) {
            break
        }
        
        shortDistance = min(shortDistance, abs(word2Indices[j]-word1Indices[i]))
        if word1Indices[i] < word2Indices[j] {
            i++
        } else if word2Indices[j] < word1Indices[i] {
            j++
        }
    }
    
    return shortDistance
}


func min(arg int, rest ...int) int {
    curr := arg
    for _, num := range rest {
        if curr > num {
            curr = num
        }
    }
    
    return curr
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */
 