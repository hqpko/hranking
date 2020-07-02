package hranking

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
func (r *Ranking) Set(key string, score float64) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if n := r.nodeMap[key]; n != nil {
		if n.score == score {
			return
		}
		r.tree = del(r.tree, n)
		n.score = score
		r.tree = add(r.tree, n)
	} else {
		node := &node{key: key, score: score}
		r.tree = add(r.tree, node)
		r.nodeMap[key] = node
	}
}

// 由大到小排序，最大元素的排序为 1
func (r *Ranking) Get(key string) int {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if n := r.nodeMap[key]; n != nil {
		return rank(r.tree, n)
	}
	return 0
}

// 获取排名区间数据，返回区间内 key,score 集合，序号从 1 开始，如数据长度不足，则返回所有有效数据
// [from,to], if to<from, to=from
func (r *Ranking) GetRange(from, to int) ([]string, []float64) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if to < from {
		to = from
	}
	return getRange(r.tree, from, to, []string{}, []float64{})
}

// 获取排名为 n 的 key,score
// n>0 && n<=ranking.len，否则返回 "",0
func (r *Ranking) GetN(n int) (string, float64) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if t := getN(r.tree, n); t != nil {
		return t.node.key, t.node.score
	}
	return "", 0
}

func (r *Ranking) Len() int {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return len(r.nodeMap)
}

// index 从 1 开始
func (r *Ranking) Walk(handler func(index int, key string, score float64)) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	walk(r.tree, 1, handler)
}

func (r *Ranking) Copy() *Ranking {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return &Ranking{
		tree:    copyTree(r.tree),
		nodeMap: copyMap(r.nodeMap),
	}
}

func copyMap(m map[string]*node) map[string]*node {
	m2 := map[string]*node{}
	for k, v := range m {
		m2[k] = &node{key: v.key, score: v.score}
	}
	return m2
}
