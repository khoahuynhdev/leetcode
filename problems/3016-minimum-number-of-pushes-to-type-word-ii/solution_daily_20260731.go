package main

import "sort"

// 3016. Minimum Number of Pushes to Type Word II
//
// Greedy + counting. The order of characters in `word` is irrelevant; only the
// frequency of each letter matters. Push costs become available in tiers of 8:
// the first slot on each of the 8 keys costs 1 push, the second slot on each key
// costs 2, and so on. So the cheapest assignment pairs the highest frequencies
// with the cheapest slots: sort the 26 counts descending and charge the letter at
// index i a cost of i/8 + 1 per occurrence.
//
// Time: O(n) — one pass to count plus a constant-size sort of 26 values.
// Space: O(1) — a fixed 26-element counter array.
func minimumPushes(word string) int {
	var counts [26]int
	for i := 0; i < len(word); i++ {
		counts[word[i]-'a']++
	}

	freqs := counts[:]
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))

	total := 0
	for i, f := range freqs {
		if f == 0 {
			break
		}
		total += f * (i/8 + 1)
	}
	return total
}
