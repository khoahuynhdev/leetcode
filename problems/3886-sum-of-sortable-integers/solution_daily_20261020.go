package main

// Approach: enumerate every divisor k of n. Rotating a block never changes
// which values it holds, so k is sortable iff
//  1. every block is a rotation of a sorted sequence, i.e. it has at most one
//     cyclic "descent" (a position i with block[i] > block[i+1], wrapping the
//     last element around to the first), and
//  2. each block's max is <= the next block's min (blocks can't exchange values).
// n has at most 128 divisors for n <= 1e5, and each check is O(n).
func sumOfSortableIntegers(nums []int) int {
	n := len(nums)
	total := 0
	for k := 1; k <= n; k++ {
		if n%k == 0 && isSortable(nums, k) {
			total += k
		}
	}
	return total
}

func isSortable(nums []int, k int) bool {
	prevMax := 0
	for start := 0; start < len(nums); start += k {
		end := start + k - 1
		descents := 0
		lo, hi := nums[start], nums[start]
		for i := start; i <= end; i++ {
			next := i + 1
			if i == end {
				next = start
			}
			if nums[i] > nums[next] {
				descents++
			}
			lo = min(lo, nums[i])
			hi = max(hi, nums[i])
		}
		if descents > 1 || lo < prevMax {
			return false
		}
		prevMax = hi
	}
	return true
}
