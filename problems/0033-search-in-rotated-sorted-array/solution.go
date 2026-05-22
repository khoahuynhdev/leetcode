package main

// search finds target in a possibly-rotated ascending array of distinct values
// in O(log n) time using a modified binary search.
//
// At each step exactly one half of the window [lo, hi] is fully sorted (the
// other half holds the rotation seam). We detect the sorted half with
// nums[lo] <= nums[mid], check whether target lies inside that sorted range,
// and discard the half that cannot contain it.
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] {
			// Left half [lo..mid] is sorted.
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			// Right half [mid..hi] is sorted.
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
