package main

// 3090. Maximum Length Substring With Two Occurrences
//
// Sliding window with a fixed 26-slot frequency array. We extend the window by
// one character at a time; the only letter that can violate the "at most two
// occurrences" rule is the one just added, so whenever its count reaches 3 we
// advance left (decrementing counts as characters leave) until it drops back to
// 2. left never moves backwards, so every character enters and leaves the window
// at most once: O(n) time, O(1) space.
func maximumLengthSubstring(s string) int {
	var count [26]int
	best, left := 0, 0

	for right := 0; right < len(s); right++ {
		c := s[right] - 'a'
		count[c]++

		for count[c] > 2 {
			count[s[left]-'a']--
			left++
		}

		if length := right - left + 1; length > best {
			best = length
		}
	}

	return best
}
