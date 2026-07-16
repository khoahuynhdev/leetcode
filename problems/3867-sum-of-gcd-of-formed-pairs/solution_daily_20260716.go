package main

import "sort"

// Approach: Simulate the three stages directly.
//  1. Build prefixGcd in one pass using a running prefix maximum:
//     prefixGcd[i] = gcd(nums[i], max(nums[0..i])).
//  2. Sort prefixGcd in non-decreasing order.
//  3. Pair the smallest unpaired with the largest unpaired using two pointers
//     (left at the front, right at the back) and sum gcd(left, right) for each
//     pair. The `left < right` guard leaves the middle element unpaired when n
//     is odd, exactly as required.
//
// The answer is accumulated in an int64 to avoid overflow, since values can be
// up to 1e9 across many pairs.

// gcd returns the greatest common divisor of a and b using the Euclidean
// algorithm. Assumes non-negative inputs (constraints guarantee nums[i] >= 1).
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func sumGcd(nums []int) int64 {
	n := len(nums)

	// Stage 1: build prefixGcd with a running maximum.
	prefixGcd := make([]int, n)
	mx := 0
	for i, v := range nums {
		if v > mx {
			mx = v
		}
		prefixGcd[i] = gcd(v, mx)
	}

	// Stage 2: sort in non-decreasing order.
	sort.Ints(prefixGcd)

	// Stage 3: two-pointer pairing from both ends.
	var sum int64
	for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
		sum += int64(gcd(prefixGcd[left], prefixGcd[right]))
	}
	return sum
}
