# Competitive Programming Library for Go

This is a collection of algorithms and data structures written in Go.

They might not all be fast enough for competitive programming platforms like Kattis or CodeForces, 
but they are good enough for Advent of Code (What I needed them for)

## Algorithms / Routines

### Binomial Coefficients
_binom.go_

`Binom(n, k int) int` O(k)

### Dijkstra's
_dijkstra.go_

`Dijkstra(g *GraphMtx, src int) []int` O(E\*log(V))

`Dijkstra_Sparse(g *GraphLst, src int) []int` O(V<sup>2</sup>)

## Data Structures

### GraphLst (Adjacency List)
_graph.go_

`NewGraphLst(N int) *GraphLst` 

`(g *GraphLst) AddEdge(u, v, w int)` O(1) - _O(N) worst case_

### GraphMtx (Adjacency Matrix)
_graph.go_

`NewGraphMtx(N int) *GraphMtx` 

`(g *GraphMtx) AddEdge(u, v, w int)` O(1)

### Priority Queue
_priority\_queue.go_

`NewPQ[T](beforeFunc func(a, b T)bool) *priority_queue[T]`

`(pq *priority_queue) Push(x T)` O(log(N))

`(pq *priority_queue) Pop(x T) T` O(log(N))

`(pq *priority_queue) Top() T` O(1)

`(pq *priority_queue) Len() int` O(1)

`(pq *priority_queue) Empty() bool` O(1)

### Union Find

`NewUnionFind(N int) *unionFind` O(N)

`(uf *unionFind) Find(x int) int` O(log(N))

`(uf *unionFind) Merge(x int) bool` O(log(N))
