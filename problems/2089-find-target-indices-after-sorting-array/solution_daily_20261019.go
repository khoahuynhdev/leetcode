package main

// Approach: counting, no sort needed.
// After sorting, every copy of target forms one contiguous block that starts
// at index `less` (the count of elements smaller than target) and has length
// `equal` (the count of elements equal to target). One pass computes both,
// giving O(n) time and O(1) extra space.
func targetIndices(nums []int, target int) []int {
	less, equal := 0, 0
	for _, v := range nums {
		if v < target {
			less++
		} else if v == target {
			equal++
		}
	}

	result := make([]int, 0, equal)
	for i := 0; i < equal; i++ {
		result = append(result, less+i)
	}
	return result
}
