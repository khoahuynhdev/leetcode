package main

// 3658. GCD of Odd and Even Sums
//
// Approach: derive closed forms for the two sums.
//   sumOdd  = 1 + 3 + ... + (2n-1) = n^2
//   sumEven = 2 + 4 + ... + 2n     = n(n+1)
// Then gcd(n^2, n(n+1)) = n * gcd(n, n+1) = n, since consecutive
// integers n and n+1 are coprime. So the answer is simply n.
//
// For clarity and to mirror the literal problem statement, this
// implementation computes both sums explicitly and applies the
// Euclidean algorithm; the result always equals n.

func gcdOfOddEvenSums(n int) int {
	sumOdd := n * n
	sumEven := n * (n + 1)
	return gcd(sumOdd, sumEven)
}

// gcd returns the greatest common divisor of a and b using the
// iterative Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
