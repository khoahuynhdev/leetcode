package main

// Approach: Simulate the transformation by calculating new positions using modular arithmetic
// For each element, move the specified number of steps in the circular array and
// store the value at the destination position in the result array.

// Claude code implementation
// func constructTransformedArray(nums []int) []int {
//     n := len(nums)
//     result := make([]int, n)
//
//     for i := 0; i < n; i++ {
//         if nums[i] == 0 {
//             result[i] = nums[i]
//         } else if nums[i] > 0 {
//             // Move right: (current_index + steps) % n
//             newIndex := (i + nums[i]) % n
//             result[i] = nums[newIndex]
//         } else {
//             // Move left: handle negative steps with proper modular arithmetic
//             // (current_index + steps%n + n) % n ensures positive result
//             steps := nums[i] % n
//             newIndex := (i + steps + n) % n
//             result[i] = nums[newIndex]
//         }
//     }
//
//     return result
// }

// my implementation
// constructTransformedArray
func constructTransformedArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range nums {
		val := (i + nums[i]) % n
		if val >= 0 {
			ans[i] = nums[val]
		} else {
			ans[i] = nums[val+n]
		}
	}
	// 0 + 3 % 3 = 0 -> nums[0]
	// 1 + (-5)  % 3 = 0 <=> -1 =
	// neg_idx = idx - len
	// 2 + 1 % 3 = 0 -> nums[0]
	return ans
}
