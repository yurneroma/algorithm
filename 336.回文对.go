/*
 * @lc app=leetcode.cn id=336 lang=golang
 *
 * [336] 回文对
 */

// @lc code=start
func palindromePairs(words []string) [][]int {
	res := make([][]int, 0)
	hash := make(map[string]int)
	for i := 0; i < len(words); i++ {
		hash[reverse(words[i])] = i + 1
	}

	for i := 0; i < len(words); i++ {
		word := words[i]
		for j := 0; j < len(word); j++ {
			left := word[:j]
			right := word[j:]
			if check(right) && hash[left] != 0 && hash[left]-1 != i {
				res = append(res, []int{i, hash[left] - 1})
			}

			if hash[right] != 0 && check(left) && hash[right]-1 != i && len(word) != len(words[hash[right]]) {
				res = append(res, []int{hash[right] - 1, i})
			}
		}
	}

	return res
}

func reverse(word string) string {
	buff := []byte{}
	for i := len(word) - 1; i >= 0; i-- {
		buff = append(buff, word[i])
	}
	return string(buff)
}

func check(word string) bool {
	i := 0
	j := len(word) - 1
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end

