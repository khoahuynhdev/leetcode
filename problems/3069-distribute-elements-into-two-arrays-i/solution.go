package main

// 3069. Distribute Elements Into Two Arrays I
//
// Straight simulation of the process described in the statement. The problem is
// 1-indexed, so operations 1 and 2 seed arr1 with nums[0] and arr2 with nums[1]
// in 0-indexed terms; the remaining operations start at index 2.
//
// Each step only inspects the most recently appended element of each array, and
// a Go slice gives that in O(1) via s[len(s)-1], so the whole distribution is a
// single pass. Both slices are preallocated with capacity len(nums) since
// together they hold exactly n elements.
//
// Time: O(n). Space: O(n), all of it the output.
func resultArray(nums []int) []int {
	if len(nums) < 2 {
		// Constraints guarantee n >= 3, but stay total rather than panicking.
		return append([]int(nil), nums...)
	}

	arr1 := make([]int, 0, len(nums))
	arr2 := make([]int, 0, len(nums))
	arr1 = append(arr1, nums[0])
	arr2 = append(arr2, nums[1])

	for _, v := range nums[2:] {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, v)
		} else {
			arr2 = append(arr2, v)
		}
	}

	return append(arr1, arr2...)
}
