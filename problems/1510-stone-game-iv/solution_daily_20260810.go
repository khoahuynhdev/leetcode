package main

// 1510. Stone Game IV
//
// The game is impartial: both players have the same legal moves, so a position
// is fully described by the number of stones left. Let dp[i] mean "the player
// about to move with i stones wins". A position wins if and only if some legal
// move hands the opponent a losing position:
//
//	dp[i] = true  iff exists k >= 1 with k*k <= i and !dp[i-k*k]
//
// dp[0] is false because a player facing an empty pile cannot move and loses.
// Filling the table bottom-up costs O(n*sqrt(n)) time and O(n) space.
func winnerSquareGame(n int) bool {
	if n < 0 {
		return false
	}

	dp := make([]bool, n+1) // dp[0] == false: no move available, mover loses
	for i := 1; i <= n; i++ {
		for k := 1; k*k <= i; k++ {
			if !dp[i-k*k] {
				dp[i] = true // one winning move is enough
				break
			}
		}
	}

	return dp[n]
}
