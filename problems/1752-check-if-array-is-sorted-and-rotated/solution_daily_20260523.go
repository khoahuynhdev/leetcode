package main

// Approach: A sorted-then-rotated array, viewed circularly, has at most one
// "drop" where an element is strictly greater than its next neighbor. We scan
// the array once and count positions where nums[i] > nums[(i+1)%n]. Zero drops
// means it's already sorted (rotation by zero); exactly one drop means a single
// valid rotation seam. Two or more drops mean it cannot be a rotated sorted
// array. Using strict ">" keeps duplicates (equal adjacent values) valid.
func check(nums []int) bool {
	n := len(nums)
	drops := 0
	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			drops++
		}
	}
	return drops <= 1
}
