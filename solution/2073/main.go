package main

func timeRequiredToBuy(tickets []int, k int) int {
  mt := tickets[k]
  ans := 0
  for idx, v := range tickets {
    if idx <= k {
      ans += Min(mt,v)
    } else {
      ans += Min(mt - 1,v)
    }
  }
  return ans
}

func Min(a, b int) int {
  if a < b { return a} else { return b}
}
