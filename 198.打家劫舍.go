/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	// f[i][1]代表前i天rob 第i个的最大值
	// f[i][0] 代表前i天 不rob 第i个的最大值
	n := len(nums)
	f := make([][2]int, n+1)
	f[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i][0] = max(f[i-1][1], f[i-1][0])
		f[i][1] = f[i-1][0] + nums[i]
	}

	return max(f[n-1][0], f[n-1][1])
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

