package android_unlock_patterns

func NumberOfPatterns(m int, n int) int {
	// Skip array represents number to skip between two pairs
	skip := make([][]int, 10)
	for i := range skip {
		skip[i] = make([]int, 10)
	}

	skip[1][3], skip[3][1] = 2, 2
	skip[1][7], skip[7][1] = 4, 4
	skip[3][9], skip[9][3] = 6, 6
	skip[7][9], skip[9][7] = 8, 8
	skip[1][9], skip[9][1], skip[2][8], skip[8][2], skip[3][7], skip[7][3], skip[4][6], skip[6][4] = 5, 5, 5, 5, 5, 5, 5, 5

	var dfs func(visited []bool, cur, remain int) int
	dfs = func(visited []bool, cur, remain int) int {
		if remain < 0 {
			return 0
		}

		if remain == 0 {
			return 1
		}
		visited[cur] = true
		res := 0
		for i := 1; i <= 9; i++ {
			// If vis[i] is not visited and (two numbers are adjacent or skip number is already visited)
			if !visited[i] && (skip[cur][i] == 0 || (visited[skip[cur][i]])) {
				res += dfs(visited, i, remain - 1)
			}
		}
		visited[cur] = false
		return res
	}

	visited := make([]bool, 10)
	res := 0
	// DFS search each length from m to n
	for i := m; i <= n; i++ {
		res += dfs(visited, 1, i - 1) * 4    // 1, 3, 7, 9 are symmetric
		res += dfs(visited, 2, i - 1) * 4    // 2, 4, 6, 8 are symmetric
		res += dfs(visited, 5, i - 1)        // 5
	}
	return res
}

func numberOfPatterns(m int, n int) int {
	var patterns int

	for depth := m; depth <= n; depth++ {
		seen := make(map[int]bool)
		patterns += dfs(1, depth, seen) * 4
		patterns += dfs(2, depth, seen) * 4
		patterns += dfs(5, depth, seen)
	}

	return patterns
}

func dfs(node, depth int, seen map[int]bool) int {
	if depth == 1 {
		return 1
	}

	var patterns int
	seen[node] = true
	neighbours := getNeighbours(node, seen)
	for _, neighbour := range neighbours {
		if seen[neighbour] {
			continue
		}
		patterns += dfs(neighbour, depth-1, seen)
	}
	seen[node] = false
	return patterns
}

func getNeighbours(node int, seen map[int]bool) [] int {
	neighbours := make([]int, 0)
	for i := 1; i <= 9; i++ {
		if node == i {
			continue
		}

		if node == 1 {
			if i == 3 && !seen[2] {
				continue
			}

			if i == 9 && !seen[5] {
				continue
			}

			if i == 7 && !seen[4] {
				continue
			}
		}

		if node == 2 {
			if i == 8 && !seen[5] {
				continue
			}
		}

		if node == 3 {
			if i == 1 && !seen[2] {
				continue
			}

			if i == 7 && !seen[5] {
				continue
			}

			if i == 9 && !seen[6] {
				continue
			}
		}


		if node == 4 {
			if i == 6 && !seen[5] {
				continue
			}
		}

		if node == 6 {
			if i == 4 && !seen[5] {
				continue
			}
		}

		if node == 7 {
			if i == 1 && !seen[4] {
				continue
			}

			if i == 3 && !seen[5] {
				continue
			}

			if i == 9 && !seen[8] {
				continue
			}
		}

		if node == 8 {
			if i == 2 && !seen[5] {
				continue
			}
		}

		if node == 9 {
			if i == 1 && !seen[5] {
				continue
			}

			if i == 7 && !seen[8] {
				continue
			}

			if i == 3 && !seen[6] {
				continue
			}
		}

		neighbours = append(neighbours, i)
	}

	return neighbours
}

// 2-1-3, 2-3-1 are valid moves
func buildGraph() map[int][]int {
	graph := make(map[int][]int)
	graph[1] = append(graph[1], []int{2, 6, 5, 8, 4}...)
	graph[2] = append(graph[2], []int{1, 4, 7, 5, 9, 6, 3}...)
	graph[3] = append(graph[3], []int{2, 4, 5, 8, 6}...)
	graph[4] = append(graph[4], []int{1, 2, 3, 5, 9, 8, 7}...)
	graph[5] = append(graph[5], []int{1, 2, 3, 6, 9, 8, 7, 4}...)
	graph[6] = append(graph[6], []int{1, 2, 3, 5, 7, 8, 9}...)
	graph[7] = append(graph[7], []int{8, 6, 5, 2, 4}...)
	graph[8] = append(graph[8], []int{1, 4, 7, 5, 3, 6, 9}...)
	graph[9] = append(graph[9], []int{8, 4, 5, 2, 6}...)
	return graph
}