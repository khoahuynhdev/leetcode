package solution

import (
	"sort"
)

// solution O(nlogn) time
// solution O(1) space
// counting the total work to reduce the largest target to smallest target
// very similar to previous problem

func reductionOperations(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			continue
		}
		target := nums[i]
		if target < nums[i+1] {
			ans += (len(nums) - i - 1) * 1
		}
	}
	return ans
}
