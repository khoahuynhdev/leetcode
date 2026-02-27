package main

// Approach: Count zeros z. We need m operations where m*k total flips can be
// assigned so zeros get odd flip counts and ones get even flip counts.
// Constraints: m*k ≡ z (mod 2), m*k >= z, and a capacity constraint based on
// m's parity. Try both even and odd m, return the minimum valid value.
func minOperations(s string, k int) int {
	n := len(s)
	z := 0
	for _, c := range s {
		if c == '0' {
			z++
		}
	}

	if z == 0 {
		return 0
	}

	// k even and z odd: m*k is always even but needs to be odd. Impossible.
	if k%2 == 0 && z%2 != 0 {
		return -1
	}

	// k == n: every operation flips all positions, so all share same flip parity.
	if k == n {
		if z == n {
			return 1
		}
		return -1
	}

	d := n - k // d >= 1 since k < n
	best := -1

	for parity := 0; parity <= 1; parity++ {
		// When k is odd, m*k has same parity as m, so m must match z's parity.
		if k%2 == 1 && parity != z%2 {
			continue
		}

		// Lower bound from m*k >= z
		lb1 := ceilDiv(z, k)

		// Lower bound from capacity constraint
		var lb3 int
		if parity == 0 { // m even: m*d >= z
			lb3 = ceilDiv(z, d)
		} else { // m odd: m*d >= n-z
			if n-z > 0 {
				lb3 = ceilDiv(n-z, d)
			}
		}

		m := lb1
		if lb3 > m {
			m = lb3
		}

		// Adjust to required parity
		if m%2 != parity {
			m++
		}

		if best == -1 || m < best {
			best = m
		}
	}

	return best
}

func ceilDiv(a, b int) int {
	return (a + b - 1) / b
}
