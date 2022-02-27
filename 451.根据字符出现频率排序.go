/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
func frequencySort(s string) string {
	heap := Heap{h: make([]Pair, 1)}

	smap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		smap[s[i]]++
	}

	for e, val := range smap {
		heap.push(Pair{val: val, e: e})
	}

	res := make([]byte, 0)
	for heap.size > 0 {
		t := heap.pop()
		for i := 0; i < t.val; i++ {
			res = append(res, t.e)
		}
	}

	return string(res)
}

type Pair struct {
	val int
	e   byte
}

type Heap struct {
	h    []Pair
	size int
}

func (heap *Heap) up(u int) {
	for u/2 != 0 && heap.h[u].val > heap.h[u/2].val {
		heap.h[u], heap.h[u/2] = heap.h[u/2], heap.h[u]
		u /= 2
	}
}

func (heap *Heap) down(u int) {
	t := u
	if u*2 <= heap.size && heap.h[t].val < heap.h[u*2].val {
		t = u * 2
	}
	if u*2+1 <= heap.size && heap.h[t].val < heap.h[u*2+1].val {
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

