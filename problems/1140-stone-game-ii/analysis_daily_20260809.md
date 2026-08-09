# 1140. Stone Game II

[LeetCode Link](https://leetcode.com/problems/stone-game-ii/)

Difficulty: Medium
Topics: Array, Math, Dynamic Programming, Minimax, Prefix Sum, Game Theory, Zero-Sum Game
Acceptance Rate: 73.0%

## Hints

### Hint 1

Don't reach for a greedy rule here — "take the biggest reachable chunk" loses, and so does "always take one pile." Both players play perfectly, so you need to *search* the game tree rather than guess a rule. The right frame is minimax with memoization: figure out what the complete description of a game position is, and how many distinct positions there can be. If that count is small (polynomial in `n`), a DP will fit comfortably inside `n <= 100`.

### Hint 2

Ask yourself what a player actually needs to know to play the rest of the game optimally. Not the history, not who has how many stones so far — moves only ever remove a *prefix* of what's left, and the rules only reference `M`. So a position is fully described by two numbers. Once you name those two numbers, you have your DP state. Also notice that a move takes a contiguous block starting at a known index, so prefix/suffix sums let you price any move in `O(1)`.

### Hint 3

The trick that makes this clean is exploiting that the game is **zero-sum**: every remaining stone ends up with somebody. Define `dp[i][m]` as the best total the player *whose turn it is* can get from `piles[i:]` — note this is deliberately player-agnostic, so the same table serves Alice and Bob.

If that player takes `X` piles, the opponent inherits the position `(i+X, max(m, X))` and, by the very same definition, will secure `dp[i+X][max(m,X)]` of the `suffix[i+X]` stones left. So the mover walks away with

```
suffix[i] - dp[i+X][max(m, X)]
```

and `dp[i][m] = max` of that over all legal `X`. That subtraction is the whole problem: it turns a two-player minimax into a single recurrence with no "whose turn is it" bookkeeping. Base case: if `i + 2m >= n` the mover can sweep the entire remainder, so `dp[i][m] = suffix[i]`. Answer: `dp[0][1]`.

## Approach

**State.** `dp[i][m]` = the maximum number of stones the player about to move can guarantee from `piles[i..n-1]`, given the current value of `M = m`. It doesn't matter *which* player that is — the rules are symmetric, so both players evaluate positions with the same table.

**Suffix sums.** Precompute `suffix[i] = piles[i] + ... + piles[n-1]` with `suffix[n] = 0`, scanning right to left. This gives the total stones still on the table at any index in `O(1)`.

**Base case.** If `i >= n` the position is empty and worth `0`. More usefully: if `i + 2m >= n`, the mover is allowed to take *every* remaining pile in a single turn, so `dp[i][m] = suffix[i]` and there is nothing to decide. This case is what terminates the recursion and it also handles the tail of the row cheaply.

**Transition.** Otherwise, try each legal `X` from `1` to `2m` (all of which leave at least one pile behind, since `i + 2m < n`):

```
dp[i][m] = max over X in [1, 2m] of ( suffix[i] - dp[i+X][max(m, X)] )
```

The opponent, facing `(i+X, max(m,X))`, plays optimally and takes `dp[i+X][max(m,X)]`. Because no stone is ever left over, everything the opponent does *not* take from that suffix eventually lands with the current mover — hence the subtraction. Maximizing our own take is therefore exactly the same as minimizing the opponent's, which is what makes minimax collapse into a plain `max`.

**Order of evaluation.** `dp[i][m]` depends only on rows with a strictly larger `i`, so iterate `i` from `n-1` down to `0`; `m` can go in any order within a row. `m` never needs to exceed `n`, since taking more than `n` piles is impossible. Return `dp[0][1]` — Alice moves first with `M = 1`.

**Worked example: `piles = [2,7,9,4,4]`, `suffix = [26,24,17,8,4,0]`.**

- `dp[3][*] = 8` and `dp[4][*] = 4` — from index 3 or 4, two or fewer piles remain, so the mover sweeps them.
- `dp[2][1]`: `2 + 2 = 4 < 5`, so we compare `X=1 → 17 - dp[3][1] = 17 - 8 = 9` against `X=2 → 17 - dp[4][2] = 17 - 4 = 13`. So `dp[2][1] = 13`: take two piles (`9 + 4 = 13`) and leave the final `4` for the opponent.
- `dp[1][1]`: `1 + 2 = 3 < 5`. `X=1 → 24 - dp[2][1] = 24 - 13 = 11`; `X=2 → 24 - dp[3][2] = 24 - 8 = 16`. So `dp[1][1] = 16`.
- `dp[0][1]`: `X=1 → 26 - dp[1][1] = 26 - 16 = 10`; `X=2 → 26 - dp[2][2] = 26 - 17 = 9` (from index 2 with `m=2`, `2+4 >= 5`, so the opponent sweeps all 17). Max is **10**. ✅

That last line is precisely the reasoning in the problem's explanation: opening with one pile beats opening with two, because opening with two hands Bob a position from which he can take everything left.

Honest note on difficulty: the recurrence is four lines of code, but the `suffix[i] - dp[next]` inversion is the kind of insight that feels obvious only *after* you've seen it. If it didn't click on first read, that's normal — the zero-sum trick recurs across the whole Stone Game family (877, 1140, 1406, 1563), so learning it here pays off repeatedly.

## Complexity Analysis

Time Complexity: O(n^3) — there are `O(n^2)` states `(i, m)` and each scans up to `2m = O(n)` choices of `X`. With `n <= 100` that's at most ~10^6 cheap operations.
Space Complexity: O(n^2) for the DP table, plus `O(n)` for the suffix sums.

## Edge Cases

- **Single pile (`n = 1`).** `i + 2m = 0 + 2 >= 1`, so the base case fires immediately and Alice takes the whole pile. Worth checking that your base-case condition is `>=` and not `>`, or you'll index out of bounds.
- **Two piles.** With `M = 1` Alice may take up to `2` piles, so she takes both — no matter how the values are distributed. A solution that greedily takes one pile gets this wrong.
- **Empty input.** LeetCode's constraints guarantee `piles.length >= 1`, but returning `0` for an empty slice keeps the function total and avoids a panic when building the DP table. Cheap insurance.
- **The `max(M, X)` update.** `M` only ever grows, and it becomes `max(M, X)` — not `X`. Taking a small `X` when `M` is already large must not shrink the opponent's allowance.
- **Sweep-the-rest condition (`i + 2m >= n`).** Getting this boundary right is what keeps `i + X` inside the array in the transition loop. Once you've handled it, the inner loop over `X in [1, 2m]` never needs its own bounds check.
- **All piles equal.** e.g. `[1,1,1,1] → 2`. A useful sanity check: with symmetric values the split is even, and any bug in the turn-alternation logic tends to show up as an off-by-one here.
- **A huge pile at the end.** e.g. `[1,2,3,4,5,100] → 104`. Optimal play means *deliberately taking little early* so that `M` stays small for the opponent and the big tail stays reachable. Greedy approaches fail this case loudly.
- **Upper value bound.** `n = 100` piles of `10^4` stones tops out at `10^6`, so plain `int` is far from overflow in Go.
