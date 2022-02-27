package main

import "fmt"

const N = 1010

var f [N][N]int

var s [N]int // s[i] 第i种个数
var v [N]int // v[i] 体积
var w [N]int // w[i]value

func main() {
	// n种， m总体积
	n, m := 0, 0
	fmt.Scan(&n, &m)

	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i], &s[i])
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k <= s[i] && k*v[i] <= j; k++ {
				f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}
	fmt.Println(f[n][m])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
