package main

// Approach: a valid string has the shape a...ab...b. It is invalid exactly
// when some 'b' is immediately followed by an 'a' (the substring "ba").
// Scan adjacent pairs once: O(n) time, O(1) space.
func checkString(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] == 'b' && s[i] == 'a' {
			return false
		}
	}
	return true
}
