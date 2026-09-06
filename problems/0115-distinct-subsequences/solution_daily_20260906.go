package main

// 0115. Distinct Subsequences
//
// Count the index selections of s that spell out t.
//
// dp[j] holds the number of distinct subsequences of the prefix of s processed
// so far that equal t[:j]. For each character of s there are two disjoint
// choices -- drop it (dp[j] unchanged) or spend it on t[j-1] when the bytes
// match (dp[j] += dp[j-1]) -- so the counts add.
//
// The inner loop runs downward so dp[j-1] still holds the previous row's value
// when it is read; sweeping upward would let one character of s match two
// positions of t and overcount.
//
// Time: O(len(s) * len(t)). Space: O(len(t)).
func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	if m > n {
		return 0
	}

	dp := make([]int, m+1)
	dp[0] = 1 // the empty target is matched exactly one way: take nothing

	for i := 0; i < n; i++ {
		for j := m; j >= 1; j-- {
			if s[i] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[m]
}
