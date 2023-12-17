package complibgo

type edge struct {
	to int
	w  int
}

type GraphAdj struct {
	V   int
	Adj [][]edge
}

func NewGraphAdj(N int) *GraphAdj {
	g := &GraphAdj{
		N,
		make([][]edge, N),
	}
	return g
}

func (g *GraphAdj) AddEdge(u, v, w int) {
	if u >= g.V || v >= g.V {
		panic("AddEdge: node out of range")
	}
	if w < 0 {
		panic("AddEdge: negative edge weight")
	}
	g.Adj[u] = append(g.Adj[u], edge{v, w})
}
