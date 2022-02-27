/*
 * @lc app=leetcode.cn id=676 lang=golang
 *
 * [676] 实现一个魔法字典
 */

// @lc code=start
type MagicDictionary struct {
	son   [26]*MagicDictionary
	isEnd bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, w := range dictionary {
		node := this
		for i := 0; i < len(w); i++ {
			u := int(w[i] - 'a')
			if node.son[u] == nil {
				node.son[u] = &MagicDictionary{}
			}
			node = node.son[u]
		}
		node.isEnd = true
	}

}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.dfs(searchWord, 0, 0)
}

func (this *MagicDictionary) dfs(word string, idx, cnt int) bool {
	if idx == len(word) && this.isEnd && cnt == 1 {
		return true
	}

	if cnt > 1 || idx == len(word) {
		return false
	}

	for i := 0; i < 26; i++ {
		node := this.son[i]
		if node == nil {
			continue
		}

		if int(word[idx]-'a') != i {
			cnt++
		}

		if node.dfs(word, idx+1, cnt) {
			return true
		}
	}
	return false

}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end

