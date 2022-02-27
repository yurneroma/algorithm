/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	size = 0
	cnt := 0
	h = make([]int, 1)
	seen = make(map[int]bool)
	h = append(h, 1)
	size++
	seen[1] = true

	for len(h) > 0 {
		cur := h[1]
		cnt++
		if cnt == n {
			return cur
		}

		ls := []int{2 * cur, 3 * cur, 5 * cur}
		for _, front := range ls {
			if seen[front] {
				continue
			}
			seen[front] = true
			h = append(h, front)
			size++
		}
		h[1] = h[size]
		h = h[:len(h)-1]
		size--
		for i := size / 2; i >= 1; i-- {
			down(i)
		}
	}

	return 1

}

var seen map[int]bool
var size int
var h []int

func up(u int) {
	for u/2 != 0 && h[u] < h[u/2] {
		h[u], h[u/2] = h[u/2], h[u]
		u /= 2
	}
}

func down(u int) {
	t := u
	if u*2 <= size && h[t] > h[u*2] {
		t = u * 2
	}
	if u*2+1 <= size && h[t] > h[u*2+1] {
		t = u*2 + 1
	}
	if t != u {
		h[t], h[u] = h[u], h[t]
		down(t)
	}
}

// @lc code=end

