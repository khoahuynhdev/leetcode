# 0486. Predict the Winner

[LeetCode Link](https://leetcode.com/problems/predict-the-winner/)

Difficulty: Medium
Topics: Array, Math, Dynamic Programming, Recursion, Game Theory
Acceptance Rate: 56.7%

## Hints

### Hint 1

Greedy fails here — always grabbing the larger end is not optimal (try `[1, 5, 233, 7]`, where player 1 must *give up* the bigger end first). Whenever a greedy choice is wrong because it changes what your opponent can do next, think about exploring the game tree. Notice that every reachable state of the array is a **contiguous subarray** `nums[i..j]`, because both players only ever remove from the ends. That is a very small state space: at most `n²/2` states. What technique does "small state space + overlapping subproblems" suggest?

### Hint 2

Define a recursive function over the interval `[i, j]` and let it answer the question *"what is the best outcome for whoever is about to move on `nums[i..j]`?"* The trick is that the two players are symmetric: the same function describes both of them, just called on a smaller interval. So you never need to track *which* player it is — only whose turn it is implicitly, by recursion depth. Memoize on `(i, j)`.

The remaining question is what "best outcome" should mean. Tracking `(player1Score, player2Score)` works but is clumsy. Is there a single number that captures everything you need?

### Hint 3

Track the **score difference** from the perspective of the player to move, not the two scores separately.

Let `dp[i][j]` = the maximum value of *(current player's total) − (other player's total)* over the subarray `nums[i..j]`, when both play optimally.

If the current player takes `nums[i]`, they gain `nums[i]`, and then the roles flip: the opponent becomes "the current player" on `nums[i+1..j]` and achieves `dp[i+1][j]` *from their own perspective*. Since a gain for them is a loss for us, we subtract:

```
dp[i][j] = max(nums[i] - dp[i+1][j],
               nums[j] - dp[i][j-1])
```

Base case: `dp[i][i] = nums[i]`. Player 1 wins or ties exactly when `dp[0][n-1] >= 0`.

That single sign flip is the whole problem — it collapses a two-player game into an ordinary interval DP.

## Approach

**The state.** Since both players remove only from the ends, every position the game can ever reach is a contiguous window `nums[i..j]`. The identity of the mover doesn't need to be part of the state: the player to move on `nums[i..j]` is determined by how many elements have been removed, and — more importantly — the *rules* are identical for both players, so one function serves both.

**The value.** Instead of two scores, store one number: the best achievable `myScore - opponentScore` for the player about to move. This is the standard "negamax" reformulation of a zero-sum game. It works because the final question ("did player 1 score at least as much as player 2?") only depends on the difference, and because the difference negates cleanly when you hand the turn over.

**The recurrence.** From `nums[i..j]` the mover has exactly two options:

- Take `nums[i]`: immediate gain `nums[i]`, then the opponent faces `nums[i+1..j]` and will achieve `dp[i+1][j]` *in their favor*, so our net is `nums[i] - dp[i+1][j]`.
- Take `nums[j]`: by the same reasoning, `nums[j] - dp[i][j-1]`.

The mover picks the maximum. The base case is a single element, `dp[i][i] = nums[i]`, since the mover takes it and the opponent gets nothing.

**Walkthrough on `nums = [1, 5, 2]`.**

Single elements: `dp[0][0] = 1`, `dp[1][1] = 5`, `dp[2][2] = 2`.

Length 2:
- `dp[1][2] = max(5 - dp[2][2], 2 - dp[1][1]) = max(5-2, 2-5) = 3`
- `dp[0][1] = max(1 - dp[1][1], 5 - dp[0][0]) = max(1-5, 5-1) = 4`

Length 3:
- `dp[0][2] = max(nums[0] - dp[1][2], nums[2] - dp[0][1]) = max(1-3, 2-4) = -2`

`dp[0][2] = -2 < 0`, so player 1 loses by 2 no matter what → `false`, matching the expected output. The DP is telling us both of player 1's opening moves hand the `5` to player 2.

**Fill order.** `dp[i][j]` depends on `dp[i+1][j]` (a larger `i`) and `dp[i][j-1]` (a smaller `j`), so iterate `i` from `n-1` down to `0` and `j` from `i+1` up to `n-1`. Equivalently, iterate by increasing interval length.

**Space optimization.** With that loop order, the row for `i` only ever reads the row for `i+1` at the same column `j`, plus the already-updated entry `dp[i][j-1]` in the current row. So a single 1-D array of length `n` suffices: when processing index `i`, `dp[j]` still holds `dp[i+1][j]` before we overwrite it, and `dp[j-1]` already holds the new `dp[i][j-1]`. That is exactly the two values the recurrence needs. The implementation uses this 1-D form.

Note that with `n <= 20` a plain memoized recursion would also pass comfortably; the bottom-up version is here because it is allocation-free per call and makes the dependency structure explicit.

## Complexity Analysis

Time Complexity: O(n²) — one constant-work transition for each of the ~n²/2 intervals `[i, j]`.
Space Complexity: O(n) — a single rolling row. The straightforward 2-D table is O(n²); the memoized recursion is O(n²) table plus O(n) stack.

## Edge Cases

- **Single element (`n == 1`)**: player 1 takes it and player 2 scores 0. `dp[0][0] = nums[0] >= 0` since values are non-negative, so the answer is `true`. The loop body never executes, so this falls out of the base case for free.
- **Ties count as a win**: the problem says equal scores means player 1 wins. This is why the final comparison is `dp[0][n-1] >= 0` and not `> 0`. Using `>` is the single most common bug here — `[1, 1]` and `[2, 2, 2, 2]` both tie and must return `true`.
- **All zeros, e.g. `[0, 0, 0]`**: every difference is `0`, so the answer is `true`. A good check that the `>=` comparison is right.
- **Even-length arrays**: player 1 can always guarantee at least a tie (take all even-indexed or all odd-indexed elements, whichever sum is larger — player 1 can force either). So the answer is always `true` for even `n`. Not needed for correctness, but a useful sanity check: if your code ever returns `false` for an even-length input, it is wrong.
- **Non-negative values including 0** (`0 <= nums[i] <= 10^7`): no negative numbers exist, so a player is never forced to take something harmful. Don't let that tempt you into a greedy shortcut — the difficulty is positional, not sign-related.
- **Overflow**: with `n <= 20` and values up to `10^7`, the maximum total is `2 * 10^8`, which fits in a 32-bit `int` and is far inside Go's `int`. No special handling needed, but worth noting for languages with narrower defaults.
- **Empty input**: not reachable under the constraints (`n >= 1`), but the code returns `true` for it rather than indexing out of bounds, which keeps the function total.
