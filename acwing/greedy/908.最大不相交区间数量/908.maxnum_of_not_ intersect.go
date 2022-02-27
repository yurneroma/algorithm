package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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
		return areas[i][1] < areas[j][1]
	})
	//fmt.Println("areas", areas)
	res := 1
	end := areas[0][1]
	for i := 1; i < len(areas); i++ {
		cur := areas[i]
		if cur[0] > end {
			res++
			end = cur[1]
		}
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
