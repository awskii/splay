package main

import (
	"math/rand"
	"time"

	"github.com/awskii/splay"
)

const elements = 800000

var random = make([]splay.Key, elements)

func init() {
	// pre-generate randomized numbers
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(random); i++ {
		random[i] = splay.Key(r.Int())
	}
}

func main() {
	T := splay.NewTree(1)

	for i := 0; i < elements; i++ {
		T.Insert(random[i])
	}

	shuffle(random)

	for i := 0; i < elements; i++ {
		T.Search(random[i])
	}
}

// Shuffle indices of randomized list
func shuffle(a []splay.Key) {
	rnd := rand.Perm(len(a))
	for i := 0; i < len(a); i++ {
		a[i], a[rnd[i]] = a[rnd[i]], a[i]
	}
}
