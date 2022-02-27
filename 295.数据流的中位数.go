/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type MedianFinder struct {
	minH MinHeap
	maxH MaxHeap
}

type MinHeap struct {
	h    []int
	size int
}

type MaxHeap struct {
	h    []int
	size int
}

func Constructor() MedianFinder {
	return MedianFinder{
		minH: MinHeap{h: make([]int, 1)},
		maxH: MaxHeap{h: make([]int, 1)},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxH.size == 0 || num <= this.maxH.top() {
		this.maxH.push(num)
		if this.maxH.size > this.minH.size+1 {
			res := this.maxH.pop()
			this.minH.push(res)
		}
	} else {
		this.minH.push(num)
		if this.minH.size > this.maxH.size {
			res := this.minH.pop()
			this.maxH.push(res)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.minH.size+this.maxH.size)%2 == 0 {
		return float64(this.minH.top()+this.maxH.top()) / 2.0
	}
	return float64(this.maxH.top())
}

func (h *MinHeap) up(u int) {
	for u/2 != 0 && h.h[u] < h.h[u/2] {
		h.h[u], h.h[u/2] = h.h[u/2], h.h[u]
		u /= 2
	}
}

func (h *MinHeap) down(u int) {
	t := u
	if u*2 <= h.size && h.h[u] > h.h[u*2] {
		t = u * 2
	}

	if u*2+1 <= h.size && h.h[t] > h.h[u*2+1] {
		t = u*2 + 1
	}

	if t != u {
		h.h[t], h.h[u] = h.h[u], h.h[t]
		h.down(t)
	}
}

func (h *MaxHeap) up(u int) {
	for u/2 != 0 && h.h[u] > h.h[u/2] {
		h.h[u], h.h[u/2] = h.h[u/2], h.h[u]
		u /= 2
	}
}

func (h *MaxHeap) down(u int) {
	t := u
	if u*2 <= h.size && h.h[u] < h.h[u*2] {
		t = u * 2
	}

	if u*2+1 <= h.size && h.h[t] < h.h[u*2+1] {
		t = u*2 + 1
	}

	if t != u {
		h.h[t], h.h[u] = h.h[u], h.h[t]
		h.down(t)
	}

}

func (h *MinHeap) top() int {
	if h.size == 0 {
		return 0
	}
	return h.h[1]
}

func (h *MaxHeap) top() int {
	if h.size == 0 {
		return 0
	}
	return h.h[1]
}

func (h *MaxHeap) push(num int) {
	h.h = append(h.h, num)
	h.size++
	h.up(h.size)
}

func (h *MinHeap) push(num int) {
	h.h = append(h.h, num)
	h.size++
	h.up(h.size)
}

func (h *MaxHeap) pop() int {
	res := h.h[1]
	h.h[1] = h.h[h.size]
	h.h = h.h[:len(h.h)-1]
	h.size--
	h.down(1)
	return res
}

func (h *MinHeap) pop() int {
	res := h.h[1]
	h.h[1] = h.h[h.size]
	h.h = h.h[:len(h.h)-1]
	h.size--
	h.down(1)
	return res
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

