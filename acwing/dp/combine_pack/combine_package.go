package main

import "fmt"

// 组合背包
const N = 1010

var f [N]int

type Good struct {
	t int
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
		if s == -1 {
			goods = append(goods, Good{t: s, v: v, w: w})
		} else if s == 0 {
			goods = append(goods, Good{t: s, v: v, w: w})
		} else {
			// 多重背包转换成01背包
			for k := 1; k <= s; k *= 2 {
				goods = append(goods, Good{t: -1, v: v * k, w: w * k})
				s -= k
			}
			if s > 0 {
				goods = append(goods, Good{t: -1, v: v * s, w: w * s})
			}
		}
	}

	for _, good := range goods {
		if good.t == -1 {
			for j := m; j >= good.v; j-- {
				f[j] = max(f[j], f[j-good.v]+good.w)
			}
		} else {
			for j := good.v; j <= m; j++ {
				f[j] = max(f[j], f[j-good.v]+good.w)
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
