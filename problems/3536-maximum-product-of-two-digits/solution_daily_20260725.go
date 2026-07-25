package main

// 3536. Maximum Product of Two Digits
//
// Every digit lies in [0, 9], so all candidate factors are non-negative and the
// product a*b is non-decreasing in both a and b. That means the answer is always
// the largest digit times the second-largest digit (by position, so a repeated
// digit may be used twice) -- no pair loop and no sorting required.
//
// We peel the decimal digits off with %10 / /=10 and keep a running top-two in
// two variables. The "else if d > second" branch is what lets an equal digit
// occupy the second slot, which is exactly the n = 22 case.
//
// Time: O(log n) -- at most 10 iterations. Space: O(1).
func maxProduct(n int) int {
	best, second := 0, 0
	for ; n > 0; n /= 10 {
		d := n % 10
		switch {
		case d > best:
			best, second = d, best
		case d > second:
			second = d
		}
	}
	return best * second
}
