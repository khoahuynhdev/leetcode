package main

// 3739. Count Subarrays With Majority Element II
//
// Approach:
//   Map each element to +1 if it equals target and -1 otherwise. For any
//   subarray, target is the strict majority iff the transformed sum is > 0.
//   Using prefix sums P (P[0]=0), subarray (l..r) has positive sum iff
//   P[r+1] - P[l] > 0, i.e. P[l] < P[r+1]. So the answer is the number of
//   ordered index pairs (i, j) with i < j and P[i] < P[j].
//
//   We sweep the prefix sums left to right, and for each new value count how
//   many earlier prefix values are strictly smaller using a Fenwick tree
//   (Binary Indexed Tree). Prefix values live in [-n, n], shifted into
//   [1, 2n+1] for 1-indexed Fenwick storage.
//
// Time:  O(n log n)
// Space: O(n)

func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)
	// Fenwick tree over shifted prefix values in [1, 2n+1].
	size := 2*n + 1
	tree := make([]int, size+1)

	add := func(i int) {
		for ; i <= size; i += i & (-i) {
			tree[i]++
		}
	}
	// query returns how many inserted values have shifted index <= i.
	query := func(i int) int {
		s := 0
		for ; i > 0; i -= i & (-i) {
			s += tree[i]
		}
		return s
	}

	// Shift a prefix value p in [-n, n] to a 1-indexed slot in [1, 2n+1].
	idx := func(p int) int { return p + n + 1 }

	var ans int64
	prefix := 0
	add(idx(prefix)) // insert sentinel P[0] = 0

	for _, v := range nums {
		if v == target {
			prefix++
		} else {
			prefix--
		}
		// Count earlier prefix values strictly less than the current one.
		ans += int64(query(idx(prefix) - 1))
		add(idx(prefix))
	}
	return ans
}
