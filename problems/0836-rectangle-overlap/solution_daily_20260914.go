package main

// 0836. Rectangle Overlap
//
// An axis-aligned rectangle is the Cartesian product of an X interval and a Y
// interval, so the 2D problem separates into the same 1D question asked twice:
// do two intervals share more than a single point?
//
// Two intervals [a1, a2] and [b1, b2] fail to overlap exactly when one ends at
// or before the other starts (a2 <= b1 || b2 <= a1). Negating that gives the
// overlap test a1 < b2 && b1 < a2. The strict inequalities are what make edges
// and corners that merely touch count as "no overlap", since those produce a
// zero-length intersection.
//
// The rectangles overlap iff both projections overlap. O(1) time, O(1) space.
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// X projections: [rec1[0], rec1[2]] and [rec2[0], rec2[2]].
	// Y projections: [rec1[1], rec1[3]] and [rec2[1], rec2[3]].
	return rec1[0] < rec2[2] && rec2[0] < rec1[2] &&
		rec1[1] < rec2[3] && rec2[1] < rec1[3]
}
