/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	f[0] = nums[0]
	ans := f[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
		ans = max(ans, f[i])
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

