/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 */

// @lc code=start
type Twitter struct {
	tmap map[int][]elem       // key is userID,  val is twitter elem list
	fmap map[int]map[int]bool // key is followerID,  val followeedID list
}

func Constructor() Twitter {
	return Twitter{
		tmap: make(map[int][]elem),
		fmap: make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if this.tmap[userId] == nil {
		this.tmap[userId] = make([]elem, 0)
	}
	this.tmap[userId] = append(this.tmap[userId], elem{tId: tweetId, utime: time.Now().UnixNano()})
}

type elem struct {
	tId   int
	utime int64
}
type Heap struct {
	h    []elem
	size int
}

func (h *Heap) up(u int) {
	for u/2 != 0 && h.h[u].utime > h.h[u/2].utime {
		h.h[u], h.h[u/2] = h.h[u/2], h.h[u]
		u /= 2
	}
}

func (h *Heap) down(u int) {
	t := u
	if u*2 <= h.size && h.h[u].utime < h.h[u*2].utime {
		t = u * 2
	}
	if u*2+1 <= h.size && h.h[t].utime < h.h[u*2+1].utime {
		t = u*2 + 1
	}
	if t != u {
		h.h[t], h.h[u] = h.h[u], h.h[t]
		h.down(t)
	}
}

func (h *Heap) push(x elem) {
	h.h = append(h.h, x)
	h.size++
	h.up(h.size)
}

func (h *Heap) pop() elem {
	val := h.h[1]
	h.h[1] = h.h[h.size]
	h.h = h.h[:h.size]
	h.size--
	h.down(1)
	return val
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	// 获取关注usrId列表
	flist := make([]int, 0)
	for uid := range this.fmap[userId] {
		flist = append(flist, uid)
	}
	flist = append(flist, userId)

	heap := &Heap{h: make([]elem, 1)}
	// 每个用户的数组索引
	for _, uId := range flist {
		elems := this.tmap[uId]
		for _, elem := range elems {
			heap.push(elem)
		}
	}

	res := make([]int, 0)
	for heap.size > 0 {
		val := heap.pop()
		//fmt.Printf("val=%+v\n", val)
		res = append(res, val.tId)
		if len(res) == 10 {
			break
		}
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.fmap[followerId] == nil {
		this.fmap[followerId] = make(map[int]bool)
	}
	this.fmap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.fmap[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

