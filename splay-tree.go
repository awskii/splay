// package splay_tree
package splay

type Key int64

type Tree struct {
	root *Node
}

func NewTree(key Key) *Tree {
	return NewTreeFromNode(newNode(nil, nil, key))
}

func NewTreeFromNode(root *Node) *Tree {
	return &Tree{root: root}
}

func (t *Tree) Insert(key Key) *Node {
	L, R := t.Split(key)
	t.root = newNode(L, R, key)
	t.root.keepParent()
	return t.root
}

func (t *Tree) Remove(key Key) *Node {
	t.root = t.Search(key)
	t.root.left.setParent(nil)
	t.root.right.setParent(nil)
	el := t.root
	t.root = Merge(t.root.left, t.root.right)
	return el
}

func (t *Tree) Split(key Key) (*Node, *Node) {
	if t.root == nil {
		return nil, nil
	}

	var (
		L, R *Node
		root = t.root.find(key)
	)

	if root.key == key {
		root.left.setParent(nil)
		root.right.setParent(nil)
		return root.left, root.right
	}

	if root.key < key {
		R, root.right = root.right, nil
		R.setParent(nil)
		return root, R
	} else {
		L, root.left = root.left, nil
		L.setParent(nil)
		return L, root
	}
	return nil, nil
}

func (t *Tree) Search(key Key) *Node {
	t.root = t.root.find(key)
	return t.root
}

// All keys of L(eft) tree should be lesser than R(ight) keys
func Merge(L, R *Node) *Node {
	if L == nil {
		return R
	}
	if R == nil {
		return L
	}

	R = R.find(L.key)
	R.left, L.parent = L, R
	return R
}
