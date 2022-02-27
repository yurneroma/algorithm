/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	/*
			单调队列方法， s作为滑动窗口， 维护当前窗口内的动态最大值。 s[0]用于保存当前窗口的最大值。
			后续的数如果比s末尾的数小， 可以加入s。 如果比s末尾的数大，那么需要从尾部, 一直剔除小的部分.
		ans := make([]int, 0)
		s := make([]int, 0)
		for i := 0; i < k; i++ {
			for len(s) > 0 && nums[i] > nums[s[len(s)-1]] {
				s = s[:len(s)-1]
			}
			s = append(s, i)
		}
		ans = append(ans, nums[s[0]])

		for i := k; i < len(nums); i++ {
			if i-s[0] >= k {
				s = s[1:]
			}

			for len(s) > 0 && nums[i] > nums[s[len(s)-1]] {
				s = s[:len(s)-1]
			}
			s = append(s, i)
			ans = append(ans, nums[s[0]])
		}
		return ans
	*/

	// heap 堆排方法

}

// @lc code=end

