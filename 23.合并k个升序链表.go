/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var (
	h    []*ListNode
	size int
)

func up(x int) {
	for x/2 != 0 && h[x].Val <= h[x/2].Val {
		h[x], h[x/2] = h[x/2], h[x]
		x /= 2
	}
}

func down(x int) {
	t := x
	if x*2 <= size && h[x].Val >= h[2*x].Val {
		t = 2 * x
	}
	if x*2+1 <= size && h[t].Val >= h[2*x+1].Val {
		t = 2*x + 1
	}

	if t != x {
		h[x], h[t] = h[t], h[x]
		down(t)
	}

}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	if len(lists) == 0 {
		return nil
	}
	size = 0
	n := len(lists)
	h = make([]*ListNode, n+1)
	for i := 1; i <= n; i++ {
		if lists[i-1] == nil {
			continue
		}
		size++
		h[size] = lists[i-1]
		up(size)
	}

	cur := dummy

	for size > 0 {
		cur.Next = h[1]
		cur = cur.Next
		if h[1].Next != nil {
			h[1] = h[1].Next
		} else {
			h[1] = h[size]
			size--
		}
		down(1)
	}

	return dummy.Next
}

// @lc code=end