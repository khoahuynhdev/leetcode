package main

// maxRotateFunction computes the maximum value of F(k) over all rotations
// using the recurrence F(k) = F(k-1) + sum(nums) - n*nums[n-k]. This avoids
// the O(n^2) simulation by reusing the previous rotation's value, giving
// an O(n) time, O(1) space solution.
func maxRotateFunction(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sum := 0
	f := 0
	for i, v := range nums {
		sum += v
		f += i * v
	}

	best := f
	for k := 1; k < n; k++ {
		f += sum - n*nums[n-k]
		if f > best {
			best = f
		}
	}
	return best
}
