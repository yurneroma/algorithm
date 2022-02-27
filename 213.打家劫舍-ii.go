/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	// f[i][j] 前i个抢不抢的最大值, j = (0|1)
	// 状态机dp， 抢和不抢间转换
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	f := make([][2]int, n+1)
	// 抢第一个
	f[0][1] = nums[0]
	f[0][0] = -1 << 31
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][1], f[i-1][0])
		f[i][1] = f[i-1][0] + nums[i]
	}
	res := f[n-1][0]
	// 不抢第一个
	f[0][0] = 0
	f[0][1] = -1 << 31
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][1], f[i-1][0])
		f[i][1] = f[i-1][0] + nums[i]
	}
	res = max(res, max(f[n-1][0], f[n-1][1]))
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

