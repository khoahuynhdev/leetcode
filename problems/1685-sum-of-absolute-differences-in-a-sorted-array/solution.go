package solution

// https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/editorial/?envType=daily-question&envId=2023-11-25
func getSumAbsoluteDifferences(nums []int) []int {
	total := 0
	ps := make([]int, len(nums))
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		ps[i] = total
	}
	for i := 0; i < len(nums); i++ {
		pre := (nums[i] * (i + 1)) - ps[i]
		pst := (total - ps[i]) - (nums[i] * (len(nums) - 1 - i))
		ans[i] = pre + pst
	}
	return ans
}
