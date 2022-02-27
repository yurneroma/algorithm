/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	res := 0
	minNum := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > minNum {
			res = max(res, prices[i]-minNum)
		} else {
			minNum = prices[i]
		}
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

