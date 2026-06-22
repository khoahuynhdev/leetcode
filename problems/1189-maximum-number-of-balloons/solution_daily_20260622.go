package main

// Maximum Number of Balloons
//
// Approach: frequency counting. Count every letter in text, then for each of the
// five letters that make up "balloon" (b, a, l, o, n) compute how many complete
// copies it can support — dividing by 2 for the doubled letters l and o. The
// answer is the minimum across those five, since the scarcest required letter is
// the bottleneck.
//
// Time:  O(n) for a single pass over text.
// Space: O(1) using a fixed 26-element counter.
func maxNumberOfBalloons(text string) int {
	var count [26]int
	for _, c := range text {
		count[c-'a']++
	}

	// Letters needed for one "balloon": b, a, n once each; l, o twice each.
	candidates := []int{
		count['b'-'a'],
		count['a'-'a'],
		count['l'-'a'] / 2,
		count['o'-'a'] / 2,
		count['n'-'a'],
	}

	best := candidates[0]
	for _, c := range candidates[1:] {
		if c < best {
			best = c
		}
	}
	return best
}
