package jmap

const count = 1024
type arrayList struct{
  array [count][]int
}

type JMap struct {
   arraylist []arrList
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
    }
    return jmap
}
