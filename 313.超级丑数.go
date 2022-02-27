/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
type H []int

func (h *H) Len() int {
	return len(*h)
}

func (h *H) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *H) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *H) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *H) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
func nthSuperUglyNumber(n int, primes []int) int {
	var res int
	hash := make(map[int]bool)
	h := H{1}
	heap.Init(&h)

	for n > 0 {
		res = heap.Pop(&h).(int)
		n--
		if n == 0 {
			return res
		}
		for _, v := range primes {
			if !hash[res*v] {
				heap.Push(&h, res*v)
				hash[res*v] = true
			}
		}
	}
	return -1
}

/*
func nthSuperUglyNumber(n int, primes []int) int {
	// 使用堆排会TLE， 算法复杂度nlogn
	/*
			heap := &Heap{h: make([]int, 1)}
			seen := make(map[int]bool)
			cnt := 0
			heap.push(1)
		seen[1] = true

			for heap.size > 0 {
				x := heap.h[1]
				cnt++
			i	f cnt == n {
					return x
				}
			f	or _, e := range primes {
					if seen[e*x] {
						continue
					}
					seen[e*x] = true
					heap.h = append(heap.h, e*x)
					heap.size++
			}

				heap.h[1] = heap.h[heap.size]
				heap.h = heap.h[:len(heap.h)-1]
				heap.size--
			f	or i := heap.size / 2; i >= 1; i-- {
					heap.down(i)
			}

			}
		return -1
*/

/*
	//多路归并 + 堆排 , 复杂度 o(n)
	// init heap
	if n == 1 {
		return 1
	}

	ans := make([]int, 0)
	ans = append(ans, 1)
	heap := &Heap{h: make([]int, 1)}
	// build index
	imap := make(map[int]int)
	seen := make(map[int]bool)
	for _, elem := range primes {
		heap.push(elem)
		imap[elem] = 0
		seen[elem] = true
	}

	var cnt = 1
	for len(ans) < n {
		m := heap.pop()
		cnt++
		//fmt.Printf("min=%d,cnt=%d\n", m, cnt)
		ans = append(ans, m)
		if cnt == n {
			return m
		}

		for elem, idx := range imap {
			if elem*ans[idx] == m {
				idx++
				imap[elem] = idx
				val := elem * ans[idx]
				if !seen[val] {
					heap.push(val)
					seen[val] = true
				}
			}
		}
	}
	return -1

}

type Heap struct {
	h    []int
	size int
}

func (heap *Heap) up(u int) {
	for u/2 != 0 && heap.h[u] < heap.h[u/2] {
		heap.h[u], heap.h[u/2] = heap.h[u/2], heap.h[u]
		u /= 2
	}
}

func (heap *Heap) down(u int) {
	t := u
	if u*2 <= heap.size && heap.h[u] > heap.h[u*2] {
		t = u * 2
	}
	if u*2+1 <= heap.size && heap.h[t] > heap.h[u*2+1] {
		t = u*2 + 1
	}
	if t != u {
		heap.h[u], heap.h[t] = heap.h[t], heap.h[u]
		heap.down(t)
	}
}
func (heap *Heap) pop() int {
	res := heap.h[1]
	heap.h[1] = heap.h[len(heap.h)-1]
	heap.h = heap.h[:len(heap.h)-1]
	heap.size--
	heap.down(1)
	return res
}

func (heap *Heap) top() int {
	if heap.size == 0 {
		return 0
	}
	return heap.h[1]
}

func (heap *Heap) push(x int) {
	heap.h = append(heap.h, x)
	heap.size++
	heap.up(heap.size)
}
*/

// @lc code=end

