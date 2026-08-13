// 2213. Longest Substring of One Repeating Character
//
// A segment tree over the positions of s. Each node summarises its range with
// three numbers: the length of the run of equal characters at the left end
// (pref), the same at the right end (suf), and the longest run anywhere inside
// (best). Two children merge in O(1): the only run in the parent that lives in
// neither child is the one straddling the split point, and it exists exactly
// when the characters on either side of the split are equal. Each query is a
// point update costing O(log n), and the answer is the root's best.

package main

// repeatNode summarises one segment-tree range.
type repeatNode struct {
	pref int // longest run of equal characters starting at the range's left end
	suf  int // longest run of equal characters ending at the range's right end
	best int // longest run of equal characters anywhere in the range
}

// repeatSegTree holds the tree plus a mutable copy of the string, which the
// merge step consults to compare the characters adjacent to a split point.
type repeatSegTree struct {
	s    []byte
	tree []repeatNode
	n    int
}

func newRepeatSegTree(s string) *repeatSegTree {
	t := &repeatSegTree{s: []byte(s), n: len(s)}
	t.tree = make([]repeatNode, 4*t.n)
	t.build(1, 0, t.n-1)
	return t
}

func (t *repeatSegTree) build(idx, lo, hi int) {
	if lo == hi {
		// A single character is a run of length one, and always will be.
		t.tree[idx] = repeatNode{pref: 1, suf: 1, best: 1}
		return
	}
	mid := (lo + hi) / 2
	t.build(2*idx, lo, mid)
	t.build(2*idx+1, mid+1, hi)
	t.tree[idx] = t.merge(idx, lo, mid, hi)
}

// merge recomputes the node at idx from its children, which cover [lo, mid] and
// [mid+1, hi].
func (t *repeatSegTree) merge(idx, lo, mid, hi int) repeatNode {
	left, right := t.tree[2*idx], t.tree[2*idx+1]
	lenL, lenR := mid-lo+1, hi-mid

	res := repeatNode{
		pref: left.pref,
		suf:  right.suf,
		best: max(left.best, right.best),
	}
	if t.s[mid] == t.s[mid+1] {
		// A run crosses the boundary.
		res.best = max(res.best, left.suf+right.pref)
		// The prefix reaches past the boundary only if the whole left child is
		// a single run; symmetrically for the suffix and the right child.
		if left.pref == lenL {
			res.pref = lenL + right.pref
		}
		if right.suf == lenR {
			res.suf = lenR + left.suf
		}
	}
	return res
}

// update re-merges the path from the leaf at pos back up to idx. The character
// at pos must already have been written into t.s.
func (t *repeatSegTree) update(idx, lo, hi, pos int) {
	if lo == hi {
		// Leaves are unaffected by a character change: still a run of length 1.
		return
	}
	mid := (lo + hi) / 2
	if pos <= mid {
		t.update(2*idx, lo, mid, pos)
	} else {
		t.update(2*idx+1, mid+1, hi, pos)
	}
	t.tree[idx] = t.merge(idx, lo, mid, hi)
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	if len(s) == 0 || len(queryIndices) == 0 {
		return []int{}
	}

	t := newRepeatSegTree(s)
	lengths := make([]int, len(queryIndices))
	for i, pos := range queryIndices {
		t.s[pos] = queryCharacters[i]
		t.update(1, 0, t.n-1, pos)
		lengths[i] = t.tree[1].best
	}
	return lengths
}
