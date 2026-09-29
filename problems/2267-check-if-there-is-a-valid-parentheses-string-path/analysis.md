# 2267.  Check if There Is a Valid Parentheses String Path

[LeetCode Link](https://leetcode.com/problems/check-if-there-is-a-valid-parentheses-string-path/)

Difficulty: Hard
Topics: Array, Dynamic Programming, Matrix, Bracket Sequences
Acceptance Rate: 48.9%

## Hints

### Hint 1

How do you check if a single parentheses string is valid without a stack? Think about a single integer counter. Then think about which grid-path technique you know for "moves only down or right".

### Hint 2

The validity of a path prefix is fully described by one number: the current balance (opens minus closes). Two different paths arriving at the same cell with the same balance are interchangeable for the future. That suggests adding the balance as an extra dimension to a grid DP.

### Hint 3

Define `reachable[i][j][b]` = "some path from `(0,0)` to `(i,j)` has balance exactly `b` and never went negative". Transition from the top and left neighbors by applying `+1` or `-1`. The answer is `reachable[m-1][n-1][0]`. Balance can never usefully exceed `(m+n-1)/2`, which keeps the state space small. Also: the path length `m+n-1` must be even — check that first.

## Approach

1. **Quick rejections.** Every path has exactly `m + n - 1` characters. An odd length can never be balanced. The first cell must be `'('` and the last must be `')'`.
2. **State.** For each cell, keep the set of balances reachable by some valid-so-far prefix ending at that cell. A balance is valid while it is `>= 0`.
3. **Transition.** For cell `(i, j)` with delta `+1` for `'('` and `-1` for `')'`, take every balance `b` reachable at `(i-1, j)` or `(i, j-1)` and mark `b + delta` if it is non-negative.
4. **Pruning.** If the balance is larger than the number of cells still left on the path, it can never come back to zero, so drop it. This also bounds balances by about `(m+n)/2`.
5. **Answer.** Return whether balance `0` is reachable at `(m-1, n-1)`.

Because transitions only look at the row above and the cell to the left, a single rolling row of sets is enough: when processing `(i, j)`, `dp[j]` still holds row `i-1` (the cell above) while `dp[j-1]` has already been updated to row `i` (the cell to the left).

**Why it works:** the future of a path depends only on where it is and its current balance, not on the exact characters so far. Grouping paths by `(cell, balance)` collapses exponentially many paths into polynomially many states.

**Example:** `grid = ["(()", "())"]` (length 4). Along the top row the balances are `{1}`, `{2}`, `{1}`. Row two: `(1,0)` gets `{0}`; `(1,1)` gets from top `{2}` and left `{0}` after `-1` → `{1}` (the `-1` from `0` is dropped); `(1,2)` gets from top `{1}` and left `{1}` after `-1` → `{0}`. Balance `0` is reachable, so the answer is `true`.

## Complexity Analysis

Time Complexity: O(m · n · (m + n))
Space Complexity: O(n · (m + n)) with a rolling row

## Edge Cases

- **Odd path length** (e.g. `1x1` or `3x3` grid): can never be valid; return `false` immediately.
- **First cell is `')'` or last cell is `'('`**: every path fails; quick rejection.
- **Single row or single column**: only one path exists; the DP must still handle missing top/left neighbors.
- **Balance dipping below zero mid-path**: a prefix like `"())"` must be discarded even if the total count later balances out.
- **Even length and correct endpoints but still no valid path**: e.g. `["())", ")))"]`, where every path goes negative, which shows that the quick checks alone aren't enough.
