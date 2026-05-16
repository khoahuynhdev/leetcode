package main

// findMin returns the minimum element of a sorted-then-rotated array that
// may contain duplicates. It uses binary search comparing nums[mid] to
// nums[right]: a strict inequality identifies which half holds the rotation
// point, while equality forces a safe one-step shrink (right--), which
// degrades the worst case to O(n) when many duplicates are present.
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] > nums[right]:
			left = mid + 1
		case nums[mid] < nums[right]:
			right = mid
		default:
			right--
		}
	}
	return nums[left]
}
