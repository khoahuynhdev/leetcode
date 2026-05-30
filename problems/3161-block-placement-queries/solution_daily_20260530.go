package main

// Approach: coordinate-compress every position that appears in any query
// (plus a sentinel at 0), then maintain two segment trees indexed by
// compressed position:
//
//   presence[i] = 1 iff an obstacle currently sits at positions[i]
//                (sum-aggregated; used to descend to predecessor / successor)
//   gap[i]      = positions[i] - positions[prevObstacle(i)] for each
//                 active obstacle (max-aggregated)
//
// Inserting x updates gap at x and at x's successor — exactly the two
// gaps that change. A type-2 query (x, sz) takes max(rangeMax over gap
// for indices in (0, idx(pred)]) and (x - pred), and checks against sz.

import "sort"

type segPresence struct {
	size int
	tree []int
}

func newSegPresence(n int) *segPresence {
	size := 1
	for size < n {
		size <<= 1
	}
	return &segPresence{size: size, tree: make([]int, 2*size)}
}

func (s *segPresence) update(i, v int) {
	i += s.size
	s.tree[i] = v
	for i > 1 {
		i >>= 1
		s.tree[i] = s.tree[2*i] + s.tree[2*i+1]
	}
}

// findPredecessor returns the largest index in [0, lim] whose value is 1,
// or -1 if none exists.
func (s *segPresence) findPredecessor(lim int) int {
	if lim < 0 {
		return -1
	}
	return s.findPredRec(1, 0, s.size-1, lim)
}

func (s *segPresence) findPredRec(node, l, r, lim int) int {
	if l > lim || s.tree[node] == 0 {
		return -1
	}
	if l == r {
		return l
	}
	mid := (l + r) >> 1
	if right := s.findPredRec(2*node+1, mid+1, r, lim); right != -1 {
		return right
	}
	return s.findPredRec(2*node, l, mid, lim)
}

// findSuccessor returns the smallest index in [lim, n-1] whose value is 1,
// or -1 if none exists.
func (s *segPresence) findSuccessor(lim int) int {
	if lim >= s.size {
		return -1
	}
	return s.findSuccRec(1, 0, s.size-1, lim)
}

func (s *segPresence) findSuccRec(node, l, r, lim int) int {
	if r < lim || s.tree[node] == 0 {
		return -1
	}
	if l == r {
		return l
	}
	mid := (l + r) >> 1
	if left := s.findSuccRec(2*node, l, mid, lim); left != -1 {
		return left
	}
	return s.findSuccRec(2*node+1, mid+1, r, lim)
}

type segMax struct {
	size int
	tree []int
}

func newSegMax(n int) *segMax {
	size := 1
	for size < n {
		size <<= 1
	}
	return &segMax{size: size, tree: make([]int, 2*size)}
}

func (s *segMax) update(i, v int) {
	i += s.size
	s.tree[i] = v
	for i > 1 {
		i >>= 1
		left, right := s.tree[2*i], s.tree[2*i+1]
		if left >= right {
			s.tree[i] = left
		} else {
			s.tree[i] = right
		}
	}
}

func (s *segMax) queryMax(l, r int) int {
	if l > r {
		return 0
	}
	res := 0
	l += s.size
	r += s.size + 1
	for l < r {
		if l&1 == 1 {
			if s.tree[l] > res {
				res = s.tree[l]
			}
			l++
		}
		if r&1 == 1 {
			r--
			if s.tree[r] > res {
				res = s.tree[r]
			}
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func getResults(queries [][]int) []bool {
	// Collect every distinct position we might need to address, plus the
	// sentinel at 0. We need x values from type-2 queries too, because
	// findPredecessor / queryMax operate on compressed indices.
	posSet := map[int]struct{}{0: {}}
	for _, q := range queries {
		posSet[q[1]] = struct{}{}
	}
	positions := make([]int, 0, len(posSet))
	for p := range posSet {
		positions = append(positions, p)
	}
	sort.Ints(positions)

	posIdx := make(map[int]int, len(positions))
	for i, p := range positions {
		posIdx[p] = i
	}
	n := len(positions)

	pres := newSegPresence(n)
	gaps := newSegMax(n)

	// Mark the sentinel at position 0 as an obstacle (its gap stays 0).
	pres.update(posIdx[0], 1)

	results := make([]bool, 0, len(queries))
	for _, q := range queries {
		switch q[0] {
		case 1:
			x := q[1]
			xi := posIdx[x]
			predIdx := pres.findPredecessor(xi - 1)
			predPos := positions[predIdx]
			succIdx := pres.findSuccessor(xi + 1)
			pres.update(xi, 1)
			gaps.update(xi, x-predPos)
			if succIdx != -1 {
				gaps.update(succIdx, positions[succIdx]-x)
			}
		case 2:
			x, sz := q[1], q[2]
			xi := posIdx[x]
			predIdx := pres.findPredecessor(xi)
			predPos := positions[predIdx]
			interior := gaps.queryMax(1, predIdx)
			trailing := x - predPos
			best := interior
			if trailing > best {
				best = trailing
			}
			results = append(results, best >= sz)
		}
	}
	return results
}
