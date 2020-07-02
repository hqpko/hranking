package hrank

import "sync"

type Ranking struct {
	lock    sync.RWMutex
	tree    *tree
	nodeMap map[string]*node
}

func NewRanking() *Ranking {
	return &Ranking{nodeMap: map[string]*node{}}
}

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

func (r *Ranking) Get(key string) int {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if n := r.nodeMap[key]; n != nil {
		return rank(r.tree, n)
	}
	return 0
}
