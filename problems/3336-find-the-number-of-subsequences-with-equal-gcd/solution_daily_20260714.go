package main

// Approach: dynamic programming over (gcd of seq1, gcd of seq2).
//
// Process the elements one at a time. A DP state dp[a][b] holds the number of
// ways to assign the elements seen so far to one of three buckets — seq1, seq2,
// or "unused" — such that the running GCD of seq1 is `a` and the running GCD of
// seq2 is `b`. We use the sentinel value 0 to mean "the subsequence is still
// empty", which works nicely because gcd(0, x) == x, so adding the first
// element to an empty subsequence simply sets its GCD to that element.
//
// For each new element `num`, every existing state can transition three ways:
//   - skip num          -> dp[a][b] stays as dp[a][b]
//   - put num in seq1    -> contributes to dp[gcd(a, num)][b]
//   - put num in seq2    -> contributes to dp[a][gcd(b, num)]
//
// The final answer is the sum over all g >= 1 of dp[g][g]: both subsequences
// are non-empty (GCD != 0) and their GCDs are equal.
const mod = 1_000_000_007

func subsequencePairCount(nums []int) int {
	maxV := 0
	for _, v := range nums {
		if v > maxV {
			maxV = v
		}
	}

	// dp[a][b]: a, b range over 0..maxV, with 0 meaning "empty subsequence".
	dp := make([][]int, maxV+1)
	for i := range dp {
		dp[i] = make([]int, maxV+1)
	}
	dp[0][0] = 1 // both subsequences start empty

	for _, num := range nums {
		ndp := make([][]int, maxV+1)
		for i := range ndp {
			ndp[i] = make([]int, maxV+1)
			copy(ndp[i], dp[i]) // "skip num" transition
		}
		for a := 0; a <= maxV; a++ {
			for b := 0; b <= maxV; b++ {
				ways := dp[a][b]
				if ways == 0 {
					continue
				}
				// put num in seq1
				na := gcd(a, num)
				ndp[na][b] = (ndp[na][b] + ways) % mod
				// put num in seq2
				nb := gcd(b, num)
				ndp[a][nb] = (ndp[a][nb] + ways) % mod
			}
		}
		dp = ndp
	}

	ans := 0
	for g := 1; g <= maxV; g++ {
		ans = (ans + dp[g][g]) % mod
	}
	return ans
}

// gcd returns the greatest common divisor of a and b.
// gcd(0, x) == x, which lets 0 act as the "empty subsequence" sentinel.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
