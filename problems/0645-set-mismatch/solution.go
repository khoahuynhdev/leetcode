package solution

func findErrorNums(nums []int) []int {
	s := make([]int, len(nums)+1)
	s[0] = 1
	ans := []int{0, 0}
	for _, v := range nums {
		s[v]++
	}
	for idx, v := range s {
		if v > 1 {
			ans[0] = idx
		}
		if v == 0 {
			ans[1] = idx
		}
	}
	return ans
}
