package hrank

import "fmt"

type node struct {
	key   string
	score float64
}

type tree struct {
	left  *tree
	right *tree
	size  int
	node  *node
}

func add(t *tree, node *node) *tree {
	if t == nil {
		return &tree{node: node, size: 1}
	} else if node.score < t.node.score {
		t.left = add(t.left, node)
	} else {
		t.right = add(t.right, node)
	}
	t.size++
	t = maintain(t, node.score > t.node.score)
	return t
}

func del(t *tree, node *node) *tree {
	if t == nil {
		return nil
	} else if t.node.key == node.key {
		if t.left == nil {
			return t.right
		} else if t.right == nil {
			return t.left
		} else {
			first := getFirst(t.right)
			t.node, first.node = first.node, t.node
			t.right = del(t.right, node)
		}
	} else if node.score <= t.node.score { // 注意此处是 <= 而不是 <，由于在插入时使用的是 < 导致相同 score 的节点在左边，此处需要使用 <= 来检查左方数据
		t.left = del(t.left, node)
		t = maintain(t, true)
	} else {
		t.right = del(t.right, node)
		t = maintain(t, false)
	}
	t.size--
	return t
}

// 由小到大，从 1 开始，0 表示无排序
func rank(t *tree, node *node) int {
	if t == nil {
		return 0
	}
	if node.key == t.node.key {
		return size(t.left) + 1
	} else if node.score < t.node.score {
		return rank(t.left, node)
	} else if t.right == nil {
		return size(t.left) + 1
	} else {
		return size(t.left) + rank(t.right, node) + 1
	}
}

func maintain(t *tree, flag bool) *tree {
	if !flag {
		if t.left != nil && size(t.left.left) > size(t.right) {
			t = rotateRight(t)
		} else if t.left != nil && size(t.left.right) > size(t.right) {
			t.left = rotateLeft(t.left)
			t = rotateRight(t)
		} else {
			return t
		}
	} else {
		if t.right != nil && size(t.right.right) > size(t.left) {
			t = rotateLeft(t)
		} else if t.right != nil && size(t.right.left) > size(t.left) {
			t.right = rotateRight(t.right)
			t = rotateLeft(t)
		} else {
			return t
		}
	}
	t.left = maintain(t.left, false)
	t.right = maintain(t.right, true)
	t = maintain(t, false)
	t = maintain(t, true)
	return t
}

func rotateRight(t *tree) *tree {
	if left := t.left; left != nil {
		t.left = left.right
		left.right = t
		left.size = t.size
		t.size = size(t.left) + size(t.right) + 1
		return left
	} else {
		return nil
	}
}

func rotateLeft(t *tree) *tree {
	if right := t.right; right != nil {
		t.right = right.left
		right.left = t
		right.size = t.size
		t.size = size(t.left) + size(t.right) + 1
		return right
	} else {
		return nil
	}
}

func getFirst(t *tree) *tree {
	for t.left != nil {
		t = t.left
	}
	return t
}

func check(t *tree) bool {
	if t.right != nil && (size(t.right.left) > size(t.left) || size(t.right.right) > size(t.left)) {
		return false
	} else if t.left != nil && (size(t.left.left) > size(t.right) || size(t.left.right) > size(t.right)) {
		return false
	} else if t.left != nil && !check(t.left) {
		return false
	} else if t.right != nil && !check(t.right) {
		return false
	} else {
		return true
	}
}

func size(t *tree) int {
	if t == nil {
		return 0
	}
	return t.size
}

func copyTree(t *tree) *tree {
	if t == nil {
		return nil
	}
	return &tree{
		left:  copyTree(t.left),
		right: copyTree(t.right),
		size:  t.size,
		node:  &node{key: t.node.key, score: t.node.score},
	}
}

func walk(t *tree, index int, handler func(index int, key string, score float64)) int {
	if t == nil {
		return index
	}
	index = walk(t.right, index, handler)
	handler(index, t.node.key, t.node.score)
	index++
	return walk(t.left, index, handler)
}

func (t *tree) print() {
	t.debug("")
}

func (t *tree) debug(pre string) {
	if t == nil || t.node == nil {
		return
	}
	fmt.Println(pre, t.node.key, t.node.score)
	pre += "-"
	if t.left != nil {
		fmt.Println("left:")
		t.left.debug(pre)
	}
	if t.right != nil {
		fmt.Println("right:")
		t.right.debug(pre)
	}
}
