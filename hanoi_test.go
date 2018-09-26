package hanoi

import (
	"testing"
)

func TestMov(t *testing.T) {
	tr := [][]uint{
		{4, 3, 2, 1},
		{},
		{},
	}
	ok := move(tr, 0, 2)
	if !ok {
		t.Fatalf("Not ok")
	}
	testEqual(t, tr[0], []uint{4, 3, 2})
	testEqual(t, tr[1], []uint{})
	testEqual(t, tr[2], []uint{1})
}

func testEqual(t *testing.T, a, b []uint) {
	r := equal(a, b)
	if !r {
		t.Fatalf("Not equal")
	}
}

func equal(a, b []uint) (r bool) {
	r = len(a) == len(b)
	n := 0
	for r && n != len(a) {
		r, n = a[n] == b[n], n+1
	}
	return
}

func TestSolve(t *testing.T) {
	sols := []solution{
		{
			n: 2,
			ps: []play{
				{src: 0, dst: 1}, //B
				{src: 0, dst: 2}, //A
				{src: 1, dst: 2}, //C
			},
			ts: [][]uint{
				{2, 1},
				{},
				{},
			},
		},
		{
			n: 3,
			ps: []play{
				{src: 0, dst: 2}, //A
				{src: 0, dst: 1}, //B
				{src: 2, dst: 1}, //D
				{src: 0, dst: 2}, //A
				{src: 1, dst: 0}, //E
				{src: 1, dst: 2}, //C
				{src: 0, dst: 2}, //A
			},
			ts: [][]uint{
				{3, 2, 1},
				{},
				{},
			},
		},
	}
	for i, j := range sols {
		ok, n := moves(j.ts, j.ps)
		if !ok {
			t.Fatalf("An invalid play in case %d at %d", i, n)
		}
		ok = solved(j.ts, j.n)
	}
}

type solution struct {
	n  int
	ps []play
	ts towers
}
