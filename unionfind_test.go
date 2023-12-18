package complibgo

import "testing"

type Op struct {
	op   string
	u, v int
	val  bool
}

var uftests = []struct {
	N int
	O []Op
}{
	{
		10,
		[]Op{
			Op{"?", 1, 3, false},
			Op{"=", 1, 8, true},
			Op{"=", 3, 8, true},
			Op{"?", 1, 3, true},
		},
	},
	{
		4,
		[]Op{
			Op{"?", 0, 0, true},
			Op{"=", 0, 1, true},
			Op{"=", 1, 2, true},
			Op{"=", 0, 2, false},
			Op{"?", 0, 3, false},
		},
	},
}

func TestUnionFind(t *testing.T) {
	for _, test := range uftests {
		uf := NewUnionFind(test.N)
		for _, op := range test.O {
			if op.op == "=" {
				b := uf.Merge(op.u, op.v)
				if b != op.val {
					t.Fatalf("unexpected value %v, for Merge(%d, %d), expected %v", b, op.u, op.v, op.val)
				}
			} else {
				b := uf.Find(op.u) == uf.Find(op.v)
				if b != op.val {
					t.Fatalf("unexpected value %v, for Find(%d) == Find(%d), expected %v", b, op.u, op.v, op.val)
				}
			}
		}
	}
}
