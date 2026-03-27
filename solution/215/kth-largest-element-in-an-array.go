package solution

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findKthLargest(nums []int, k int) int {
	var c int = 1.e4
	vals := make([]int, 2*c+1)
	for _, v := range nums {
		vals[v+c]++
	}
	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] == 0 {
			continue
		}
		if k > 0 {
			m := minInt(k, vals[i])
			k -= m
			vals[i] -= m
		}
		if k == 0 {
			return i - c
		}

	}
	return k
}
