package complibgo

type priority_queue[T any] struct {
	heap   []T
	before func(a, b T) bool
}

type beforeFunc[T any] func(i, j T) bool

func NewPQ[T any](fn beforeFunc[T]) *priority_queue[T] {
	pq := &priority_queue[T]{
		make([]T, 0, 16),
		fn,
	}
	return pq
}

func (pq *priority_queue[T]) Push(x T) {
	if cap(pq.heap) == len(pq.heap) {
		pq.heap = append(make([]T, 0, cap(pq.heap)*2), pq.heap...)
	}

	pq.heap = append(pq.heap, x)
	pq.shiftUp(len(pq.heap) - 1)
}

func (pq *priority_queue[T]) Pop() T {
	if pq.Empty() {
		panic("Top: priority queue is empty")
	}
	top := pq.heap[0]

	pq.heap[0] = pq.heap[len(pq.heap)-1]
	pq.shiftDown(0)
	pq.heap = pq.heap[:len(pq.heap)-1]

	return top
}

func (pq *priority_queue[T]) Top() T {
	if pq.Empty() {
		panic("Top: priority queue is empty")
	}
	return pq.heap[0]
}

func (pq *priority_queue[_]) Empty() bool {
	return len(pq.heap) == 0
}

func (pq *priority_queue[_]) Len() int {
	return len(pq.heap)
}

func parent(i int) int {
	return (i - 1) / 2
}

func lChild(i int) int {
	return (2 * i) + 1
}

func rChild(i int) int {
	return (2 * i) + 2
}

func (pq *priority_queue[_]) shiftUp(i int) {
	for i > 0 && pq.before(pq.heap[i], pq.heap[parent(i)]) {
		pq.heap[i], pq.heap[parent(i)] = pq.heap[parent(i)], pq.heap[i]
		i = parent(i)
	}
}

func (pq *priority_queue[T]) shiftDown(i int) {
	i2 := i

	l := lChild(i)
	if l < len(pq.heap) && pq.before(pq.heap[l], pq.heap[i2]) {
		i2 = l
	}

	r := rChild(i)
	if r < len(pq.heap) && pq.before(pq.heap[r], pq.heap[i2]) {
		i2 = r
	}

	if i != i2 {
		pq.heap[i], pq.heap[i2] = pq.heap[i2], pq.heap[i]
		pq.shiftDown(i2)
	}
}
