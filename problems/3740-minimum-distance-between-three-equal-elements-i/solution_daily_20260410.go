package main

// Approach: Group indices by value using a hash map.
// Key insight: for i < j < k, distance = 2*(k-i) since the middle index cancels.
// Slide a window of 3 over each group's indices to find the minimum span.
func minimumDistance(nums []int) int {
	indices := make(map[int][]int)
	for i, v := range nums {
		indices[v] = append(indices[v], i)
	}

	res := -1
	for _, idx := range indices {
		if len(idx) < 3 {
			continue
		}
		for i := 0; i+2 < len(idx); i++ {
			d := 2 * (idx[i+2] - idx[i])
			if res == -1 || d < res {
				res = d
			}
		}
	}
	return res
}
