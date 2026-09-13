package main

// Approach: pair-voting on translation offsets.
//
// A 1 at img1[i][j] can only land on a 1 at img2[k][l] under exactly one
// translation: the offset (k-i, l-j). So every cross pair of ones votes for a
// single offset, and the number of votes an offset receives is precisely the
// overlap produced by that translation. Collect the coordinates of the ones in
// each image, tally offsets in a map, and return the largest tally.
//
// Ones that would slide off the grid simply have no partner to pair with, so
// the "bits translated outside the borders are erased" rule needs no explicit
// bounds handling.
//
// Time: O(m1*m2) where m1, m2 are the counts of ones (O(n^4) worst case).
// Space: O(n^2) — at most (2n-1)^2 distinct offsets.
func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	if n == 0 {
		return 0
	}

	ones1 := onesOf(img1)
	ones2 := onesOf(img2)
	if len(ones1) == 0 || len(ones2) == 0 {
		return 0
	}

	// Offsets range over [-(n-1), n-1] on both axes, so bias by n before
	// packing into a single int key to keep negative and positive shifts apart.
	counts := make(map[int]int)
	best := 0
	for _, a := range ones1 {
		for _, b := range ones2 {
			key := (b[0]-a[0]+n)*(2*n) + (b[1] - a[1] + n)
			counts[key]++
			if counts[key] > best {
				best = counts[key]
			}
		}
	}
	return best
}

// onesOf returns the [row, col] coordinates of every 1 in the matrix.
func onesOf(img [][]int) [][2]int {
	ones := make([][2]int, 0, len(img)*len(img))
	for i, row := range img {
		for j, v := range row {
			if v == 1 {
				ones = append(ones, [2]int{i, j})
			}
		}
	}
	return ones
}
