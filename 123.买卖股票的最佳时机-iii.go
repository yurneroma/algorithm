/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	return solution2(prices)
}

const MinInf = -1 << 31

// f[i][j] 代表前i天某种状态的最大值
func solution2(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	f := make([][]int, n+10)
	for i := 0; i < n+10; i++ {
		f[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			f[i][j] = MinInf
		}
		f[i][0] = 0
	}

	f[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][1] = max(f[i-1][0]-prices[i], f[i-1][1])
		f[i][2] = max(f[i-1][1]+prices[i], f[i-1][2])
		f[i][3] = max(f[i-1][3], f[i-1][2]-prices[i])
		f[i][4] = max(f[i-1][3]+prices[i], f[i-1][4])
	}
	//fmt.Println(f[5][0], f[5][1], f[5][2], f[5][3], f[5][4])
	return max(0, max(f[n-1][4], f[n-1][2]))
}

func solution1(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	firb := make([]int, n+10)
	secb := make([]int, n+10)
	cmin := prices[0]

	tmp := 0
	// 从前往后统计第一次买卖，第i天卖出最大收益
	for i := 0; i < n; i++ {
		if prices[i] <= cmin {
			cmin = prices[i]
			firb[i] = 0
		}
		tmp = max(tmp, prices[i]-cmin)
		firb[i] = tmp
	}

	cmax := prices[n-1]
	// 从后往前统计第二次买卖, 第i天买入最大收益
	tmp = 0
	for i := n - 1; i >= 0; i-- {
		if cmax <= prices[i] {
			cmax = prices[i]
		}
		tmp = max(tmp, cmax-prices[i])
		secb[i] = tmp
	}

	res := 0
	for i := 1; i <= n; i++ {
		fmt.Println(firb[i], secb[i+1])
		res = max(res, firb[i-1]+secb[i])
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

