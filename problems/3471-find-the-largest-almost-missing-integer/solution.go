package main

// 3471. Find the Largest Almost Missing Integer
//
// A value is "almost missing" when the windows of size k covering its
// occurrences number exactly one. Index i is covered by
// min(i, n-k) - max(0, i-k+1) + 1 windows, which equals 1 only at i == 0 and
// i == n-1 once 1 < k < n; every interior index is covered at least twice.
// So only the two endpoints are candidates, and each qualifies only if its
// value occurs nowhere else (a second occurrence, including one at the other
// end, would pull in a second window).
//
// Two degenerate cases fall outside that rule:
//   - k == n: a single window covers everything, so every value in it counts.
//   - k == 1: each window is one element, so the window count equals the
//     occurrence count.
//
// Values are non-negative, so -1 doubles as the "not found" sentinel.
// Time O(n), space O(n).
func largestInteger(nums []int, k int) int {
	n := len(nums)

	freq := make(map[int]int, n)
	for _, v := range nums {
		freq[v]++
	}

	best := -1

	switch {
	case k == n:
		// One window covering the whole array: every value is almost missing.
		for _, v := range nums {
			if v > best {
				best = v
			}
		}
	case k == 1:
		// Window count == occurrence count.
		for v, c := range freq {
			if c == 1 && v > best {
				best = v
			}
		}
	default:
		// Only the endpoints sit in a single window, and only if unique.
		for _, i := range [2]int{0, n - 1} {
			if v := nums[i]; freq[v] == 1 && v > best {
				best = v
			}
		}
	}

	return best
}
