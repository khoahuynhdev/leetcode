package main

import (
  "strings"
)

func lengthOfLastWord(s string) int {
  words := strings.Fields(strings.TrimSpace(s))
  return len(words[len(words)-1]) 
}
