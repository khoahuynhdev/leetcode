package main

// findMin returns the minimum element of a sorted array that has been rotated
// between 1 and n times. We use binary search, comparing nums[mid] to nums[right]
// to decide which half of the current window contains the pivot (minimum):
//   - If nums[mid] > nums[right], the minimum is strictly to the right of mid.
//   - Otherwise, the right half from mid onward is sorted, so the minimum is
//     at mid or to its left.
// Runs in O(log n) time and O(1) space.
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
