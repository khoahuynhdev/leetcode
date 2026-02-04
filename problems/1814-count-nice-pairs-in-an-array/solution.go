package solution

// nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
// nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
// https://leetcode.com/problems/count-nice-pairs-in-an-array/editorial/?envType=daily-question&envId=2023-11-21
func reverse(n int) int {
	ans := 0
	tmp := n
	for tmp > 0 {
		num := tmp % 10
		ans = (ans * 10) + num
		tmp = tmp / 10
	}
	return ans
}

func countNicePairs(nums []int) int {
	fmap := make(map[int]int)
	f := make([]int, len(nums))
	ans := 0
	for i := 0; i < len(nums); i++ {
		f[i] = nums[i] - reverse(nums[i])
	}
	for _, v := range f {
		fmap[v]++
	}

	for _, v := range fmap {
		ans = (ans + v*(v-1)/2) % 1000000007
	}
	return ans
}
