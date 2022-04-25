package minimum_knight_moves

import (
	"container/list"
)

// MinKnightMovesBFS
// Bidirectional BFS
func MinKnightMovesBFS(x int, y int) int {
	// the offsets in the eight directions
	offsets := [][]int{
		{1, 2}, {2, 1}, {2, -1}, {1, -2},
		{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2},
	}

	// data structures needed to move from the origin point
	originQueue := list.New()
	originQueue.PushBack([]int{0, 0, 0})
	originDistance := make(map[string]int)
	originDistance[generateKey(0, 0)] = 0

	// data structures needed to move from the target point
	targetQueue := list.New()
	targetQueue.PushBack([]int{x, y, 0})
	targetDistance := make(map[string]int)
	targetDistance[generateKey(x, y)] = 0

	for {
		// check if we reach the circle of target
		first := originQueue.Front()
		originQueue.Remove(first)
		origin := first.Value.([]int)

		if value, ok := targetDistance[generateKey(origin[0], origin[1])]; ok {
			return origin[2] + value
		}

		// check if we reach the circle of origin
		first = targetQueue.Front()
		targetQueue.Remove(first)
		target := first.Value.([]int)

		if value, ok := originDistance[generateKey(target[0], target[1])]; ok {
			return target[2] + value
		}

		for _, offset := range offsets {
			// expand the circle of origin
			nextOrigin := []int{origin[0] + offset[0], origin[1] + offset[1]}

			if _, ok := originDistance[generateKey(nextOrigin[0], nextOrigin[1])]; !ok {
				originQueue.PushBack([]int{nextOrigin[0], nextOrigin[1], origin[2] + 1})
				originDistance[generateKey(nextOrigin[0], nextOrigin[1])] = origin[2] + 1
			}

			// expand the circle of target
			nextTarget := []int{target[0] + offset[0], target[1] + offset[1]}
			if _, ok := targetDistance[generateKey(nextTarget[0], nextTarget[1])]; !ok {
				targetQueue.PushBack([]int{nextTarget[0], nextTarget[1], target[2] + 1})
				targetDistance[generateKey(nextTarget[0], nextTarget[1])] = origin[2] + 1
			}
		}
	}
}

// unidirectional BFS
func minKnightMovesBFS(x int, y int) int {
	offsets := [][]int{
		{1, 2}, {2, 1}, {2, -1}, {1, -2},
		{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2},
	}

	bfs := func(x, y int) int {
		visited := make(map[string]bool)

		queue := list.New()
		queue.PushBack([]int{0, 0})
		steps := 0

		for queue.Len() > 0 {
			currLevel := queue.Len()
			for i := 0; i < currLevel; i++ {
				current := queue.Front()
				queue.Remove(current)

				curr := current.Value.([]int)
				if curr[0] == x && curr[1] == y {
					return steps
				}

				for _, offset := range offsets {
					next := []int{
						curr[0] + offset[0],
						curr[1] + offset[1],
					}
					nextKey := generateKey(next[0]+302, next[1]+302)
					if !visited[nextKey] {
						visited[nextKey] = true
						queue.PushBack(next)
					}
				}
			}
			steps++
		}
		return steps
	}

	return bfs(x, y)
}

