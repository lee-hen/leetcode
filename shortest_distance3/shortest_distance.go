package shortest_distance

func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    wordsIndex := make(map[string]int)
    var distance = len(wordsDict)
    for i, str := range wordsDict {
        if word1 == str || word2 == str {
            if word1 == word2 {
                if idx, ok := wordsIndex[str]; ok {
                    distance = min(distance, abs(idx-i))
                }
            }
            wordsIndex[str] = i     
        }
        
        if word1 != word2 {
            idx1, ok1 := wordsIndex[word1] 
            idx2, ok2 := wordsIndex[word2]
            if ok1 && ok2 {
                distance = min(distance, abs(idx1-idx2))
            }
        }
    }
    return distance
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
