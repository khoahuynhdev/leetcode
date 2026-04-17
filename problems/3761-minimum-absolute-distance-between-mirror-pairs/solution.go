package main

// Single left-to-right pass with a hash map keyed by reversed values.
// For each index j, lastSeen[nums[j]] (if present) is the most recent i
// with reverse(nums[i]) == nums[j], which minimizes j - i among all such i.
// After checking, record lastSeen[reverse(nums[j])] = j so future indices
// can pair against j. Time O(n * d), space O(n).
func minAbsoluteDistance(nums []int) int {
	lastSeen := make(map[int]int, len(nums))
	best := -1
	for j, v := range nums {
		if i, ok := lastSeen[v]; ok {
			d := j - i
			if best == -1 || d < best {
				best = d
			}
		}
		lastSeen[reverseDigits(v)] = j
	}
	return best
}

func reverseDigits(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return r
}
