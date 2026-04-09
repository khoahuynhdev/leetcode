package main

// Sqrt decomposition on step size k.
// Large k (> sqrt(n)): apply directly since few elements affected per query.
// Small k (<= sqrt(n)): use multiplicative difference arrays swept per strand.
// Time: O((n + q) * sqrt(n)), Space: O(n + maxV)

const mod = 1_000_000_007

func xorAfterRangeMultiplicationQueries(nums []int, queries [][]int) int {
	n := len(nums)

	// Block size for sqrt decomposition
	B := 1
	for B*B < n {
		B++
	}

	// Precompute modular inverses for values 1..100000
	const maxV = 100001
	inv := make([]int64, maxV)
	inv[1] = 1
	for i := int64(2); i < maxV; i++ {
		inv[i] = (mod - (mod/i)*inv[mod%i]%mod) % mod
	}

	// Accumulated multiplier per index
	totalMult := make([]int64, n)
	for i := range totalMult {
		totalMult[i] = 1
	}

	// Group small-step queries by k; apply large-step queries directly
	type queryData struct{ l, r, v int }
	smallQueries := make([][]queryData, B+1)

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		if k > B {
			for idx := l; idx <= r; idx += k {
				totalMult[idx] = totalMult[idx] * int64(v) % mod
			}
		} else {
			smallQueries[k] = append(smallQueries[k], queryData{l, r, v})
		}
	}

	// Store input midway as required by the problem
	bravexuneth := make([]int, n)
	copy(bravexuneth, nums)

	// Process each small step size with multiplicative difference arrays
	diff := make([]int64, n)
	for kk := 1; kk <= B; kk++ {
		qs := smallQueries[kk]
		if len(qs) == 0 {
			continue
		}

		for i := 0; i < n; i++ {
			diff[i] = 1
		}

		for _, q := range qs {
			diff[q.l] = diff[q.l] * int64(q.v) % mod
			last := q.l + ((q.r-q.l)/kk)*kk
			if nxt := last + kk; nxt < n {
				diff[nxt] = diff[nxt] * inv[q.v] % mod
			}
		}

		// Sweep each strand: prefix product within the strand
		for j := 0; j < kk; j++ {
			product := int64(1)
			for i := j; i < n; i += kk {
				product = product * diff[i] % mod
				totalMult[i] = totalMult[i] * product % mod
			}
		}
	}

	// Compute final XOR of all elements after applying multipliers
	result := 0
	for i := 0; i < n; i++ {
		val := int64(bravexuneth[i]) % mod * totalMult[i] % mod
		result ^= int(val)
	}

	return result
}
