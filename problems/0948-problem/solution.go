package main

import (
  "sort"
)

// Intuition: sort to get most power -> by it with score to gain retain max score after that

func bagOfTokensScore(tokens []int, power int) int {
  score, l, r := 0,0, len(tokens)
  sort.Ints(tokens)
  for l < r {
    if tokens[l] <= power {
      power-= tokens[l]
      l++
      score++
    } else {
      r--
      if r == l { break }
      if score < 1 { break }
      power+=tokens[r]
      score--
    }
  }
  return score
}
