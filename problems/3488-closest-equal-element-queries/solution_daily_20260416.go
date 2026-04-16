package main

// Approach: bucket indices by value, then for every index precompute its
// circular next-same and previous-same neighbors. For each query the answer is
// the minimum of the forward and backward circular distances to those two
// neighbors, computed in O(1). If the bucket for nums[q] has size 1 the answer
// is -1.
//
// Time:  O(n + q)
// Space: O(n)
func solveQueries(nums []int, queries []int) []int {
	n := len(nums)

	positions := make(map[int][]int, n)
	for i, v := range nums {
		positions[v] = append(positions[v], i)
	}

	nextSame := make([]int, n)
	prevSame := make([]int, n)
	for i := 0; i < n; i++ {
		nextSame[i] = -1
		prevSame[i] = -1
	}

	for _, idxList := range positions {
		k := len(idxList)
		if k < 2 {
			continue
		}
		for m, idx := range idxList {
			nextSame[idx] = idxList[(m+1)%k]
			prevSame[idx] = idxList[(m-1+k)%k]
		}
	}

	answer := make([]int, len(queries))
	for i, q := range queries {
		if nextSame[q] == -1 {
			answer[i] = -1
			continue
		}
		forward := (nextSame[q] - q + n) % n
		backward := (q - prevSame[q] + n) % n
		if forward < backward {
			answer[i] = forward
		} else {
			answer[i] = backward
		}
	}
	return answer
}
