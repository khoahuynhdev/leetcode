package main

import "sort"

// Approach: sort a copy of nums by (digit sum, value) to get each element's
// target index, then treat the mapping current index -> target index as a
// permutation. A cycle of length k needs k-1 swaps, so the answer is
// n - (number of cycles).
func minSwaps(nums []int) int {
	n := len(nums)

	type item struct {
		val, sum int
	}
	items := make([]item, n)
	for i, v := range nums {
		items[i] = item{v, digitSum(v)}
	}
	sort.Slice(items, func(a, b int) bool {
		if items[a].sum != items[b].sum {
			return items[a].sum < items[b].sum
		}
		return items[a].val < items[b].val
	})

	pos := make(map[int]int, n)
	for i, it := range items {
		pos[it.val] = i
	}

	visited := make([]bool, n)
	cycles := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		cycles++
		for j := i; !visited[j]; j = pos[nums[j]] {
			visited[j] = true
		}
	}
	return n - cycles
}

func digitSum(x int) int {
	s := 0
	for x > 0 {
		s += x % 10
		x /= 10
	}
	return s
}
