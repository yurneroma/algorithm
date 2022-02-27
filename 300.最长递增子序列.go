/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	arr := make([]int, 0)
	arr = append(arr, nums[0])
	for _, num := range nums {
		if num <= arr[0] {
			arr[0] = num
		} else if num > arr[len(arr)-1] {
			arr = append(arr, num)
		} else {
			// 二分
			l := 0
			r := len(arr) - 1
			for l < r {
				mid := (l + r - 1) >> 1
				if arr[mid] <= num {
					l = mid + 1
				} else {
					r = mid
				}
			}
			if arr[l-1] == num {
				continue
			}
			arr[l] = num

		}
	}
	return len(arr)
}

// @lc code=end

