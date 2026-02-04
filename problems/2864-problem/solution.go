package main

func maximumOddBinaryNumber(str string) string {
  s, pts, last, odd := []byte(str), 0, len(str) - 1, false
  if s[last] == '1' { odd = true }
  for i:=0;i<len(s) - 1;i++ {
    if s[i] == '1' {
      if odd {
        s[i], s[pts] = s[pts], s[i]
        pts++
      } else {
        s[i], s[last] = s[last], s[i]
        odd = true
      }
    }
  }
  return string(s)
}
