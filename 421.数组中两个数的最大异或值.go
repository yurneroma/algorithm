const N = 10000 * 64

var son [N][2]int
var idx int

func findMaximumXOR(nums []int) int {
	idx = 0
	son = [N][2]int{}
	// build trie tree
	for _, num := range nums {
		insert(num)
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(ans, query(nums[i]))
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func insert(x int) {
	p := 0
	for i := 31; i >= 0; i-- {
		u := (x >> i) & 1
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
}

func query(x int) int {
	p := 0
	res := 0
	for i := 31; i >= 0; i-- {
		u := (x >> i) & 1
		if son[p][1-u] > 0 {
			res += (1 << i)
			p = son[p][1-u]
		} else {
			p = son[p][u]
		}
	}
	return res
}

// @lc code=end