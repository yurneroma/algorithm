package main

import "fmt"

const N = 2010

var f [N]int

type Good struct {
	v int
	w int
}

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)

	goods := make([]Good, 0)
	for i := 1; i <= n; i++ {
		v, w, s := 0, 0, 0
		fmt.Scan(&v, &w, &s)
		for k := 1; k <= s; k *= 2 {
			goods = append(goods, Good{v * k, w * k})
			s -= k
		}
		if s > 0 {
			goods = append(goods, Good{v * s, w * s})
		}

	}

	for i := 0; i < len(goods); i++ {
		for j := m; j >= 1; j-- {
			t := goods[i]
			if j >= t.v {
				f[j] = max(f[j], f[j-t.v]+t.w)
			}
		}
	}
	fmt.Println(f[m])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
