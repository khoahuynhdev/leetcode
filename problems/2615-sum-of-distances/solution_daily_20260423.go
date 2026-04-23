package main

// Group indices by value, then for each group use a running prefix sum over
// the (already sorted) index list so every answer is computed in O(1).
// For sorted indices g[0..k-1] at position p:
//   leftSum  = p * g[p] - prefixSum
//   rightSum = (totalSum - prefixSum - g[p]) - (k - 1 - p) * g[p]
// Overall O(n) time and O(n) space. Uses int64 because distances can exceed int32.

func distance(nums []int) []int64 {
	n := len(nums)
	arr := make([]int64, n)

	groups := make(map[int][]int)
	for i, v := range nums {
		groups[v] = append(groups[v], i)
	}

	for _, indices := range groups {
		k := len(indices)
		if k <= 1 {
			continue
		}

		var total int64
		for _, idx := range indices {
			total += int64(idx)
		}

		var prefix int64
		for p, idx := range indices {
			idx64 := int64(idx)
			leftSum := int64(p)*idx64 - prefix
			rightSum := (total - prefix - idx64) - int64(k-1-p)*idx64
			arr[idx] = leftSum + rightSum
			prefix += idx64
		}
	}

	return arr
}
