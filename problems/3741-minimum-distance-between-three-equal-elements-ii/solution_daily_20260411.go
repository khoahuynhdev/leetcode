package main

// Approach: Group indices by value, then slide a window of 3 consecutive
// indices over each group. The distance formula for sorted i < j < k
// simplifies to 2*(k-i), so we minimize the span of the outermost indices.
// Time: O(n), Space: O(n).
func minimumDistance(nums []int) int {
	indexMap := make(map[int][]int)
	for i, v := range nums {
		indexMap[v] = append(indexMap[v], i)
	}

	result := -1
	for _, indices := range indexMap {
		if len(indices) < 3 {
			continue
		}
		for i := 0; i+2 < len(indices); i++ {
			dist := 2 * (indices[i+2] - indices[i])
			if result == -1 || dist < result {
				result = dist
			}
		}
	}

	return result
}
