/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
func replaceWords(dic []string, s string) string {
	// build trie tree
	t := &Trie{}

	for _, e := range dic {
		t.insert(e)
	}

	res := make([]string, 0)
	//check
	j := 0
	for i := 0; i < len(s); i++ {
		j = i
		for j < len(s) && s[j] != ' ' {
			j++
		}
		word := s[i:j]
		res = append(res, t.query(word))
		i = j
	}
	return strings.Join(res, " ")

}

type Trie struct {
	son   [26]*Trie
	isEnd bool
}

func (t *Trie) insert(w string) {
	node := t
	for i := 0; i < len(w); i++ {
		u := int(w[i] - 'a')
		if node.son[u] == nil {
			node.son[u] = &Trie{}
		}
		node = node.son[u]
	}
	node.isEnd = true
}

func (t *Trie) query(w string) string {
	node := t
	for i := 0; i < len(w); i++ {
		u := int(w[i] - 'a')
		if node.son[u] == nil {
			return w
		}
		node = node.son[u]
		if node.isEnd {
			return w[:i+1]
		}

	}
	return w
}

// @lc code=end

