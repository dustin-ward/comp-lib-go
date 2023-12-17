package complibgo

import "testing"

func TestDijkstraSparse1(t *testing.T) {
	N := 9
	g := NewGraphAdj(N)

	g.AddEdge(0, 1, 4)
	g.AddEdge(1, 0, 4)
	g.AddEdge(0, 7, 8)
	g.AddEdge(7, 0, 8)
	g.AddEdge(1, 2, 8)
	g.AddEdge(2, 1, 8)
	g.AddEdge(1, 7, 11)
	g.AddEdge(7, 1, 11)
	g.AddEdge(2, 3, 7)
	g.AddEdge(3, 2, 7)
	g.AddEdge(2, 8, 2)
	g.AddEdge(8, 2, 2)
	g.AddEdge(2, 5, 4)
	g.AddEdge(5, 2, 4)
	g.AddEdge(3, 4, 9)
	g.AddEdge(4, 3, 9)
	g.AddEdge(3, 5, 14)
	g.AddEdge(5, 3, 14)
	g.AddEdge(4, 5, 10)
	g.AddEdge(5, 4, 10)
	g.AddEdge(5, 6, 2)
	g.AddEdge(6, 5, 2)
	g.AddEdge(6, 7, 1)
	g.AddEdge(7, 6, 1)
	g.AddEdge(6, 8, 6)
	g.AddEdge(8, 6, 6)
	g.AddEdge(7, 8, 7)
	g.AddEdge(8, 7, 7)

	D := Dijkstra_Sparse(g, 0)
	t.Log(D)

	ans := []int{0, 4, 12, 19, 21, 11, 9, 8, 14}
	for i, v := range D {
		if v != ans[i] {
			t.Fatalf("exected D[%d] == %d, but got %d", i, ans[i], D[i])
		}
	}
}
