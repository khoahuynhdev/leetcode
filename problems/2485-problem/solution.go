package main

func pivotInteger(n int) int {
  l, r, total := 1, n, (n * (n+1)) / 2
  for l < r {
    m := (r + l)/2
    if m * m - total < 0 {
      l = m + 1
    } else {
      r = m 
    }
  }

  if l * l - total == 0 {
    return l
  }
  return -1
}
