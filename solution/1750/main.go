package main

func minimumLength(s string) int {
  l, r := 0, len(s) - 1
  for l < r {
    if s[l] != s[r] { break }
    tmp := s[l]
    for s[l] == tmp && l < r{
      l++
    }
    if l == r { return 0}
    for s[r] == tmp && r > l{
      r--
    }
    // fmt.Println(l, r)
  }
  return r - l + 1
}
