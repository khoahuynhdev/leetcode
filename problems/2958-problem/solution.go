package main

func maxSubarrayLength(nums []int, k int) int {
  f,l,r,ans := make(map[int]int),0,0,0
  for _,n := range nums {
    f[n]++
    r++
    for f[n] > k {
      f[nums[l]]--
      l++
    }
    if ans < r - l { ans = r - l }
  }
  return ans
}
