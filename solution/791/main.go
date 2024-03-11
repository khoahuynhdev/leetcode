package main

# running: 0ms

func customSortString(order string, s string) string {
  ans, dict := make([]rune,0), make(map[rune]int)
  for _,v := range s {
    dict[v]++
  }
  for _,v := range order {
    tmp := dict[v]
    for tmp > 0 {
      ans = append(ans, v)
      tmp--
    }
    delete(dict,v)
  }
  for k,v := range dict {
    tmp := v
    for tmp > 0 {
      ans = append(ans, k)
      tmp--
    }
  }

  return string(ans)
}
