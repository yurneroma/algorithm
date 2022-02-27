package main

import "fmt"

const N = 1010

var f [N]int
var v [N]int
var w [N]int

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	// read v, w array, index from 1 to n
	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			f[j] = max(f[j], f[j-v[i]]+w[i])
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
