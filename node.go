package splay

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	key    Key
}

func newNode(l, r *Node, key Key) *Node {
	return &Node{left: l, right: r, key: key}
}

func (n *Node) setParent(p *Node) {
	if n != nil {
		n.parent = p
	}
}

func (n *Node) keepParent() {
	n.left.setParent(n)
	n.right.setParent(n)
}

func (n *Node) Rotate(child *Node) {
	gp := n.parent
	if gp != nil {
		if gp.left == n {
			gp.left = child
		} else {
			gp.right = child
		}
	}

	if n.left == child {
		n.left, child.right = child.right, n
	} else {
		n.right, child.left = child.left, n
	}

	child.keepParent()
	n.keepParent()
	child.parent = gp
}

func (n *Node) Splay() *Node {
	if n.parent == nil {
		return n
	}

	parent := n.parent
	gparent := parent.parent

	for {
		// if n.parent == nil {
		// return n
		// }
		if gparent == nil {
			// make zig
			parent.Rotate(n)
			return n
		}

		if gparent.left == parent && parent.left == n {
			// zig-zig
			gparent.Rotate(parent)
			parent.Rotate(n)
		} else {
			// zig-zag
			parent.Rotate(n)
			gparent.Rotate(n)
		}

		if parent = n.parent; parent == nil {
			break
		}

		if gparent = parent.parent; gparent == nil {
			parent.Rotate(n)
			break
		}
	}
	return n

	// parent := n.parent
	// gparent := parent.parent

	// if gparent == nil {
	// // make zig
	// parent.Rotate(n)
	// return n
	// }

	// if gparent.left == parent && parent.left == n {
	// // zig-zig
	// gparent.Rotate(parent)
	// parent.Rotate(n)
	// } else {
	// // zig-zag
	// parent.Rotate(n)
	// gparent.Rotate(n)
	// }
	// return n.Splay()
}

func (n *Node) find(key Key) *Node {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n.Splay()
		// } else if n.key > key && n.left != nil {
		// return n.left.find(key)
		// } else if n.key < key && n.right != nil {
		// return n.right.find(key)
		// }
	}

	var curNode *Node
	if n.key > key && n.left != nil {
		curNode = n.left
	} else if n.key < key && n.right != nil {
		curNode = n.right
	} else {
		return n.Splay()
	}

	for {
		if curNode.key == key {
			break
		}
		if curNode.key > key && curNode.left != nil {
			curNode = curNode.left
			continue
		} else if curNode.key < key && curNode.right != nil {
			curNode = curNode.right
			continue
		}
		break
	}
	return curNode.Splay()
}
