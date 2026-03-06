package main

import "strings"

// checkOnesSegment checks if the binary string has at most one contiguous
// segment of ones. Since the string has no leading zeros, it always starts
// with '1'. A second segment of ones can only exist if "01" appears in s.
func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
