package main

import "math"

// maximumProduct returns the largest product obtainable from any three elements
// of nums.
//
// Approach: only five values can ever matter — the three largest and the two
// smallest. A winning triple has either zero negatives (take the three largest)
// or two negatives (the two most negative values multiply to the largest
// positive, paired with the single largest value). Triples with one or three
// negatives are negative, and when the whole array is negative the three
// largest are already the product closest to zero. So the answer is
//
//	max(max1*max2*max3, min1*min2*max1)
//
// Both candidates are found in a single O(n) pass with O(1) extra space, which
// avoids the O(n log n) sort the obvious solution needs.
func maximumProduct(nums []int) int {
	// Seeded outside the [-1000, 1000] input range so the first elements
	// cascade into every slot, including for all-negative arrays.
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt

	for _, v := range nums {
		switch {
		case v >= max1:
			max1, max2, max3 = v, max1, max2
		case v >= max2:
			max2, max3 = v, max2
		case v >= max3:
			max3 = v
		}

		switch {
		case v <= min1:
			min1, min2 = v, min1
		case v <= min2:
			min2 = v
		}
	}

	byLargest := max1 * max2 * max3
	byNegativePair := min1 * min2 * max1
	if byNegativePair > byLargest {
		return byNegativePair
	}
	return byLargest
}
