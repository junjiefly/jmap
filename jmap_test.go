package jmap

import (
	"fmt"
	"testing"
)

var jMap *JMap

var needleCount = 1000000

var mmap map[uint64]int

func TestMain(m *testing.M) {
	jMap = NewJMap(1024) //length has to be the power of 2
	mmap = make(map[uint64]int,needleCount)
	//for i := 0; i < needleCount; i++ {
	//	jMap.Set(uint64(i), i)
	//	mmap[uint64(i)] = i
	//}
	fmt.Println("j-map ready! len:", jMap.count)
	fmt.Println("goMap ready! len:", len(mmap))
	m.Run()
}

func BenchmarkJMapSet(b *testing.B) {
	for i:=0;i<b.N;i++{
		jMap.Set(uint64(i), 0)
	}
}

func BenchmarkJMapGet(b *testing.B) {
	for i:=0;i<b.N;i++{
		_, _ = jMap.Get(uint64(i))
	}
}

func BenchmarkJMapDelete(b *testing.B) {
	for i:=0;i<b.N;i++{
		_ = jMap.Delete(uint64(i))
	}
}

func BenchmarkMapSet(b *testing.B) {
	for i:=0;i<b.N;i++{
		mmap[uint64(i)] = 0
	}
}

func BenchmarkMapGet(b *testing.B) {
	for i:=0;i<b.N;i++{
		_, _ = mmap[uint64(i)]
	}
}

func BenchmarkMapDelete(b *testing.B) {
	for i:=0;i<b.N;i++{
		delete(mmap, uint64(i))
	}
}
