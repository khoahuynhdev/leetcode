package main

// 1464. Maximum Product of Two Elements in an Array
//
// Since 1 <= nums[i] <= 10^3, every factor (nums[i]-1) is non-negative, so the
// product is maximized by independently maximizing both factors. That means the
// answer always comes from the two largest elements (counted with multiplicity,
// because i != j constrains indices, not values).
//
// A single pass tracks the largest and second-largest values: when a new maximum
// appears, the old maximum demotes to second place. 0 is a safe initial sentinel
// because no element can be smaller than 1.
//
// Time: O(n), Space: O(1).
func maxProduct(nums []int) int {
	largest, second := 0, 0
	for _, v := range nums {
		switch {
		case v > largest:
			second = largest
			largest = v
		case v > second:
			second = v
		}
	}
	return (largest - 1) * (second - 1)
}
