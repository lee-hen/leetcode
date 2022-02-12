package most_stones_removed_with_same_row_or_column

func removeStones1(stones [][]int) int {
	visited := make(map[Point]bool)

	var dfs func(p Point, visited map[Point]bool, stones [][]int)
	dfs = func(p Point, visited map[Point]bool, stones [][]int) {
		visited[p] = true

		for _, s := range stones {
			q := Point{s[0], s[1]}

			if !visited[q] {
				// stone with same row or column. group them into island
				if q.row == p.row || p.col == q.col {
					dfs(q, visited, stones)
				}
			}
		}
	}

	numOfIslands := 0
	for _, s := range stones {
		p := Point{s[0], s[1]}

		if !visited[p] {
			dfs(p, visited, stones)
			numOfIslands++
		}
	}

	return len(stones) - numOfIslands
}

type Point struct {
	row, col int
}

type graph map[Point][]Point

func removeStones(stones [][]int) int {
	g := buildGraph(stones)
	visited := make(map[Point]bool)
	removed := make(map[Point]bool)

	for i := range stones {
		p := Point{stones[i][0], stones[i][1]}
		dfs(g, visited, removed, p)
	}

	return len(removed)
}

func dfs(g graph, visited, removed map[Point]bool, p Point) {
	visited[p] = true
	adj := g[p]

	for _, point := range adj {
		if visited[point] {
			continue
		}
		removed[point] = true
		dfs(g, visited, removed, point)
	}
}

func buildGraph(stones [][]int) graph {
	groupByRow := make(map[int]map[Point]struct{})
	groupByCol := make(map[int]map[Point]struct{})

	for i := range stones {
		p := Point{stones[i][0], stones[i][1]}

		if _, ok := groupByRow[p.row]; !ok {
			groupByRow[p.row] = make(map[Point]struct{})
		}

		if _, ok := groupByCol[p.col]; !ok {
			groupByCol[p.col] = make(map[Point]struct{})
		}

		groupByRow[p.row][p] = struct{}{}
		groupByCol[p.col][p] = struct{}{}
	}

	g := make(graph)

	for i := range stones {
		p := Point{stones[i][0], stones[i][1]}

		for point := range groupByRow[p.row] {
			if p == point {
				continue
			}
			g[p] = append(g[p], point)
		}

		for point := range groupByCol[p.col] {
			if p == point {
				continue
			}
			g[p] = append(g[p], point)
		}
	}

	return g
}
