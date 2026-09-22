package main

// 3525. Find X Value of Array II
//
// A query (index, value, start, x) asks: after nums[index] = value, how many
// indices j in [start, n-1] satisfy
//
//	nums[start] * nums[start+1] * ... * nums[j] % k == x
//
// Since k <= 5, every element collapses to one of k residues, so a segment
// tree node only needs two pieces of information about its range:
//
//	prod    - product of the whole range modulo k
//	base[p] - how many prefixes of the range have product == p (mod k)
//
// Merging a left child L with a right child R shifts every prefix of R by
// L.prod, while the prefixes of L are unchanged:
//
//	prod    = L.prod * R.prod % k
//	base[v] = L.base[v] + sum of R.base[p] over p with (L.prod*p)%k == v
//
// A query walks the O(log n) nodes covering the suffix [start, n-1] from left
// to right, carrying the running residue r of everything already consumed, and
// adds the number of prefixes p inside the node with (r*p)%k == x.
//
// Build O(n*k), update O(k*log n), query O(k*log n).

// xSegTree is a segment tree over residues modulo k.
type xSegTree struct {
	k    int
	n    int
	prod []int32 // prod[node] = product of node's range mod k
	base []int32 // base[node*k+p] = prefixes of node's range with product p
}

func newXSegTree(nums []int, k int) *xSegTree {
	n := len(nums)
	t := &xSegTree{
		k:    k,
		n:    n,
		prod: make([]int32, 4*n),
		base: make([]int32, 4*n*k),
	}
	t.build(1, 0, n-1, nums)
	return t
}

// setLeaf stores a single element whose residue is v.
func (t *xSegTree) setLeaf(node int, v int32) {
	b := t.base[node*t.k : node*t.k+t.k]
	for i := range b {
		b[i] = 0
	}
	b[v] = 1
	t.prod[node] = v
}

// pull recomputes node from its children l and r.
func (t *xSegTree) pull(node, l, r int) {
	k := t.k
	lp := int(t.prod[l])
	t.prod[node] = int32(lp * int(t.prod[r]) % k)

	nb := t.base[node*k : node*k+k]
	lb := t.base[l*k : l*k+k]
	rb := t.base[r*k : r*k+k]
	copy(nb, lb)
	for p := 0; p < k; p++ {
		nb[lp*p%k] += rb[p]
	}
}

func (t *xSegTree) build(node, lo, hi int, nums []int) {
	if lo == hi {
		t.setLeaf(node, int32(nums[lo]%t.k))
		return
	}
	mid := (lo + hi) / 2
	t.build(2*node, lo, mid, nums)
	t.build(2*node+1, mid+1, hi, nums)
	t.pull(node, 2*node, 2*node+1)
}

func (t *xSegTree) update(node, lo, hi, pos, val int) {
	if lo == hi {
		t.setLeaf(node, int32(val%t.k))
		return
	}
	mid := (lo + hi) / 2
	if pos <= mid {
		t.update(2*node, lo, mid, pos, val)
	} else {
		t.update(2*node+1, mid+1, hi, pos, val)
	}
	t.pull(node, 2*node, 2*node+1)
}

// query counts the prefixes of nums[start..n-1] whose product is x mod k.
// r is the residue accumulated by the segments already visited; the updated
// residue is returned alongside the count so that segments stay ordered.
func (t *xSegTree) query(node, lo, hi, start, r, x int) (int, int) {
	if hi < start {
		return 0, r
	}
	k := t.k
	if lo >= start {
		cnt := 0
		b := t.base[node*k : node*k+k]
		for p := 0; p < k; p++ {
			if r*p%k == x {
				cnt += int(b[p])
			}
		}
		return cnt, r * int(t.prod[node]) % k
	}
	mid := (lo + hi) / 2
	left, r := t.query(2*node, lo, mid, start, r, x)
	right, r := t.query(2*node+1, mid+1, hi, start, r, x)
	return left + right, r
}

func resultArray(nums []int, k int, queries [][]int) []int {
	n := len(nums)
	tree := newXSegTree(nums, k)

	res := make([]int, len(queries))
	for i, q := range queries {
		index, value, start, x := q[0], q[1], q[2], q[3]
		tree.update(1, 0, n-1, index, value)
		cnt, _ := tree.query(1, 0, n-1, start, 1%k, x)
		res[i] = cnt
	}
	return res
}
