package main

func tribonacci(n int) int {
  dict := make(map[int]int)
  dict[0] = 0
  dict[1] = 1
  dict[2] = 1
  var fn func(int) int
  fn = func (num int) int {
    if num == 0 { return dict[0]}
    if num == 1 { return dict[1]}
    if num == 2 { return dict[2]}
    if v,ok:= dict[num];ok  { return v}
    v1,v2,v3 := fn(num-1), fn(num-2), fn(num-3)
    dict[num-1] = v1
    dict[num-2] = v2
    dict[num-3] = v3
    return v1 + v2 + v3
  }

  return fn(n)
}
