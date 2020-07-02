package hrank

import (
	"sync"
)

type Ranking struct {
	lock    sync.RWMutex
	tree    *tree
	nodeMap map[string]*node
}

func NewRanking() *Ranking {
	return &Ranking{nodeMap: map[string]*node{}}
}

// Set 设置 key 对应的分数
// 排名由大到小，得分相同情况下，最新更新的 key 排名更靠前
func (r *Ranking) Set(key string, source float64) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if n := r.nodeMap[key]; n != nil {
		if n.source == source {
			return
		}
		r.tree = del(r.tree, n)
		n.source = source
		r.tree = add(r.tree, n)
	} else {
		node := &node{key: key, source: source}
		r.tree = add(r.tree, node)
		r.nodeMap[key] = node
	}
}

// 由大到小排序，最大元素的排序为 1
func (r *Ranking) Get(key string) int {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if n := r.nodeMap[key]; n != nil {
		return len(r.nodeMap) - rank(r.tree, n) + 1
	}
	return 0
}
