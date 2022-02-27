/*
 * @lc app=leetcode.cn id=502 lang=golang
 *
 * [502] IPO
 */

// @lc code=start
func findMaximizedCapital(k int, w int, pro []int, cap []int) int {
	//成本先排个序,每次从可选的项目中(当前利润>项目成本)选一个利润最大的，可以使当前利润最大，
	// 并且不影响后面的选择。所以这个题是个贪心做法。

	pl := make([]Pair, 0)
	for i := 0; i < len(cap); i++ {
		pl = append(pl, Pair{c: cap[i], p: pro[i]})
	}
	sort.Slice(pl, func(i, j int) bool { return pl[i].c < pl[j].c })

	heap := Heap{h: make([]Pair, 1)}
	var i int
	for i < len(pl) && pl[i].c <= w {
		heap.push(pl[i])
		i++
	}

	for k > 0 && heap.size > 0 {
		t := heap.pop()
		w += t.p
		for i < len(pl) && w >= pl[i].c {
			heap.push(pl[i])
			i++
		}
		k--
	}

	return w

}

type Pair struct {
	c int // cap
	p int // profit
}

type Heap struct {
	h    []Pair
	size int
}

// 大的上浮
func (heap *Heap) up(u int) {
	for u/2 != 0 && heap.h[u].p > heap.h[u/2].p {
		heap.h[u], heap.h[u/2] = heap.h[u/2], heap.h[u]
		u /= 2
	}
}

// 小的下沉
func (heap *Heap) down(u int) {
	t := u
	if u*2 <= heap.size && heap.h[u].p < heap.h[u*2].p {
		t = u * 2
	}
	if u*2+1 <= heap.size && heap.h[t].p < heap.h[u*2+1].p {
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

