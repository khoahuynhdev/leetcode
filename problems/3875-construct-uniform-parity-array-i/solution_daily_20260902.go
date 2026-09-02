package main

// 3875. Construct Uniform Parity Array I
// https://leetcode.com/problems/construct-uniform-parity-array-i/
//
// Approach: pure parity reasoning. Only the parities of nums1 matter, and the
// move nums2[i] = nums1[i] - nums1[j] flips index i's parity exactly when
// nums1[j] is odd. Two cases cover every input:
//
//  1. nums1 is already uniform (all even or all odd): keep every element via
//     nums2[i] = nums1[i].
//  2. nums1 is mixed: pick any index k holding an odd value. Keep every odd
//     element, and for every even element take nums2[i] = nums1[i] - nums1[k],
//     which is even - odd = odd. Since nums1[i] is even and nums1[k] is odd,
//     i != k holds automatically, and nothing forbids reusing the same k for
//     every even index. The result is all odd.
//
// Both cases succeed, so the answer is unconditionally true. Rather than
// returning a bare constant, this builds the witness nums2 and checks it, so
// the reasoning above is exercised instead of merely asserted.
//
// Time O(n), space O(n). Deciding the answer alone needs only O(1) of each.
func canConstructUniformParityArray(nums1 []int) bool {
	return isUniformParity(buildUniformParityArray(nums1))
}

// buildUniformParityArray returns a legal nums2 whose elements share a single
// parity, following the construction described above.
func buildUniformParityArray(nums1 []int) []int {
	odd := -1
	for i, v := range nums1 {
		if v%2 != 0 {
			odd = i
			break
		}
	}

	nums2 := make([]int, len(nums1))
	for i, v := range nums1 {
		// Keep v when it is already odd, or when there is no odd element to
		// borrow from (in which case every element is even and keeping them
		// all yields a uniformly even array).
		if odd == -1 || v%2 != 0 {
			nums2[i] = v
			continue
		}
		// v is even and nums1[odd] is odd, so i != odd is guaranteed.
		nums2[i] = v - nums1[odd]
	}
	return nums2
}

// isUniformParity reports whether every element of nums shares one parity.
// An array of fewer than two elements is trivially uniform.
func isUniformParity(nums []int) bool {
	for _, v := range nums {
		if (v-nums[0])%2 != 0 {
			return false
		}
	}
	return true
}
