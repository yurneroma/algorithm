package main

import "fmt"

const Xmax = 1<<31 - 1

func main() {
	n := 0
	fmt.Scan(&n)
	f := make([][]int, n+10)
	s := make([]int, n+10)

	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
	}
	for i := 1; i <= n; i++ {
		s[i] += s[i-1]
	}

	for i := 0; i < n+10; i++ {
		f[i] = make([]int, n+10)
		for j := 0; j < n+10; j++ {
			f[i][j] = Xmax
			f[i][i] = 0
		}
	}

	for length := 2; length <= n; length++ { // 枚举长度
		for i := 1; i+length-1 <= n; i++ { // 枚举起点
			l := i
			r := i + length - 1
			for k := l; k < r; k++ {
				f[l][r] = min(f[l][r], f[l][k]+f[k+1][r]+s[r]-s[l-1])
			}
		}
	}
	fmt.Println(f[1][n])

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
