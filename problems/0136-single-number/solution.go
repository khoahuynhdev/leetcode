package solution

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
