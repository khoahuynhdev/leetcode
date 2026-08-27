package main

// 3720. Lexicographically Smallest Permutation Greater Than Target
//
// Only the letter counts of s matter, never its order. Any answer must agree
// with target on some prefix [0, i) and exceed it at index i, after which the
// suffix is free. A candidate that breaks later is always smaller (it still
// matches target where the earlier candidate already went above it), so we scan
// the breaking index downward from the longest prefix of target that our counts
// can actually spell (capped at n-1, since matching all n would only tie) and
// return the first candidate we can build: the smallest available letter above
// target[i], followed by every leftover letter in ascending order.
//
// Time: O(n * 26), Space: O(1) beyond the output.
func smallestPermutation(s string, target string) string {
	n := len(s)

	var counts [26]int
	for i := 0; i < n; i++ {
		counts[s[i]-'a']++
	}

	// Longest prefix of target spellable from s; no breaking index may exceed it.
	maxPrefix := 0
	remaining := counts
	for maxPrefix < n && remaining[target[maxPrefix]-'a'] > 0 {
		remaining[target[maxPrefix]-'a']--
		maxPrefix++
	}

	// Cap at n-1: breaking at n would reproduce target, which is not strictly greater.
	start := maxPrefix
	if start > n-1 {
		start = n - 1
	}

	// avail holds the multiset left after consuming target[:i].
	avail := counts
	for k := 0; k < start; k++ {
		avail[target[k]-'a']--
	}

	for i := start; i >= 0; i-- {
		for c := int(target[i]-'a') + 1; c < 26; c++ {
			if avail[c] == 0 {
				continue
			}
			avail[c]--
			out := make([]byte, 0, n)
			out = append(out, target[:i]...)
			out = append(out, byte('a'+c))
			for d := 0; d < 26; d++ {
				for j := 0; j < avail[d]; j++ {
					out = append(out, byte('a'+d))
				}
			}
			return string(out)
		}
		if i > 0 {
			// Moving to i-1 frees the letter that was pinned at position i-1.
			avail[target[i-1]-'a']++
		}
	}

	return ""
}
