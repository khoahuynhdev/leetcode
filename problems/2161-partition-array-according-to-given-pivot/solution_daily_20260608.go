package main

// Problem 2161: Partition Array According to Given Pivot
//
// Approach: stable three-way partition around pivot using two write pointers.
// We scan nums left-to-right placing every value < pivot into the front of the
// result (preserving their original relative order), then scan right-to-left
// placing every value > pivot into the back of the result (also preserving
// order). The remaining middle slots are exactly the elements equal to pivot,
// so we fill them with pivot. Runs in O(n) time with a single result slice.
func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	result := make([]int, n)

	left := 0
	for _, v := range nums {
		if v < pivot {
			result[left] = v
			left++
		}
	}

	right := n - 1
	for i := n - 1; i >= 0; i-- {
		if nums[i] > pivot {
			result[right] = nums[i]
			right--
		}
	}

	// Every slot in [left, right] is an element equal to the pivot.
	for i := left; i <= right; i++ {
		result[i] = pivot
	}

	return result
}
