package main

// 1621. Number of Sets of K Non-Overlapping Line Segments
//
// A set of k non-overlapping segments is a chain of indices
//
//	0 <= a1 < b1 <= a2 < b2 <= ... <= ak < bk <= n-1
//
// Shifting the i-th segment (1-indexed) right by i-1 turns every "<=" into "<",
// so the 2k shifted values are strictly increasing and live in [0, n+k-2], a
// range of n+k-1 values. The shift is reversible, so configurations correspond
// one-to-one with 2k-element subsets of that range:
//
//	answer = C(n+k-1, 2k) mod 1e9+7
//
// The binomial is built multiplicatively, with a single modular inverse (via
// Fermat's little theorem) applied at the end. O(k + log MOD) time, O(1) space.

const mod1621 = 1_000_000_007

func numberOfSets(n int, k int) int {
	return binomMod1621(n+k-1, 2*k)
}

// binomMod1621 returns C(n, r) modulo mod1621, and 0 when r is out of range.
func binomMod1621(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}

	num, den := 1, 1
	for i := 0; i < r; i++ {
		num = num * ((n - i) % mod1621) % mod1621
		den = den * (i + 1) % mod1621
	}
	return num * powMod1621(den, mod1621-2) % mod1621
}

// powMod1621 computes base^exp modulo mod1621 by fast exponentiation.
func powMod1621(base, exp int) int {
	base %= mod1621
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod1621
		}
		base = base * base % mod1621
		exp >>= 1
	}
	return result
}
