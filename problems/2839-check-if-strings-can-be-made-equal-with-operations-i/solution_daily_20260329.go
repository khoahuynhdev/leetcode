package main

// Approach: Characters at even indices (0,2) can only swap with each other,
// and characters at odd indices (1,3) can only swap with each other.
// Check that both groups have matching character multisets across s1 and s2.
func canBeEqual(s1 string, s2 string) bool {
	// Check even-indexed characters {0, 2}
	if !sameChars(s1[0], s1[2], s2[0], s2[2]) {
		return false
	}
	// Check odd-indexed characters {1, 3}
	return sameChars(s1[1], s1[3], s2[1], s2[3])
}

// sameChars checks if the multiset {a, b} equals the multiset {c, d}.
func sameChars(a, b, c, d byte) bool {
	return (a == c && b == d) || (a == d && b == c)
}
