/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	t := &Trie{}
	// 构建trie树
	for i := 0; i < len(words); i++ {
		t.insert(words[i])
	}

	m := len(board)
	n := len(board[0])

	dir := []struct{ x, y int }{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	res := make(map[string]bool)
	st := [12][12]bool{}
	var dfs func(i, j int, node *Trie)
	dfs = func(i, j int, node *Trie) {
		//ch := board[i][j]
		u := int(board[i][j] - 'a')
		node = node.son[u]
		if node == nil {
			return
		}

		if node.word != "" {
			res[node.word] = true
		}
		st[i][j] = true
		for _, d := range dir {
			newi := i + d.x
			newj := j + d.y
			if newi < m && newi >= 0 && newj < n && newj >= 0 && !st[newi][newj] {
				dfs(newi, newj, node)
			}
		}
		st[i][j] = false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, t)
		}
	}
	ans := []string{}
	for val := range res {
		ans = append(ans, val)
	}
	return ans
}

type Trie struct {
	son  [26]*Trie
	word string
}

func (t *Trie) insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		u := int(word[i] - 'a')
		if node.son[u] == nil {
			node.son[u] = &Trie{}
		}
		node = node.son[u]
	}
	node.word = word
}

// @lc code=end

