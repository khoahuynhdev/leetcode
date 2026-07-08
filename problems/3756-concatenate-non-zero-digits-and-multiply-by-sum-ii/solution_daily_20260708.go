package main

// 3756. Concatenate Non-Zero Digits and Multiply by Sum II
//
// For each query [l, r] we need x*sum (mod 1e9+7), where x is the integer
// formed by concatenating the non-zero digits of s[l..r] and sum is the sum
// of those digits.
//
// Key insight: a non-zero digit at position j contributes
//     d_j * 10^(number of non-zero digits after j within the range).
// Letting nz[i] be the count of non-zero digits in s[0..i-1], that exponent
// is nz[r+1] - nz[j+1]. Factoring out the range-dependent part:
//     x = 10^(nz[r+1]) * sum_j ( d_j * inv10^(nz[j+1]) )
// The inner sum is a prefix sum (pref), so every query is answered in O(1)
// after an O(m) precomputation. The digit sum is a plain prefix sum (dsum).

func processQueries(s string, queries [][]int) []int {
	const mod = 1_000_000_007
	m := len(s)

	// Powers of 10 and their modular inverses, indexed by exponent.
	pow10 := make([]int64, m+1)
	invPow10 := make([]int64, m+1)
	pow10[0] = 1
	for i := 1; i <= m; i++ {
		pow10[i] = pow10[i-1] * 10 % mod
	}
	inv10 := modpow(10, mod-2, mod)
	invPow10[0] = 1
	for i := 1; i <= m; i++ {
		invPow10[i] = invPow10[i-1] * inv10 % mod
	}

	// Prefix arrays over s[0..i-1].
	nz := make([]int, m+1)     // count of non-zero digits
	pref := make([]int64, m+1) // sum of d_j * inv10^(nz[j+1]) for non-zero j
	dsum := make([]int64, m+1) // sum of non-zero digit values
	for i := 0; i < m; i++ {
		d := int64(s[i] - '0')
		nz[i+1] = nz[i]
		pref[i+1] = pref[i]
		dsum[i+1] = dsum[i]
		if d != 0 {
			nz[i+1]++
			pref[i+1] = (pref[i+1] + d*invPow10[nz[i+1]]) % mod
			dsum[i+1] += d
		}
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		if nz[r+1]-nz[l] == 0 {
			ans[qi] = 0
			continue
		}
		sumContrib := ((pref[r+1]-pref[l])%mod + mod) % mod
		x := pow10[nz[r+1]] * sumContrib % mod
		sum := (dsum[r+1] - dsum[l]) % mod
		ans[qi] = int(x * sum % mod)
	}
	return ans
}

// modpow computes base^exp mod m using fast exponentiation.
func modpow(base, exp, m int64) int64 {
	result := int64(1)
	base %= m
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % m
		}
		base = base * base % m
		exp >>= 1
	}
	return result
}
