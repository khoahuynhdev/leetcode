package main

// Maximum Total Subarray Value I
//
// The value of any subarray is max(subarray) - min(subarray), which can never
// exceed max(nums) - min(nums). The full array attains that bound, and since we
// may choose the same subarray repeatedly, the best total is to pick that single
// best subarray all k times. Thus the answer is k * (max(nums) - min(nums)).
//
// We use int64 for the result because k (up to 1e5) times the difference (up to
// 1e9) can reach 1e14, which overflows 32-bit integers.
//
// Time: O(n), Space: O(1).
func maxTotalValue(nums []int, k int) int64 {
	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}
	return int64(k) * int64(mx-mn)
}
