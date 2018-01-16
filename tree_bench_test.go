package splay

import (
	"math/rand"
	"runtime/debug"
	"testing"
	"time"
)

var T = NewTree(1)
var vals []Key

func init() {
	debug.SetGCPercent(-1)

	itemsCount := 5000000
	vals = make([]Key, itemsCount)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < itemsCount; i++ {
		vals[i] = Key(r.Int63n(int64(itemsCount)))
	}
}

func BenchmarkInsertIdle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T.Insert(Key(i))
	}
}

func BenchmarkSearchIdle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T.Search(Key(i))
	}
}

func BenchmarkInsertRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T.Insert(vals[i])
	}
}

func BenchmarkSearchRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := T.Search(vals[i])
		if e.key != vals[i] {
			b.Fatalf("Wrong elemnt got: %d want: %d", e.key, vals[i])
		}
	}
}
