package complibgo

import (
	"math/rand"
	"sort"
	"testing"
)

func TestPQ(t *testing.T) {
	const TEST_SIZE = 100000

	data := make([]int, TEST_SIZE)
	for i, _ := range data {
		data[i] = rand.Int()
	}

	pq := NewPQ[int](func(i, j int) bool { return i < j })
	for _, v := range data {
		pq.Push(v)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	for i, v := range data {
		if pq.Empty() {
			t.Fatalf("pq empty when %v elements should be remaining", len(data)-i)
		}

		top := pq.Top()
		t.Logf("data: %v == top: %v?", v, top)
		if top != v {
			t.Fatalf("expected %v, got %v", v, top)
		}

		top2 := pq.Pop()
		if top2 != top {
			t.Fatalf("Pop() and Top() returned different values: %v & %v", top2, top)
		}
		if pq.Len() != (len(data)-i)-1 {
			t.Fatalf("expected %d elements remaining, got %d", (len(data)-i)-1, pq.Len())
		}
	}

	if !pq.Empty() {
		t.Fatalf("expected pq to be empty, actually length:%d", pq.Len())
	}
}
