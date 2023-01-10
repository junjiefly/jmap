package jmap

import (
	"sync"
	"sync/atomic"
)

const count = 1024
const mask = count - 1

func IsPowerOfTwo(x int) bool {
	return x > 0 && (x&(x-1)) == 0
}

type value struct {
	key   uint64
	value int //value can be any data type, not just int
}

type arrayList struct {
	array        [count][]value //for better performance, length(length = len(array[x])) should not be too large, better less than 100
	sync.RWMutex                //every list has its lock for better con-rw performance
}

type JMap struct {
	jmap   []arrayList
	count  int64
	length int
}

//length: cap/(count * 10)  #cap is the capacity of your k-v objects
//length should be a power value for 2
func NewJMap(length int) *JMap {
	if !IsPowerOfTwo(length) {
		return nil
	}
	jMap := &JMap{
		length: length,
	}
	for i := 0; i < length; i++ {
		wl := arrayList{}
		/*for ii := range wl.array {
			wl.array[ii] = make([]value,0)
		}*/
		jMap.jmap = append(jMap.jmap, wl)
	}
	return jMap
}

func (wmap *JMap) Set(key uint64, v int) {
	idx := uint64(wmap.length-1) & key
	crc := key / uint64(wmap.length)
	seq := mask & crc
	wmap.jmap[idx].Lock()
	for i := range wmap.jmap[idx].array[seq] {
		if wmap.jmap[idx].array[seq][i].key == key {
			wmap.jmap[idx].array[seq][i].value = v
			wmap.jmap[idx].Unlock()
			return
		}
	}
	wmap.jmap[idx].array[seq] = append(wmap.jmap[idx].array[seq], value{key: key, value: v})
	atomic.AddInt64(&wmap.count, 1)
	wmap.jmap[idx].Unlock()
}

func (jMap *JMap) Get(key uint64) (n *value, ok bool) {
	idx := uint64(jMap.length-1) & key
	crc := key / uint64(jMap.length)
	seq := mask & crc
	jMap.jmap[idx].RLock()
	for i := range jMap.jmap[idx].array[seq] {
		if jMap.jmap[idx].array[seq][i].key == key {
			n = &jMap.jmap[idx].array[seq][i]
			jMap.jmap[idx].RUnlock()
			return n, true
		}
	}
	jMap.jmap[idx].RUnlock()
	return nil, false
}

func (jMap *JMap) Delete(key uint64) bool {
	idx := uint64(jMap.length-1) & key
	crc := key / uint64(jMap.length)
	seq := mask & crc
	var sn = -1
	jMap.jmap[idx].Lock()
	for i := range jMap.jmap[idx].array[seq] {
		if jMap.jmap[idx].array[seq][i].key == key {
			sn = i
			break
		}
	}
	if sn >= 0 {
		length := len(jMap.jmap[idx].array[seq])
		lastNV := &jMap.jmap[idx].array[seq][length-1]
		myNV := &jMap.jmap[idx].array[seq][sn]
		if sn == length-1 {
		} else {
			myNV.key, myNV.value = lastNV.key, lastNV.value
		}
		jMap.jmap[idx].array[seq] = jMap.jmap[idx].array[seq][:length-1]
		atomic.AddInt64(&jMap.count, -1)
		jMap.jmap[idx].Unlock()
		return true
	}
	jMap.jmap[idx].Unlock()
	return false
}
