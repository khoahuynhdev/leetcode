package solution

import (
	"sort"
)

// For this one we will need sliding window technique
// and intuition:
// assume we need to find the target A
// and the frequency of the target is F
// the length of F is define is L, R where len(F) = R - L + 1
// the work we need to get done is equal W
// W = (A * F) - sum(R - L + 1)
// We will adjust the W (increase R, increase/decrease L) until W <= K and len(F) is maximum

func maxFrequency(nums []int, k int) int {
	// O(nlogn)
	sort.Ints(nums)
	left, cur := 0, 0
	for i := 0; i < len(nums); i++ {
		target := nums[i]
		cur += target
		if (i-left+1)*target-cur > k {
			cur -= nums[left]
			left++
		}
	}
	return len(nums) - left
}
