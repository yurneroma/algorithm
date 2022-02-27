package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	l int
	r int
}

func main() {
	st, end := 0, 0
	fmt.Scan(&st, &end)

	n := 0
	fmt.Scan(&n)
	areas := make([][]int, 0)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		l, r := readline(reader)
		areas = append(areas, []int{l, r})
	}

	// sort by right point of area
	sort.Slice(areas, func(i, j int) bool {
		return areas[i][0] < areas[j][0]
	})

	res := 0
	success := false
	for i := 0; i < n; i++ {
		r := -1 << 31
		j := i
		// 从剩下的里面找到可以覆盖st的右端点的最大值
		for j < n && areas[j][0] <= st {
			r = max(r, areas[j][1])
			j++
		}
		if r < st {
			res = -1
			break
		}
		res++
		if r >= end {
			success = true
			break
		}

		i = j - 1
		st = r
	}

	if !success {
		res = -1
	}
	fmt.Println(res)

}

func readline(reader *bufio.Reader) (int, int) {
	line, _ := reader.ReadString('\n')
	params := strings.Fields(strings.Trim(line, "\n"))
	l, _ := strconv.Atoi(params[0])
	r, _ := strconv.Atoi(params[1])
	return l, r
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
