package main

func maxDepth(s string) int {
  ans, cnt := 0,0
  for _, c := range s {
    if c == '(' { cnt++ }
    if c == ')' { cnt--}
    if ans < cnt { ans = cnt}
  }
  return ans
}
