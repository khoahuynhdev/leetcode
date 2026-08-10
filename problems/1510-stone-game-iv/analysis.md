# 1510. Stone Game IV

[LeetCode Link](https://leetcode.com/problems/stone-game-iv/)

Difficulty: Hard
Topics: Math, Dynamic Programming, Minimax, Game Theory, Nim Game, Sprague–Grundy Theorem, Zero-Sum Game
Acceptance Rate: 60.5%

## Hints

### Hint 1

Don't try to find a formula or a clever pattern in the perfect squares first — that road has swallowed a lot of people on this problem. Instead, notice that the game state is described by a *single* number: how many stones are left. Whose turn it is doesn't change the rules at all, because both players have the exact same set of legal moves. That means every position is either "good for whoever is about to move" or "bad for whoever is about to move." What kind of technique lets you label every one of `n + 1` states?

### Hint 2

Define a boolean `dp[i]` = "the player who must move with `i` stones on the table wins, assuming both play optimally." The base case is easy: `dp[0]` is a loss, because a player facing zero stones has no legal move and therefore loses by the rules. Now think about how `dp[i]` relates to smaller states. From `i` stones you can move to `i - 1`, `i - 4`, `i - 9`, ... — any `i - k*k` that stays non-negative. Each of those hands the turn to your opponent.

### Hint 3

Here is the insight that makes it click: after you remove `k*k` stones, your opponent becomes "the player about to move" in state `i - k*k`. So the position `i - k*k` being a **losing** position is exactly what you want — you'd be handing your opponent a doomed state.

That gives the recurrence:

```
dp[i] = true  if there exists k >= 1 with k*k <= i and dp[i - k*k] == false
dp[i] = false otherwise
```

In words: **a position wins if and only if at least one move leads to a losing position.** If every move you can make leaves your opponent in a winning position, you are lost. This "win iff some move reaches a loss" pattern is the standard characterization of normal-play impartial games, and it's worth memorizing — it recurs in almost every take-away game (Nim, Stone Game series, coin games).

Once you have that, the whole problem is a bottom-up fill from `0` to `n`, trying each square at each index.

## Approach

We build the table bottom-up so that when we evaluate `dp[i]`, every state it depends on (`i - 1`, `i - 4`, `i - 9`, ...) is strictly smaller and already computed.

1. Allocate `dp` of length `n + 1`. In Go, `make([]bool, n+1)` zero-fills to `false`, which conveniently gives us the base case `dp[0] = false` for free: with zero stones the mover cannot move and loses.
2. For each `i` from `1` to `n`, iterate `k = 1, 2, 3, ...` while `k*k <= i`. That enumerates exactly the legal moves from state `i`.
3. If any `dp[i - k*k]` is `false`, the current player can hand the opponent a losing position, so set `dp[i] = true` and break out immediately — one winning move is all it takes, and there's no benefit to examining the rest.
4. If the loop finishes without finding such a move, `dp[i]` stays `false`: every legal move leaves the opponent in a winning position.
5. Return `dp[n]`, which answers the question for Alice, since Alice is the player about to move at the start.

Worth stressing why the "whose turn is it" dimension is absent: the game is *impartial* (both players have identical move options) and *symmetric* in the stone count, so a state is fully described by the number of stones. This is what collapses what looks like a minimax search into a one-dimensional DP.

Walking the first few values by hand:

| `i` | legal moves (`i - k*k`) | reachable states | `dp[i]` |
| --- | --- | --- | --- |
| 0 | — | — | `false` (no move, mover loses) |
| 1 | `1-1` | `dp[0] = false` | `true` |
| 2 | `2-1` | `dp[1] = true` | `false` |
| 3 | `3-1` | `dp[2] = false` | `true` |
| 4 | `4-1`, `4-4` | `dp[3] = true`, `dp[0] = false` | `true` |
| 5 | `5-1`, `5-4` | `dp[4] = true`, `dp[1] = true` | `false` |
| 6 | `6-1`, `6-4` | `dp[5] = false`, ... | `true` |
| 7 | `7-1`, `7-4` | `dp[6] = true`, `dp[3] = true` | `false` |

This matches all three examples: `n = 1` → `true`, `n = 2` → `false`, `n = 4` → `true`. Notice the losing positions so far are `0, 2, 5, 7, ...` — they don't follow any obvious arithmetic pattern, which is exactly why the DP is the right tool and why hunting for a closed form is a trap.

An equivalent top-down memoized recursion works too and reads a little closer to the recurrence, but the bottom-up version avoids recursion depth concerns and has less overhead. Both are the same complexity.

## Complexity Analysis

Time Complexity: O(n * sqrt(n)) — for each of the `n` states we try at most `sqrt(i) <= sqrt(n)` squares. At `n = 10^5` that's roughly 2·10^7 primitive operations in the worst case, comfortably fast in Go. The early `break` makes the typical case much cheaper, since winning positions are common and usually found quickly.

Space Complexity: O(n) for the boolean table. Go's `[]bool` uses one byte per entry, so `n = 10^5` is ~100 KB. There's no way to shrink the window to a constant number of rows here, because `dp[i]` can reach all the way back to `dp[0]`.

## Edge Cases

- **`n = 1`** — the smallest legal input per the constraints. Alice removes the single stone and Bob has no move, so the answer is `true`. Guard that your loop bounds don't skip `i = 1`.
- **`n = 0` is not in the constraints**, but the code should still not panic if handed it. `make([]bool, 1)` and returning `dp[0] = false` handles it gracefully — it also happens to be the game-theoretically correct answer (the mover has no move and loses).
- **Perfect squares** (`n = 4`, `9`, `16`, `100`, ...) are always `true`: Alice sweeps the pile in one move. A quick sanity check on your recurrence, since `dp[n - n] = dp[0] = false`.
- **Losing positions** like `n = 2`, `5`, `7`, `10`, `17` are the ones that actually catch bugs. If you accidentally write `dp[i] = true` when `dp[i - k*k]` is `true` (inverting the condition), the examples with `n = 1` and `n = 4` still pass while every losing position breaks. Always test at least one `false` case.
- **Off-by-one in the square loop** — the condition must be `k*k <= i`, not `k*k < i`. Dropping the equality loses exactly the "clear the pile in one move" option and silently breaks every perfect square.
- **Upper bound `n = 10^5`** — verify the solution finishes quickly and doesn't overflow. `k*k` peaks near `10^5`, so plain `int` is far from any limit; the real concern is only that an exponential brute-force search without memoization would time out.
