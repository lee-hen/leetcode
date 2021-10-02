package shortest_distance

func ShortestDistance(wordsDict []string, word1 string, word2 string) int {
    index1, index2 := -1, -1
    var distance = len(wordsDict)
    for i, str := range wordsDict {
        if word1 == str {
            index1 = i
        } else if word2 == str {
            index2 = i
        }
        
        if index1 != -1 && index2 != -1 {
            distance = min(distance, abs(index1-index2))
        }
    }
    return distance
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
