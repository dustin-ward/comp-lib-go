package complibgo

func Dijkstra() {
	println("TODO")
}

const INF = (1 << 62) - 1

type pii struct {
	fst int
	snd int
}

func Dijkstra_Sparse(G *GraphLst, src int) []int {
	D := make([]int, G.V)
	for i, _ := range D {
		D[i] = INF
	}

	pq := NewPQ[pii](func(a, b pii) bool {
		if a.fst == b.fst {
			return b.snd < a.snd
		}
		return a.fst < b.fst
	})

	pq.Push(pii{0, src})
	D[src] = 0

	for !pq.Empty() {
		cur := pq.Pop()
		u := cur.snd

		for _, e := range G.Adj[u] {
			v := e.To
			if D[v] > D[u]+e.W {
				D[v] = D[u] + e.W
				pq.Push(pii{D[v], v})
			}
		}
	}

	return D
}
