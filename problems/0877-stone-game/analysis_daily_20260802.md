# 0877. Stone Game

[LeetCode Link](https://leetcode.com/problems/stone-game/)

Difficulty: Medium
Topics: Array, Math, Dynamic Programming, Game Theory
Acceptance Rate: 74.0%

## Hints

### Hint 1

Both players play *optimally*, so you cannot simulate a greedy "take the bigger end" strategy — taking the bigger pile now can hand the opponent an even better position later. Whenever a game alternates turns and each player maximizes their own outcome, think about **recursion over the remaining state** plus memoization: interval DP on the subarray `piles[i..j]` that is still on the table.

### Hint 2

The state is the pair `(i, j)` — the untaken window. The hard part is that a "score" for one player is not enough: you would have to track whose turn it is *and* both totals. Try collapsing all of that into a single number. Since the game is symmetric (both players want the maximum), you can define the value of a state from the perspective of **whoever is about to move**, and then flip the sign when you recurse into the opponent's turn.

### Hint 3

Let `dp[i][j]` = the best achievable **score difference** (current player's stones minus the other player's stones) for the window `piles[i..j]`. The current player takes `piles[i]` or `piles[j]`, and then the opponent faces the smaller window and achieves *their* best difference — which is negative from our point of view:

```
dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
dp[i][i] = piles[i]
```

Alice wins exactly when `dp[0][n-1] > 0`. Separately, there is a one-line answer worth understanding: with an **even** number of piles, Alice can pre-commit to taking *all* even-indexed piles or *all* odd-indexed piles, and she gets to choose whichever group is larger. Since the total is odd, the two groups can never tie — so Alice always wins.

## Approach

The intended learning value here is the interval DP, so that is what the solution implements; the math shortcut is discussed at the end.

**Reframing the game.** Tracking two separate scores makes the state explode. Instead, evaluate a position by a single number: the *lead* the player to move can guarantee. Because both players use the same rule ("maximize my lead"), the value of a position does not depend on *who* is moving — only on the window of piles left. That is what makes `(i, j)` a sufficient state.

**Recurrence.** From window `piles[i..j]` the mover has two choices:

- Take the left pile: gain `piles[i]`, then the opponent moves on `piles[i+1..j]` and guarantees themselves a lead of `dp[i+1][j]`. Relative to us that lead is negative, so our net is `piles[i] - dp[i+1][j]`.
- Take the right pile: symmetrically, `piles[j] - dp[i][j-1]`.

The mover picks the better of the two, giving `dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])`. The base case is a single pile, `dp[i][i] = piles[i]`: you take it and the game ends, so your lead is exactly that pile. Alice moves first on the full row, so she wins iff `dp[0][n-1] > 0`.

**Walking through `[5,3,4,5]`.** Base row: `dp[0][0]=5, dp[1][1]=3, dp[2][2]=4, dp[3][3]=5`. Windows of length 2: `dp[0][1] = max(5-3, 3-5) = 2`, `dp[1][2] = max(3-4, 4-3) = 1`, `dp[2][3] = max(4-5, 5-4) = 1`. Length 3: `dp[0][2] = max(5-1, 4-2) = 4`, `dp[1][3] = max(3-1, 5-1) = 4`. Length 4: `dp[0][3] = max(5-4, 5-4) = 1 > 0`, so Alice wins — matching the problem's explanation that she finishes 10–9 or so.

**Space optimization.** Rows are filled by increasing window length, and computing length `L` only ever reads length `L-1` entries. Keep one array `dp` where, after processing length `L`, `dp[i]` holds `dp[i][i+L-1]`. When you then compute length `L+1` sweeping `i` upward, `dp[i]` still holds the old `dp[i][j-1]` and `dp[i+1]` still holds the old `dp[i+1][j]` — exactly the two values the recurrence needs. Initialize `dp` to a copy of `piles` for the base case. This drops memory from `O(n^2)` to `O(n)`.

**The O(1) shortcut.** Index the piles `0..n-1` with `n` even. Alice can take pile `0`; that exposes piles `1` and `n-1` — both odd indices — so Bob is forced to take an odd-index pile, which re-exposes only even indices, and so on. By the same argument starting with pile `n-1`, Alice can instead force herself all the odd-index piles. She simply picks the larger of `sum(even indices)` and `sum(odd indices)`; the odd total guarantees they differ. Hence `return true`. It is a genuinely clever observation, but it only works because of the "even count, odd sum" constraints — the DP above is the transferable skill and solves the general version (odd `n`, ties possible) too.

## Complexity Analysis

Time Complexity: O(n^2) — one constant-work transition for each of the ~n^2/2 windows. With `n <= 500` that is well under 10^6 operations.
Space Complexity: O(n) — the rolling one-dimensional DP array. The straightforward two-dimensional table would be O(n^2).

## Edge Cases

- **Empty input** (`piles == nil` or length 0): outside the stated constraints, but the loop bounds and the final `dp[0]` read would panic on an empty slice, so guard explicitly and return `false` — nobody wins a game with no stones.
- **Minimum size, `n == 2`**: the DP never runs a length-3 iteration; the answer comes straight from `max(piles[0]-piles[1], piles[1]-piles[0])`, which is positive whenever the piles differ. Verifies the base case and the length-2 sweep.
- **Exact ties (`dp[0][n-1] == 0`)**: impossible under the constraints since the total is odd, but the comparison must be strict `> 0`, not `>= 0`. A tie is not a win for Alice, and using `>=` would silently return `true` for inputs like `[1,1]`.
- **Odd number of piles**: also excluded by the constraints, and the reason the `return true` shortcut is unsafe as a general answer. The DP handles it correctly — e.g. `[1,5,2]` is a loss for Alice — which is a good way to convince yourself the recurrence is right rather than accidentally correct.
- **Large monotone rows** (e.g. strictly increasing 500 piles): no overflow risk in Go's `int` since the total is at most `500 * 500 = 250000`, but worth noting that intermediate `dp` values are *differences* and are frequently negative — the array must be signed, and initializing it to zeros instead of `piles` would break the base case.
