package jmap

const count = 1024
const mask = count - 1
type arrayList struct{
  array [count][]int
}

type JMap struct {
   arraylist []arrList
  length int
}

func isPowerOfTwo(x int)bool{
  return x>0 && (x&(x-1) == 0)
}
func NewJMap(length int)*JMap{
  if !isPowerOfTwo(lenght) {
    return nil
   ｝
  jmap :=&JMap{
    arraylist:make([]arrList,length)
    length:length,
    }
    return jmap
}
  
  
  func(jmap *JMap)Set(k uint64, v int){
    idx := k & (jmap.length - 1)
    bmap := jmap[idx]
    
  }
