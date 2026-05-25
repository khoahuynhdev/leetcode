package main

// Jump Game VII
//
// dp[i] is true when index i is reachable from index 0. Reversing the jump
// rule, index i is reachable iff s[i] == '0' and some reachable index lies in
// the window [i-maxJump, i-minJump]. A prefix sum of dp answers "is any
// reachable index in this window?" in O(1), giving an overall O(n) solution.
func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	dp := make([]bool, n)
	dp[0] = true

	// pre[i] = count of reachable indices among dp[0..i].
	pre := make([]int, n)
	pre[0] = 1

	for i := 1; i < n; i++ {
		if s[i] == '0' {
			lo := i - maxJump
			hi := i - minJump
			if lo < 0 {
				lo = 0
			}
			if hi >= 0 {
				left := 0
				if lo > 0 {
					left = pre[lo-1]
				}
				if pre[hi]-left > 0 {
					dp[i] = true
				}
			}
		}

		pre[i] = pre[i-1]
		if dp[i] {
			pre[i]++
		}
	}

	return dp[n-1]
}
