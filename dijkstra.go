package complibgo

import "slices"

func Dijkstra() {
	println("TODO")
}

const INF = (1 << 62) - 1

type pii struct {
	fst int
	snd int
}

func Dijkstra_Sparse(G *GraphLst, src int) (D []int, P []int) {
	D = make([]int, G.V)
	P = make([]int, G.V)
	for i, _ := range D {
		D[i] = INF
		P[i] = -1
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
				P[v] = u
				pq.Push(pii{D[v], v})
			}
		}
	}

	return D, P
}

func GetPath(v int, P []int) (path []int) {
	path = make([]int, 1)
	path[0] = v

	for P[v] != -1 {
		v = P[v]
		path = append(path, v)
	}

	slices.Reverse(path)
	return
}
