package main

func findDuplicates(nums []int) []int {
  ans := make([]int, 0)
  for _, n := range nums {
    p := Abs(n)
    if nums[p - 1] < 1 {
      ans = append(ans, p)
    } else {
      nums[p - 1] *= -1
    }
  }
  return ans
}

func Abs(n int) int{
  if n < 1 { return -1 * n}
  return n
}
