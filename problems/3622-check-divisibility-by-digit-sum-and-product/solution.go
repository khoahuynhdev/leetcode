package main

// 3622. Check Divisibility by Digit Sum and Product
//
// Extract the digits of n in a single pass using the repeated `% 10` / `/ 10`
// idiom, accumulating the digit sum (identity 0) and the digit product
// (identity 1) at the same time. The scratch variable `rest` is what shrinks,
// so the original n survives for the final divisibility check.
//
// The divisor can never be zero: n >= 1 guarantees a digit sum of at least 1,
// even when a 0 digit zeroes out the product.
//
// Time: O(log n) — one iteration per digit. Space: O(1).
func checkDivisibility(n int) bool {
	digitSum, digitProduct := 0, 1

	for rest := n; rest > 0; rest /= 10 {
		d := rest % 10
		digitSum += d
		digitProduct *= d
	}

	return n%(digitSum+digitProduct) == 0
}
