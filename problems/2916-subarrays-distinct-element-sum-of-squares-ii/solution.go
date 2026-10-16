package main

// Approach: sweep the right endpoint r and let f(l) be the distinct count of
// nums[l..r]. Appending nums[r] (previous occurrence at p) increments f(l) for
// every l in (p, r]. Since (x+1)^2 = x^2 + 2x + 1, the sum of squares over that
// range grows by 2*sum(f over range) + length. A lazy segment tree supporting
// range add / range sum gives sum(f over range) in O(log n). We keep the running
// total of squares for the current r and add it to the answer at each step.
// Time: O(n log n), Space: O(n).

const mod = 1_000_000_007

type segTree struct {
	sum  []int64
	lazy []int64
}

func newSegTree(n int) *segTree {
	return &segTree{sum: make([]int64, 4*n), lazy: make([]int64, 4*n)}
}

func (t *segTree) push(node, l, r int) {
	if t.lazy[node] == 0 {
		return
	}
	m := (l + r) / 2
	v := t.lazy[node]
	t.sum[2*node] += v * int64(m-l+1)
	t.lazy[2*node] += v
	t.sum[2*node+1] += v * int64(r-m)
	t.lazy[2*node+1] += v
	t.lazy[node] = 0
}

// query returns the sum of f(l) for l in [ql, qr].
func (t *segTree) query(node, l, r, ql, qr int) int64 {
	if qr < l || r < ql {
		return 0
	}
	if ql <= l && r <= qr {
		return t.sum[node]
	}
	t.push(node, l, r)
	m := (l + r) / 2
	return t.query(2*node, l, m, ql, qr) + t.query(2*node+1, m+1, r, ql, qr)
}

// add increments f(l) by 1 for l in [ql, qr].
func (t *segTree) add(node, l, r, ql, qr int) {
	if qr < l || r < ql {
		return
	}
	if ql <= l && r <= qr {
		t.sum[node] += int64(r - l + 1)
		t.lazy[node]++
		return
	}
	t.push(node, l, r)
	m := (l + r) / 2
	t.add(2*node, l, m, ql, qr)
	t.add(2*node+1, m+1, r, ql, qr)
	t.sum[node] = t.sum[2*node] + t.sum[2*node+1]
}

func sumCounts(nums []int) int {
	n := len(nums)
	tree := newSegTree(n)
	last := make(map[int]int)

	var squares, ans int64 // squares = sum of f(l)^2 for current r
	for r, v := range nums {
		lo := 0
		if p, ok := last[v]; ok {
			lo = p + 1
		}
		last[v] = r

		s := tree.query(1, 0, n-1, lo, r)
		squares = (squares + 2*s + int64(r-lo+1)) % mod
		tree.add(1, 0, n-1, lo, r)

		ans = (ans + squares) % mod
	}
	return int(ans)
}
