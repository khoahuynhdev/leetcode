package solution

func findJudge(n int, trust [][]int) int {
  if len(trust) == 0 && n > 1 { return -1 }
  if len(trust) == 0 && n == 1 { return n }
  trstLst := make(map[int][2]*[]int)
  for _, relation := range trust {
    if entry,ok := trstLst[relation[0]]; !ok {
      trstLst[relation[0]] = [2]*[]int{&[]int{relation[1]}, &[]int{}}
    } else {
      *entry[0] = append(*entry[0], relation[1])
    }
    if entry,ok := trstLst[relation[1]]; !ok {
      trstLst[relation[1]] = [2]*[]int{&[]int{}, &[]int{relation[0]}}
    } else {
      *entry[1] = append(*entry[1], relation[0])
    }
  }
  for k,v := range trstLst {
    // fmt.Println(k, v[0], v[1])
    if len(*v[0]) == 0 && len(*v[1]) == n - 1 { return k}
  }
  return -1
}
