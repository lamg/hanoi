package hanoi

import "fmt"

type towers [][]uint

func move(t towers, s, d int) (ok bool) {
	ok = len(t[s]) != 0
	var d0 uint
	if ok {
		d0 = t[s][len(t[s])-1]
		if len(t[d]) != 0 {
			d2 := t[d][len(t[d])-1]
			ok = d0 < d2
		}
	}
	if ok {
		t[s] = t[s][:len(t[s])-1]
		t[d] = append(t[d], d0)
	}
	return
}

type play struct {
	src int
	dst int
}

func moves(t towers, ps []play) (ok bool, n int) {
	n, ok = 0, true
	for ok && n != len(ps) {
		printTowers(t)
		ok, n = move(t, ps[n].src, ps[n].dst), n+1
	}
	if !ok {
		n = n - 1
	}
	return
}

func solved(t towers, n int) (r bool) {
	r = len(t[0]) == 0 && len(t[1]) == 0 && len(t[2]) == n
	i := 0
	for r && i != n {
		r, i = t[2][i] == uint(n-i), i+1
	}
	return
}

func printTowers(t towers) {
	for i, j := range t {
		fmt.Printf("%d:{", i)
		for _, k := range j {
			fmt.Printf("%d,", k)
		}
		fmt.Printf("}\n")
	}
}
