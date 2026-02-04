package main

func sortedSquares(nums []int) []int {
  l, r, ans := 0, len(nums) - 1, make([]int, len(nums))
  for i:=len(ans) - 1;i>=0;i-- {
    sl, sr := Abs(nums[l]), Abs(nums[r])
    if sr > sl {
      ans[i] = sr*sr
      r--
    } else {
      ans[i] = sl*sl
      l++
    }
  }
  return ans
}

func Abs(n int) int {
  if n < 0 { return -n}
  return n
}
