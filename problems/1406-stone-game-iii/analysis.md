# 1406. Stone Game III

[LeetCode Link](https://leetcode.com/problems/stone-game-iii/)

Difficulty: Hard
Topics: Array, Math, Dynamic Programming, Game Theory
Acceptance Rate: 64.2%

## Hints

### Hint 1

Both players always take from the **front** of the row, so the entire remaining game is described by a single number: the index `i` of the first untaken stone. That is a strong signal — a one-dimensional state means a one-dimensional DP. Try to define `f(i)` = "the best outcome achievable by whoever moves first when only `stoneValue[i:]` is left", and notice that the two players are interchangeable: they face exactly the same rules, so one function serves both.

### Hint 2

The obvious definition, `f(i)` = "maximum score the current player can collect", does not work. The current player's score alone tells you nothing about who wins — you would also need to track what the opponent collects, and that depends on choices made later. Look for a value that folds *both* scores into one number, so the recurrence stays self-contained. Also note the answer only depends on the *difference* between the two final scores, never on their absolute sizes.

### Hint 3

Define `dp[i]` = (current player's score − opponent's score) over the suffix `stoneValue[i:]`, both playing optimally. If the current player takes `k ∈ {1,2,3}` stones, they bank `sum(stoneValue[i..i+k-1])`, and then the roles flip: from position `i+k` the *opponent* becomes the "current player", and their advantage there is `dp[i+k]`. Relative to the player who just moved, that advantage counts negatively. So

```
dp[i] = max over k in {1,2,3}, i+k <= n  of  ( sum(stoneValue[i..i+k-1]) - dp[i+k] )
dp[n] = 0
```

That single sign flip is the whole trick. `dp[0]` is Alice's final margin: positive → `"Alice"`, negative → `"Bob"`, zero → `"Tie"`.

## Approach

This is the classic **minimax-as-score-difference** pattern. The naive framing ("maximize my own score") forces you to track two quantities at once; reframing as "maximize my lead" collapses the game to one number per state, and because the game is symmetric — Bob faces exactly the rules Alice faces — the same array serves both players.

**State.** `dp[i]` = the largest value of `(score of the player to move) − (score of the other player)` achievable from the suffix starting at index `i`, assuming optimal play by both.

**Base case.** `dp[n] = 0`. With no stones left, neither player scores again, so the lead is 0.

**Transition.** From `i` the mover picks `k = 1`, `2`, or `3` stones (only choices with `i + k <= n` are legal). They immediately add `prefix = stoneValue[i] + ... + stoneValue[i+k-1]` to their own score. Then the turn passes. From `i+k`, the *opponent* is the mover, and `dp[i+k]` is the opponent's lead over our player. A lead for them is a deficit for us, hence the subtraction:

```
dp[i] = max_{k=1..3} ( prefix_k - dp[i+k] )
```

Taking the max is exactly "play optimally": each player, on their own turn, picks the move maximizing their own final margin, and the sign flip automatically encodes that the opponent will do the same to us.

**Order of evaluation.** `dp[i]` depends only on `dp[i+1..i+3]`, so fill the table from right to left. Since the window is at most three cells wide, you never need the full array — three rolling variables suffice, giving O(1) extra space. The implementation keeps `next1`, `next2`, `next3` for `dp[i+1]`, `dp[i+2]`, `dp[i+3]` and shifts them each iteration. Out-of-range positions are initialized to 0 and guarded by the `i+k <= n` check, so they are never read as if they were real states.

**Answer.** Alice moves first at index 0, so `dp[0]` *is* Alice's final margin. Return `"Alice"` if `dp[0] > 0`, `"Bob"` if `dp[0] < 0`, `"Tie"` if `dp[0] == 0`.

**Worked example — `stoneValue = [1,2,3,7]` (expected `"Bob"`).** Filling right to left:

- `dp[4] = 0` (base case)
- `dp[3] = 7 - dp[4] = 7` — only one stone left, take it.
- `dp[2] = max(3 - dp[3], (3+7) - dp[4]) = max(3-7, 10-0) = 10` — grab both.
- `dp[1] = max(2 - dp[2], 5 - dp[3], 12 - dp[4]) = max(-8, -2, 12) = 12`
- `dp[0] = max(1 - dp[1], 3 - dp[2], 6 - dp[3]) = max(1-12, 3-10, 6-7) = -1`

`dp[0] = -1 < 0`, so Bob wins by exactly 1 — matching the problem's explanation that Alice's best line is taking `[1,2,3] = 6` and conceding the `7`. Notice how every branch is bad for Alice; the DP does not need to know *why*, it just picks the least-bad one.

A word of encouragement: the code here is about fifteen lines, and it is genuinely short. The hard part of this problem is not implementation, it is trusting the reframing in Hint 3. If it feels slippery, hand-simulate `[1,2,3,6]` (the Tie example) and watch the signs alternate — that usually makes it click for good. The same trick powers Stone Game I/II, Predict the Winner, and most "two players alternate, both optimal" problems, so it is worth the investment.

## Complexity Analysis

Time Complexity: O(n) — one pass over the array, with a constant three-branch max at each index.
Space Complexity: O(1) — three rolling variables replace the `dp` table. (The straightforward table version is O(n) space and equally acceptable at `n <= 5 * 10^4`.)

## Edge Cases

- **Single stone (`n = 1`).** Alice must take it, and the sign of that lone value decides everything: `[5]` → `"Alice"`, `[-1]` → `"Bob"`, `[0]` → `"Tie"`. The `i + k <= n` guard is what keeps `k = 2, 3` from reading past the end.
- **Fewer than 3 stones remaining.** Near the right edge, only some of the three moves are legal. Forgetting the bounds check either panics or silently treats a nonexistent state as `0`, which would let a player "pass" — a move the rules do not allow.
- **Negative values.** Values range down to `-1000`, so taking *more* stones can be strictly worse. This breaks any greedy instinct ("always take the largest prefix") and is exactly what Example 2 (`[1,2,3,-9]`) is testing. It also means `dp` values can legitimately be negative, so initializing the running max to `0` instead of `-inf` (or to the `k = 1` branch) is a real bug.
- **All zeros / exact ties.** `dp[0] == 0` must map to `"Tie"`, not fall through to a two-way Alice/Bob branch. Example 3 (`[1,2,3,6]`) is the non-trivial version of this.
- **Total sum overflow — not an issue here.** `5 * 10^4` stones times `1000` caps the magnitude at `5 * 10^7`, comfortably inside `int32`, and Go's `int` is 64-bit on all target platforms. Worth confirming rather than assuming.
- **Recursion depth.** A top-down memoized solution recurses up to `5 * 10^4` frames deep. Go's growable goroutine stacks handle that, but the bottom-up loop sidesteps the question entirely and is faster in practice.
- **Empty input.** Outside the stated constraints (`n >= 1`), but the loop body simply never runs and `dp[0] = 0` yields `"Tie"` — a sane, non-panicking fallback.
