package main

// Two-pointer sweep over two non-increasing arrays.
// Advance i when nums1[i] is too large for the current nums2[j] (no later j can
// rescue it, since nums2 only decreases). Advance j when the pair is valid to
// try to widen j-i. Skip i>j by nudging j forward. O(n+m) time, O(1) space.
func maxDistanceDaily20260419(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	ans, i, j := 0, 0, 0
	for i < n && j < m {
		switch {
		case i > j:
			j++
		case nums1[i] > nums2[j]:
			i++
		default:
			if j-i > ans {
				ans = j - i
			}
			j++
		}
	}
	return ans
}
