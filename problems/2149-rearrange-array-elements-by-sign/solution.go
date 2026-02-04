package solution


func rearrangeArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	ip, in := 0, 0
	for i := 0; i < n; i += 2 {
		for in < n && nums[in] > 0 {
			in++
		}
		for ip < n && nums[ip] < 0 {
			ip++
		}
		
		ans[i], ans[i+1] = nums[ip], nums[in]
		
		ip++
		in++
	}

	return ans
}
