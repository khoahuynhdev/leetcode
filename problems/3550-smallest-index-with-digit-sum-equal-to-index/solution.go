package main

// Approach: scan indices left to right, computing the digit sum of each
// element; the first index whose digit sum equals the index is the smallest.
// Time: O(n * d), Space: O(1).

func smallestIndex(nums []int) int {
	for i, num := range nums {
		if digitSum(num) == i {
			return i
		}
	}
	return -1
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
