package evaluate_division

type Edge struct {
	v, w string
	weight float64
}

func NewEdge(v, w string, weight float64) *Edge {
	return &Edge{
		v, w, weight,
	}
}

type Graph struct {
	adj    map[string][]*Edge
}

func NewGraph() *Graph {
	return &Graph {
		adj: make(map[string][]*Edge),
	}
}

func (g *Graph) AddEdge(e *Edge) {
	v := e.v

	g.adj[v] = append(g.adj[v], e)
}

func (g *Graph) GetAdj() map[string][]*Edge {
	return g.adj
}

func (g *Graph) Adj(v string) []*Edge {
	return g.adj[v]
}

func (g *Graph) ValidateVertex(v string) bool {
	_, ok := g.adj[v]
	return ok
}
