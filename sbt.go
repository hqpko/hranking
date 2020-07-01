package hrank

import "fmt"

type node struct {
	key    string
	source float64
	tree   *tree
}

type tree struct {
	left  *tree
	right *tree
	size  int
	node  *node
}

func add(t *tree, node *node) *tree {
	if t == nil {
		t = &tree{node: node}
		node.tree = t
	} else if node.source <= t.node.source {
		t.left = add(t.left, node)
	} else {
		t.right = add(t.right, node)
	}
	t.size++
	t = maintain(t, node.source > t.node.source)
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
		}
	} else if node.source <= t.node.source {
		t.left = del(t.left, node)
	} else {
		t.right = del(t.right, node)
	}
	t = maintain(t, false)
	t.size = size(t.left) + size(t.right) + 1
	return t
}

func rank(t *tree, node *node) int {
	if node.key == t.node.key {
		return size(t.left) + 1
	} else if node.source <= t.node.source {
		rank := rank(t.left, node)
		return rank
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

func size(t *tree) int {
	if t == nil {
		return 0
	}
	return t.size
}

func (t *tree) print() {
	t.debug("")
}

func (t *tree) debug(pre string) {
	if t == nil || t.node == nil {
		return
	}
	fmt.Println(pre, t.node.key, t.node.source)
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
