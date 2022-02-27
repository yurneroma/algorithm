package main

import "fmt"

// solution1
const N = 1010

var f [N][N]int // f[i][j] 前i个物品， 体积小于等于j的最大值
var v [N]int    // v[i] 体积
var w [N]int    // w[i] value

func solution1() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k*v[i] <= j; k++ {
				f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}
	fmt.Println(f[n][m])
}

//一维数组优化
var ff [N]int //ff[j]:体积不大于j的情况下的最大值
func solution2() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
			ff[j] = max(ff[j], ff[j-v[i]]+w[i])
		}
	}
	fmt.Println(ff[m])
}

func main() {
	solution2()
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
