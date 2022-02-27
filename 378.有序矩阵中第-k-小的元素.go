/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start
func kthSmallest(a [][]int, k int) int {
	heap := &Heap{h: make([]Pair, 1)}
	n := len(a)
	for i := 0; i < n; i++ {
		heap.push(Pair{val: a[i][0], x: i, y: 0})
	}

	cnt := 0
	for cnt < k {
		t := heap.pop()
		cnt++
		if cnt == k {
			return t.val
		}
		if t.y+1 < n {
			heap.push(Pair{val: a[t.x][t.y+1], x: t.x, y: t.y + 1})
		}
	}
	return -1
}

type Pair struct {
	val int // val
	x   int // row
	y   int // column
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
	if u*2 <= heap.size && heap.h[u].val > heap.h[u*2].val {
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

func (heap *Heap) push(pair Pair) {
	heap.h = append(heap.h, pair)
	heap.size++
	heap.up(heap.size)
}

func (heap *Heap) pop() Pair {
	x := heap.h[1]
	heap.h[1] = heap.h[heap.size]
	heap.h = heap.h[:heap.size]
	heap.size--
	heap.down(1)
	return x
}

// @lc code=end

