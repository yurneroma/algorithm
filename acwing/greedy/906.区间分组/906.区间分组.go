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

	heap := &Heap{h: make([]int, 1)}
	for i := 0; i < len(areas); i++ {
		l := areas[i][0]
		r := areas[i][1]
		if heap.size == 0 || heap.top() >= l {
			heap.push(r)
		} else {
			heap.pop()
			heap.push(r)
		}
	}
	fmt.Println(heap.size)
}

type Heap struct {
	h    []int
	size int
}

func (heap *Heap) top() int {
	if heap.size == 0 {
		return -1
	}
	return heap.h[1]
}

func (heap *Heap) push(x int) {
	heap.h = append(heap.h, x)
	heap.size++
	heap.up(heap.size)
}

func (heap *Heap) pop() int {
	x := heap.h[1]
	heap.h[1] = heap.h[heap.size]
	heap.h = heap.h[:heap.size]
	heap.size--
	heap.down(1)
	return x
}

func (heap *Heap) up(u int) {
	for u/2 != 0 && heap.h[u] < heap.h[u/2] {
		heap.h[u], heap.h[u/2] = heap.h[u/2], heap.h[u]
		u /= 2
	}
}

func (heap *Heap) down(u int) {
	t := u
	if u*2 <= heap.size && heap.h[t] > heap.h[u*2] {
		t = u * 2
	}

	if u*2+1 <= heap.size && heap.h[t] > heap.h[u*2+1] {
		t = u*2 + 1
	}

	if t != u {
		heap.h[t], heap.h[u] = heap.h[u], heap.h[t]
		heap.down(t)
	}
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
