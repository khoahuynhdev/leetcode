package solution

func missingNumber(nums []int) int {
  ans := 0
  for i,v := range nums {
    ans ^= (i+1 ^ v)
  }
  return ans
}
