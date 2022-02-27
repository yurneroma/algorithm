/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	return solution1(prices)
}

func solution1(prices []int) int {
	//f[i][4]
	// 每天总共有四种状态， 0:不买不卖，非持股状态; 1:持股状态; 2:卖出状态; 3:冷冻状态
	n := len(prices)
	f := make([][]int, n+10)
	for i := 0; i < n+10; i++ {
		f[i] = make([]int, 4)
	}

	// 初始化
	f[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		// 4种状态转换
		f[i][0] = max(f[i-1][0], f[i-1][3])
		f[i][1] = max(f[i-1][1], max(f[i-1][0]-prices[i], f[i-1][3]-prices[i]))
		f[i][2] = f[i-1][1] + prices[i]
		f[i][3] = f[i-1][2]
	}
	res := max(max(f[n-1][0], f[n-1][3]), f[n-1][2])
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

