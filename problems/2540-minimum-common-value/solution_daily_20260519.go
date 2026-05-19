package main

// getCommon returns the minimum integer that appears in both sorted, non-decreasing
// arrays nums1 and nums2, or -1 if no such integer exists.
//
// Approach: two-pointer sweep. Both arrays are sorted, so the first time the two
// pointers land on equal values that value is necessarily the smallest common one,
// since every value we previously stepped over on each side was strictly smaller.
// Whenever the pointers disagree, the side with the smaller value cannot match
// anything that remains on the other side, so we advance it.
//
// Time:  O(n + m)
// Space: O(1)
func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		switch {
		case nums1[i] == nums2[j]:
			return nums1[i]
		case nums1[i] < nums2[j]:
			i++
		default:
			j++
		}
	}
	return -1
}
