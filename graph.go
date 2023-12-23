package complibgo

type Edge struct {
	To int
	W  int
}

type GraphLst struct {
	V   int
	Adj [][]edge
}

func NewGraphLst(N int) *GraphLst {
	g := &GraphLst{
		N,
		make([][]edge, N),
	}
	return g
}

func (g *GraphLst) AddEdge(u, v, w int) {
	if u >= g.V || v >= g.V {
		panic("AddEdge: node out of range")
	}
	if w < 0 {
		panic("AddEdge: negative edge weight")
	}
	g.Adj[u] = append(g.Adj[u], edge{v, w})
}
