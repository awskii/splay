package splay

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	t.Parallel()

	T := NewTree(1)
	if T == nil {
		t.Fatal("Got nil tree")
	}
	if T.root == nil {
		t.Fatal("Got nil tree root")
	}

	if T.root.left != nil || T.root.right != nil {
		t.Fatal("L or R of root is not nil")
	}

	if T.root.parent != nil {
		t.Fatal("Root's parent is not nil")
	}
}

func TestInsert(t *testing.T) {
	t.Parallel()

	T := NewTree(1)
	T.Insert(2)
	if T.root.key != 2 {
		t.Fatalf("Inserted %d but root element has key %d", 2, T.root.key)
	}

	if T.root.left == nil || T.root.left.key != 1 {
		t.Fatalf("Wrong splay over first element")
	}

	if T.root.right != nil {
		t.Fatalf("Non-empty R node after insertion 1-2 with value %d", T.root.right.key)
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()

	T := NewTree(1)
	T.Insert(2)
	T.Insert(3)
	T.Insert(3)
	T.Insert(10)
	T.Insert(101)
	T.Insert(60)
	T.Insert(35)
	T.Insert(47)

	lookup := Key(10)
	if n := T.Search(lookup); n.key != lookup {
		t.Fatalf("Searched element was %d but got %d", lookup, n.key)
	}

	if T.root.key != lookup {
		t.Fatalf("Root element is not equals to lookup element: %d", T.root.key)
	}

	if n := T.Search(1000000); n.key != 101 {
		t.Fatalf("Should return maximum key, but got %d", n.key)
	}

	if n := T.Search(-50); n.key != 1 {
		t.Fatalf("Should return minimum key, but got %d", n.key)
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()

	T := NewTree(1)
	T.Insert(2)
	T.Insert(3)

	el := T.Remove(2)
	if el.key != 2 {
		t.Fatalf("Returned node is not equals to removed: %d", el.key)
	}

	// if el.left != nil || el.right != nil {
	// t.Fatal("Returned node has non-flushed L or R pointers")
	// }

	el = T.Search(2)
	if el == nil {
		t.Fatal("After removal Search of not existed element should return new root, not nil")
	}

	if el.key == 2 {
		t.Fatalf("Element was not properly removed: %+v", el)
	}
}

func TestSplit(t *testing.T) {
	t.Parallel()

	T := new(Tree)
	L, R := T.Split(1)
	if L != nil || R != nil {
		t.Fatal("Left or Right subtree is not nil")
	}

	T = NewTree(1)
	T.Insert(2)
	T.Insert(3)

	L, R = T.Split(3)
	if L == nil {
		t.Fatal("Left subtree shouldn't be nil")
	} else if L.key != 2 {
		t.Fatalf("Left subtree root key should be %d but got %d", 2, L.key)
	}

	if R != nil {
		t.Fatalf("Left subtree should be nil, but got %+v", R)
	}

	T = NewTree(1)
	T.Insert(2)
	T.Insert(3)

	L, R = T.Split(2)
	if L == nil {
		t.Fatal("Left subtree shouldn't be nil")
	} else if L.key != 1 {
		t.Fatalf("Left subtree root key should be %d but got %d", 1, L.key)
	}

	if R == nil {
		t.Fatal("Right subtree shouldn't be nil")
	} else if R.key != 3 {
		t.Fatalf("Left subtree root key should be %d but got %d", 3, R.key)
	}
}

func TestMerge(t *testing.T) {
	var (
		L = newNode(nil, nil, 1)
		R *Node
	)

	T := Merge(L, R)
	if T != L {
		t.Fatalf("Merge with nil right subtree should return only left subtree, got %+v", T)
	}

	L = new(Node)
	R = newNode(nil, nil, 3)
	T = Merge(L, R)
	if T != R {
		t.Fatalf("Merge with nil left subtree should return only right subtree, got %+v", T)
	}

	L = newNode(nil, nil, 1)
	R = newNode(nil, nil, 3)
	T = Merge(L, R)
	if T.key != 3 {
		t.Fatalf("Root key should be %d but got %d", 3, T.key)
	}
	if T.left == nil {
		t.Fatal("Left child of returned root should be not nil")
	} else if T.left.key != 1 {
		t.Fatalf("Left child of returned root should be equal to %d but got %d", 1, T.left.key)
	}
}
