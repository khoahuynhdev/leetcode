package main

// Count Subarrays With Majority Element I
//
// Approach: map each element to +1 if it equals target and -1 otherwise.
// target is the strict majority of a subarray iff the transformed sum is > 0.
// Using prefix sums pre[], the answer is the number of pairs i < j with
// pre[i] < pre[j]. We scan left to right and, for each new prefix value,
// count how many earlier prefixes are strictly smaller using a Fenwick tree
// (Binary Indexed Tree) keyed by the prefix value shifted into [0, 2n].
//
// Time:  O(n log n)
// Space: O(n)

// fenwick is a 1-indexed Binary Indexed Tree storing counts.
type fenwick struct {
	tree []int
}

func newFenwick(n int) *fenwick {
	return &fenwick{tree: make([]int, n+1)}
}

// update adds 1 to position i (1-indexed).
func (f *fenwick) update(i int) {
	for ; i < len(f.tree); i += i & (-i) {
		f.tree[i]++
	}
}

// query returns the sum of counts in positions [1, i] (1-indexed).
func (f *fenwick) query(i int) int {
	sum := 0
	for ; i > 0; i -= i & (-i) {
		sum += f.tree[i]
	}
	return sum
}

func countSubarrays(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Prefix values range in [-n, n]; shift by n so they fall in [0, 2n].
	// Fenwick is 1-indexed, so the stored index is (value + n + 1) in [1, 2n+1].
	const baseOffset = 1
	bit := newFenwick(2*n + baseOffset)

	shift := func(value int) int { return value + n + baseOffset }

	answer := 0
	pre := 0
	bit.update(shift(pre)) // insert pre[0] = 0 before scanning.

	for _, v := range nums {
		if v == target {
			pre++
		} else {
			pre--
		}
		// Number of earlier prefixes strictly smaller than pre:
		// values <= pre-1, i.e. Fenwick indices up to shift(pre) - 1.
		answer += bit.query(shift(pre) - 1)
		bit.update(shift(pre))
	}

	return answer
}
