package splay

import (
	"testing"
)

func TestSetParent(t *testing.T) {
	t.Parallel()

	node := newNode(nil, nil, 1)
	parent := newNode(nil, nil, 2)
	node.setParent(parent)
	if node.parent != parent {
		t.Fatalf("Node's parent is wrong: %v", node.parent)
	}
}

func TestKeepParent(t *testing.T) {
	t.Parallel()

	node := newNode(nil, nil, 1)
	node.keepParent()
	if node.left != nil && node.left.parent != node || node.right != nil && node.right.parent != node {
		t.Fatal("Parent of L or R is not equal to current node")
	}
}

func TestSplay(t *testing.T) {
	t.Parallel()

	// node without parent
	N := newNode(nil, nil, 1)
	res := N.Splay()
	if res != N {
		t.Fatalf("Node without parent after splay should be equal to himself, got: %v", res)
	}

	T := NewTree(1)
	for i := 0; i <= 10; i++ {
		T.Insert(Key(i))
	}

	// Test zig on pre-last inserted element
	res = T.root.left.Splay()
	if res == nil {
		t.Fatal("Returned element after splay cant be nil")
	} else if res.key != 9 {
		t.Fatalf("After splay should be returned node with key %d, got %d", 9, res.key)
	}

	if res.right == nil {
		t.Errorf("Right child element should be not nil, got %v", res.right)
	} else if res.right.key != 10 {
		t.Errorf("Right child should be equal to %d, got %d", 10, res.right.key)
	}

	if res.left == nil {
		t.Errorf("Left child element should be not nil, got %v", res.left)
	} else if res.left.key != 8 {
		t.Errorf("Left child should be equal to %d, got %d", 8, res.left.key)
	}

	// Test zig-zig, 7th element
	res = res.left.left.Splay()
	if res.left == nil || res.right == nil {
		t.Fatal("L and R subtrees should be not nil")
	}

	if res.left.key > res.key {
		t.Errorf("Left child key should be less or equal than %d, got %d", res.key, res.left.key)
	} else if res.right.key < res.key {
		t.Errorf("Right child key should be less or equal than %d, got %d", res.key, res.right.key)
	}

	// Test zig-zag on 9th element
	res = res.right.right.Splay()
	if res.key != 9 {
		t.Fatalf("Returned node should have key %d, got %d", 9, res.key)
	}

	if res.left == nil || res.right == nil {
		t.Fatal("L and R subtrees shouldn't be nil")
	} else if res.parent != nil {
		t.Fatal("After splay returned node shouldn't have a parent")
	}

	if res.left.key > res.key {
		t.Errorf("Left child key should be less or equal than %d, got %d", res.key, res.left.key)
	} else if res.right.key < res.key {
		t.Errorf("Right child key should be less or equal than %d, got %d", res.key, res.right.key)
	}
}

func TestRotate(t *testing.T) {

}

func TestFind(t *testing.T) {
	t.Parallel()

	T := NewTree(1)
	for i := 2; i <= 20; i++ {
		T.Insert(Key(i))
	}

	res := T.root.find(20)
	if res == nil {
		t.Fatalf("Result can be nil only if node is nil, got node %v", res)
	} else if res.key != 20 {
		t.Fatalf("Expected node with key %d but got %d", 20, res.key)
	}

	res = T.root.find(10)
	if res == nil {
		t.Fatalf("Result can be nil only if node is nil, got node %v", res)
	} else if res.key != 10 {
		t.Fatalf("Expected node with key %d but got %d", 10, res.key)
	}
}
