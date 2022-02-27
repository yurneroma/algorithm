/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return heapSort(nums, k)
}

/*
基于快排的快选算法， 时间复杂度O(n), 证明: s = n/2+ n/4+ n/8 + ...= n
func quickSort(q []int, l, r, k int) int {
	if l == r {
		return q[l]
	}

	i, j := l-1, r+1
	x := q[(l+r)>>1]
	for i < j {
		for {
			i++
			if q[i] <= x {
				break
			}
		}
		for {
			j--
			if q[j] >= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	if j-l+1 >= k {
		return quickSort(q, l, j, k)
	} else {
		return quickSort(q, j+1, r, k-j+l-1)
	}
}
*/

/*
	基于堆排的算法，维护一个k大小的小顶堆， 遍历数组，比堆顶小的则忽略， 比堆顶大的则替换堆顶. 最终， 堆顶就是第k大的元素
*/

var h []int
var size int

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
		h[u], h[t] = h[t], h[u]
		down(t)
	}
}

func heapSort(nums []int, k int) int {
	h = make([]int, k+1)
	size = 0
	for i := 0; i < len(nums); i++ {
		if size < k {
			size++
			h[size] = nums[i]
			up(size)
		} else {
			if nums[i] > h[1] {
				h[1] = nums[i]
				down(1)
			}
		}

	}
	return h[1]
}

// @lc code=end

