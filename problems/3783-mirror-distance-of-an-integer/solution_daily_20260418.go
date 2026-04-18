package main

// mirrorDistance returns abs(n - reverse(n)) where reverse(n) is the integer
// formed by reversing the decimal digits of n. We compute the reverse by
// repeatedly peeling off the last digit with n%10 and shifting it onto an
// accumulator with rev = rev*10 + d. Trailing zeros are absorbed naturally
// because they are multiplied into rev while rev is still 0.
func mirrorDistance(n int) int {
	rev := 0
	for x := n; x > 0; x /= 10 {
		rev = rev*10 + x%10
	}
	diff := n - rev
	if diff < 0 {
		diff = -diff
	}
	return diff
}
