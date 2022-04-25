package minimum_knight_moves

import (
	"fmt"
	"math"
)

func MinKnightMovesDFS(x int, y int) int {
	memo := make(map[string]int)
	return dfs(Abs(x), Abs(y), memo)
}

func dfs(x, y int, memo map[string]int) int {
	key := generateKey(x, y)
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	if x + y == 0 {
		return 0
	} else if x + y == 2 {
		return 2
	} else {
		memo[key] = Min(dfs(Abs(x-2), Abs(y-1), memo), dfs(Abs(x-1), Abs(y-2), memo)) + 1
		return memo[key]
	}
}

func minKnightMovesDFS(x int, y int) int {
	moves, _ := helper(0, 0, []int{Abs(x), Abs(y)}, map[string]bool{}, map[string]int{})
	return moves
}

func helper(i, j int, target []int, visited map[string]bool, cache map[string]int) (int, bool) {
	key := generateKey(i, j)

	if value, ok := cache[key]; ok {
		return cache[key], value != math.MaxInt32
	}

	if outOfBounds(i, j) {
		return math.MaxInt32, false
	}

	if i == target[0] && j == target[1] {
		return 0, true
	}

	visited[key] = true

	var moves = math.MaxInt32
	nextPositions := positions(i, j, target)
	for _, next := range nextPositions {
		nextKey := generateKey(next[0], next[1])
		if visited[nextKey] {
			continue
		}

		steps, found := helper(next[0], next[1], target, visited, cache)
		if found {
			moves = Min(moves, steps+1)
		}
		visited[nextKey] = false
	}

	cache[key] = moves
	return cache[key], moves != math.MaxInt32
}


func positions(x, y int, target []int) [][]int {
	nextPositions := make([][]int, 0)

	if x < target[0] && y < target[1]+2  {
		nextPositions = append(nextPositions, []int{x + 2, y + 1})
		nextPositions = append(nextPositions, []int{x + 1, y + 2})
	}

	if x > target[0]-2 && y < target[1] {
		nextPositions = append(nextPositions, []int{x - 2, y + 1})
		nextPositions = append(nextPositions, []int{x - 1, y + 2})
	}

	if y > target[1] && x < target[0] {
		nextPositions = append(nextPositions, []int{x + 2, y - 1})
		nextPositions = append(nextPositions, []int{x + 1, y - 2})
	}

	return nextPositions
}

func generateKey(i, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}

func outOfBounds(i, j int) bool {
	if Abs(i) + Abs(j) > 300 || Abs(i) + Abs(j) < 0 {
		return true
	}
	return false
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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
