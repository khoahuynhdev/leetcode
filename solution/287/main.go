package main

import (
  "sort"
)

func findDuplicate(nums []int) int {
  // n == 4
  // a b c d b
  // a a a c d
  // 1 3 3 3 4

  sort.Ints(nums)
  l, r := nums[0], nums[1]
  for nums[l] != nums[r] && r < len(nums) {
    l++
    r++
  }
  return nums[r]
}
