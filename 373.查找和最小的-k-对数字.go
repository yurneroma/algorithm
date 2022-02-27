/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的 K 对数字
 */

// @lc code=start
func kSmallestPairs(a []int, b []int, k int) [][]int {
	res := make([][]int, 0)
	heap := &Heap{h: make([]Pair, 1)}
	n := len(a)
	m := len(b)
	if n == 0 || m == 0 {
		return res
	}

	for i := 0; i < m; i++ {
		heap.push(Pair{val: a[0] + b[i], x: 0, y: i})
	}

	for len(res) < k && heap.size > 0 {
		t := heap.pop()
		res = append(res, []int{a[t.x], b[t.y]})
		if t.x+1 < n {
			heap.push(Pair{val: a[t.x+1] + b[t.y], x: t.x + 1, y: t.y})
		}
	}
	return res

}

type Pair struct {
	val int
	x   int
	y   int
}

type Heap struct {
	h    []Pair
	size int
}

func (heap *Heap) up(u int) {
	for u/2 != 0 && heap.h[u].val < heap.h[u/2].val {
		heap.h[u], heap.h[u/2] = heap.h[u/2], heap.h[u]
		u /= 2
	}
}

func (heap *Heap) down(u int) {
	t := u
	if u*2 <= heap.size && heap.h[t].val > heap.h[u*2].val {
		t = u * 2
	}
	if u*2+1 <= heap.size && heap.h[t].val > heap.h[u*2+1].val {
		t = u*2 + 1
	}

	if t != u {
		heap.h[u], heap.h[t] = heap.h[t], heap.h[u]
		heap.down(t)
	}
}

func (heap *Heap) push(x Pair) {
	heap.h = append(heap.h, x)
	heap.size++
	heap.up(heap.size)
}

func (heap *Heap) pop() Pair {
	val := heap.h[1]
	heap.h[1] = heap.h[heap.size]
	heap.h = heap.h[:heap.size]
	heap.size--
	heap.down(1)
	return val
}

// @lc code=end

