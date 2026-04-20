package main

// maxDistance returns the maximum distance between two houses with different
// colors. The key observation is that at least one endpoint of the optimal
// pair is either the first house (index 0) or the last house (index n-1).
// So we perform two linear sweeps: one anchored at the left end scanning
// rightward from the far right, and one anchored at the right end scanning
// leftward from index 0. The answer is the larger of the two distances.
// Runs in O(n) time and O(1) extra space.
func maxDistance(colors []int) int {
	n := len(colors)
	best := 0

	// Anchor left end at index 0: find the furthest j with a different color.
	for j := n - 1; j > 0; j-- {
		if colors[j] != colors[0] {
			if j > best {
				best = j
			}
			break
		}
	}

	// Anchor right end at index n-1: find the earliest i with a different color.
	for i := 0; i < n-1; i++ {
		if colors[i] != colors[n-1] {
			if d := n - 1 - i; d > best {
				best = d
			}
			break
		}
	}

	return best
}
