package main

// Segment tree approach: sweep right endpoint r, maintaining balance[l] =
// distinct_even(l,r) - distinct_odd(l,r) for all left endpoints l.
// When processing nums[r] with previous occurrence at prev:
//   range-add +1 (even) or -1 (odd) on [prev+1, r].
// Query leftmost l in [0,r] where balance[l] == 0.
// Segment tree stores min/max per node for pruning the zero-search.

type segNode struct {
	mn, mx int
	lazy   int
}

func longestBalancedSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	tree := make([]segNode, 4*n)

	pushDown := func(idx int) {
		if tree[idx].lazy != 0 {
			for _, c := range []int{2 * idx, 2*idx + 1} {
				tree[c].mn += tree[idx].lazy
				tree[c].mx += tree[idx].lazy
				tree[c].lazy += tree[idx].lazy
			}
			tree[idx].lazy = 0
		}
	}

	var update func(idx, lo, hi, l, r, val int)
	update = func(idx, lo, hi, l, r, val int) {
		if l > hi || r < lo {
			return
		}
		if l <= lo && hi <= r {
			tree[idx].mn += val
			tree[idx].mx += val
			tree[idx].lazy += val
			return
		}
		pushDown(idx)
		mid := (lo + hi) / 2
		update(2*idx, lo, mid, l, r, val)
		update(2*idx+1, mid+1, hi, l, r, val)
		tree[idx].mn = minInt(tree[2*idx].mn, tree[2*idx+1].mn)
		tree[idx].mx = maxInt(tree[2*idx].mx, tree[2*idx+1].mx)
	}

	var findLeftmostZero func(idx, lo, hi, ql, qr int) int
	findLeftmostZero = func(idx, lo, hi, ql, qr int) int {
		if lo > qr || hi < ql || tree[idx].mn > 0 || tree[idx].mx < 0 {
			return -1
		}
		if lo == hi {
			if tree[idx].mn == 0 {
				return lo
			}
			return -1
		}
		pushDown(idx)
		mid := (lo + hi) / 2
		res := findLeftmostZero(2*idx, lo, mid, ql, qr)
		if res != -1 {
			return res
		}
		return findLeftmostZero(2*idx+1, mid+1, hi, ql, qr)
	}

	ans := 0
	prev := make(map[int]int) // value -> last seen index

	for r := 0; r < n; r++ {
		v := nums[r]
		lo := 0
		if p, ok := prev[v]; ok {
			lo = p + 1
		}
		delta := 1
		if v%2 != 0 {
			delta = -1
		}
		update(1, 0, n-1, lo, r, delta)
		prev[v] = r

		leftmost := findLeftmostZero(1, 0, n-1, 0, r)
		if leftmost != -1 {
			if length := r - leftmost + 1; length > ans {
				ans = length
			}
		}
	}

	return ans
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
