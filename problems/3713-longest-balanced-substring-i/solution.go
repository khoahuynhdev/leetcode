package main

// Enumerate all substrings with two nested loops, maintaining a frequency array.
// A substring is balanced when maxFreq * distinct == length, checked in O(1).
// Overall time: O(n^2), space: O(1).
func longestBalancedSubstring(s string) int {
	n := len(s)
	ans := 1 // every single character is balanced

	for i := 0; i < n; i++ {
		var freq [26]int
		distinct := 0
		maxFreq := 0

		for j := i; j < n; j++ {
			c := s[j] - 'a'
			freq[c]++
			if freq[c] == 1 {
				distinct++
			}
			if freq[c] > maxFreq {
				maxFreq = freq[c]
			}

			length := j - i + 1
			if maxFreq*distinct == length && length > ans {
				ans = length
			}
		}
	}

	return ans
}
