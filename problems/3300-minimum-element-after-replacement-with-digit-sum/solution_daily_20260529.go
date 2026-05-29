package main

// Approach: single linear pass. For each element, compute its decimal digit
// sum by repeated mod-10 / div-10, and track the minimum digit sum seen so
// far. O(n * d) time where d is the number of digits (<= 5 here), O(1) space.

func minElementDaily20260529(nums []int) int {
	best := -1
	for _, n := range nums {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		if best == -1 || sum < best {
			best = sum
		}
	}
	return best
}
