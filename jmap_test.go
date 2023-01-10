package jmap

import (
	"fmt"
	"math/rand"
	"testing"
)

var jMap *JMap

var needleCount = 10000000

var mmap map[uint64]int

func TestMain(m *testing.M) {
	jMap = NewJMap(1024)
	mmap = make(map[uint64]int)
	for i := 0; i < needleCount; i++ {
		jMap.Set(uint64(i), i)
		mmap[uint64(i)] = i
	}
	fmt.Println("j-map ready! len:", jMap.count)
	fmt.Println("goMap ready! len:", len(mmap))
	m.Run()
}

func Benchmark_JMapSet(b *testing.B) {
	k := rand.Intn(needleCount)
	jMap.Set(uint64(k), 0)
}

func Benchmark_JMapGet(b *testing.B) {
	k := rand.Intn(needleCount)
	_, _ = jMap.Get(uint64(k))
}

func Benchmark_JMapDel(b *testing.B) {
	k := rand.Intn(needleCount)
	_ = jMap.Delete(uint64(k))
}

func Benchmark_MapSet(b *testing.B) {
	k := rand.Intn(needleCount)
	mmap[uint64(k)] = 0
}

func Benchmark_MapGet(b *testing.B) {
	k := rand.Intn(needleCount)
	_, _ = mmap[uint64(k)]
}

func Benchmark_MapDelete(b *testing.B) {
	k := rand.Intn(needleCount)
	delete(mmap, uint64(k))
}
