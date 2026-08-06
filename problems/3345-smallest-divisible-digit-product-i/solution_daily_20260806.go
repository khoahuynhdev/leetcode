package main

// 3345. Smallest Divisible Digit Product I
//
// Approach: enumeration with a constant bound.
//
// Scan candidates upward from n and return the first whose digit product is
// divisible by t. The scan is guaranteed to terminate within 10 iterations:
// the next multiple of ten is at most 9 steps away, it ends in the digit 0, so
// its digit product is 0, and 0 % t == 0 for every t >= 1. Testing candidates
// in strictly increasing order makes the first hit the smallest one.

// digitProduct returns the product of the decimal digits of x.
// x is assumed to be positive.
func digitProduct(x int) int {
	product := 1
	for x > 0 {
		product *= x % 10
		x /= 10
	}
	return product
}

// smallestNumber returns the smallest number >= n whose digit product is
// divisible by t.
func smallestNumber(n int, t int) int {
	for x := n; ; x++ {
		if digitProduct(x)%t == 0 {
			return x
		}
	}
}
