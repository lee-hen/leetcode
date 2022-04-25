package evaluate_division

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := buildGraph(equations, values)
	result := make([]float64, len(queries))

	for i := range queries {
		v := queries[i][0]
		w := queries[i][1]
		value := 1.0
		visited := make(map[string]bool)
		validated := g.ValidateVertex(v) && g.ValidateVertex(w)

		if validated && dfs(g, visited, &value, v, w) {
			result[i] = value
		} else {
			result[i] = -1.0
		}
	}

	return result
}

func dfs(g *Graph, visited map[string]bool, value *float64,  v, w string) bool {
	if v == w {
		return true
	}

	visited[v] = true
	edges := g.Adj(v)

	for _, edge := range edges {
		if visited[edge.w] {
			continue
		}

		*value *= edge.weight

		if dfs(g, visited, value, edge.w, w) {
			return true
		} else {
			*value /= edge.weight
		}
	}

	return false
}

func buildGraph(equations [][]string, values []float64) *Graph {
	g := NewGraph()
	for i := range equations {
		v := equations[i][0]
		w := equations[i][1]
		e1 := NewEdge(v, w, values[i])
		e2 := NewEdge(w, v, 1.0 / values[i])

		g.AddEdge(e1)
		g.AddEdge(e2)
	}

	return g
}
