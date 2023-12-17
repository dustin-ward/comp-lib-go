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
	g.Adj[u] = append(g.Adj[u], edge{v, w})
}
