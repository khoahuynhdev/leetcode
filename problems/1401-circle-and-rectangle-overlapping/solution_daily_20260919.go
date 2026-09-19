package main

// 1401. Circle and Rectangle Overlapping
//
// The circle and the axis-aligned rectangle overlap exactly when the point of
// the rectangle closest to the circle's center lies within radius of it.
// Because the rectangle is [x1,x2] x [y1,y2], the closest point is found by
// clamping the center's coordinates into each interval independently. Comparing
// squared distances keeps everything in exact integer arithmetic.
//
// Time: O(1), Space: O(1).

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	// Closest point of the rectangle to the circle's center.
	cx := clampInt(xCenter, x1, x2)
	cy := clampInt(yCenter, y1, y2)

	dx := xCenter - cx
	dy := yCenter - cy

	// "<=" because sharing a single boundary point counts as overlapping.
	return dx*dx+dy*dy <= radius*radius
}

// clampInt returns v restricted to the closed interval [lo, hi].
func clampInt(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
