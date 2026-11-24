package main

// Approach: enumerate every starting index and extend the subarray to the
// right while maintaining its running product, GCD, and LCM. Because every
// value is in [1, 10], the LCM of any subarray is at most lcm(1..10) = 2520
// and the GCD is at most 10, so lcm*gcd <= 25200. The product never shrinks
// (all values are >= 1), so once it exceeds that bound, no longer subarray
// from this start can be product equivalent and we stop extending.
// Time: O(n^2 * log V), Space: O(1).
func maxLength(nums []int) int {
	const limit = 2520 * 10
	best := 0
	for i := range nums {
		prod, g, l := 1, 0, 1
		for j := i; j < len(nums); j++ {
			prod *= nums[j]
			if prod > limit {
				break
			}
			g = gcd(g, nums[j])
			l = l / gcd(l, nums[j]) * nums[j]
			if prod == l*g && j-i+1 > best {
				best = j - i + 1
			}
		}
	}
	return best
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
