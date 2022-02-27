/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	heap := &Heap{h: make([]*Item, 1)}
	for name, val := range m {
		heap.push(&Item{name: name, val: val})
	}

	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, heap.pop().name)
	}
	return res
}

type Item struct {
	name int
	val  int
}
type Heap struct {
	h    []*Item
	size int
}

func (h *Heap) up(u int) {
	for u/2 != 0 && h.h[u].val > h.h[u/2].val {
		h.h[u], h.h[u/2] = h.h[u/2], h.h[u]
		u /= 2
	}
}

func (h *Heap) down(u int) {
	t := u
	if u*2 <= h.size && h.h[u].val < h.h[u*2].val {
		t = u * 2
	}

	if u*2+1 <= h.size && h.h[t].val < h.h[u*2+1].val {
		t = u*2 + 1
	}
	if t != u {
		h.h[u], h.h[t] = h.h[t], h.h[u]
		h.down(t)
	}
}

func (h *Heap) pop() *Item {
	item := h.h[1]
	h.h[1] = h.h[len(h.h)-1]
	h.h = h.h[:len(h.h)-1]
	h.size--
	h.down(1)
	return item
}

func (h *Heap) push(item *Item) {
	h.h = append(h.h, item)
	h.size++
	h.up(h.size)
}

// @lc code=end

