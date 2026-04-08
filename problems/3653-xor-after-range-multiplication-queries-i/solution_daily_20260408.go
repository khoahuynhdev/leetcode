// Simulation approach: process each query by stepping through the range
// with stride k, multiplying elements by v mod 10^9+7, then XOR all elements.
package main

const mod = 1_000_000_007

func xorAfterRangeMultiplication(nums []int, queries [][]int) int {
	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		for idx := l; idx <= r; idx += k {
			nums[idx] = nums[idx] * v % mod
		}
	}

	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
