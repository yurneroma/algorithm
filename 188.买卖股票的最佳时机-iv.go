/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start

const MIN = -1 << 31

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k >= n/2 {
		res := 0
		for i := 0; i+1 < n; i++ {
			if prices[i+1] > prices[i] {
				res += (prices[i+1] - prices[i])
			}
		}
		return res
	}

	// f[i][j] 代表前i天交易了j次， 且手中没有股票的情况
	// g[i][j] 代表前i天交易了j次， 且手中有股票的情况

	f := make([][]int, n+1)
	g := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, n+1)
		g[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			f[i][j] = MIN
			g[i][j] = MIN

		}
	}
	f[0][0] = 0
	res := 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			f[i][j] = f[i-1][j]
			g[i][j] = max(g[i-1][j], f[i-1][j]-prices[i-1])
			if j > 0 {
				f[i][j] = max(f[i][j], g[i-1][j-1]+prices[i-1])
			}
			res = max(res, f[i][j])
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

