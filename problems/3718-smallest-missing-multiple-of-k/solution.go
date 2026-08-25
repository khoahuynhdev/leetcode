package main

// Store every value of nums in a hash set, then walk the multiples of k in
// increasing order (k, 2k, 3k, ...) and return the first one the set does not
// contain. Generating candidates smallest-first means the first miss is the
// answer. The loop always terminates: nums[i] <= 100, so any multiple above
// 100 is necessarily missing, bounding the scan to 100/k + 1 iterations.
func smallestMissingMultiple(nums []int, k int) int {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		seen[n] = struct{}{}
	}

	for multiple := k; ; multiple += k {
		if _, ok := seen[multiple]; !ok {
			return multiple
		}
	}
}
