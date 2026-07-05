# 1301. Number of Paths with Max Score

[LeetCode Link](https://leetcode.com/problems/number-of-paths-with-max-score/)

Difficulty: Hard
Topics: Array, Dynamic Programming, Matrix
Acceptance Rate: 49.6%

## Hints

### Hint 1

You start at the bottom-right `'S'` and must reach the top-left `'E'`, moving only
up, left, or up-left. Notice that every legal move strictly decreases your position
toward `(0, 0)`. Whenever a problem asks for the *best value over all paths in a grid*
where movement is restricted to a fixed set of directions, think **dynamic programming
over the grid** rather than brute-force search — the number of paths is exponential,
but the number of cells is not.

### Hint 2

This problem has two coupled quantities: the **maximum score** reachable at a cell, and
the **number of distinct paths** that achieve that maximum. A single DP value is not
enough — carry a *pair* `(bestSum, count)` for each cell. When you combine several
incoming directions, the maximum sum is decided first, and only the neighbors that tie
for that maximum contribute to the path count. Remember to take the count modulo
`10^9 + 7`.

### Hint 3

Fill the table in an order where a cell's predecessors are already computed. Because a
move goes *toward* `(0,0)`, a cell `(i, j)` is reached *from* one of `(i+1, j)`,
`(i, j+1)`, or `(i+1, j+1)` (the cells you were standing on before the move). So iterate
`i` from `n-1` down to `0` and `j` from `n-1` down to `0`. For each non-obstacle cell,
look at those three predecessors, pick the maximum reachable sum, sum the path counts of
every predecessor tying for that maximum, then add the cell's own digit value (`'E'` and
`'S'` contribute `0`). Mark obstacles and truly unreachable cells so they can never feed
a path.

## Approach

We run a bottom-up DP on the `n x n` board.

Define two tables:

- `dp[i][j]` = the maximum sum of digits collectable on a path from `S` (the
  bottom-right cell) to cell `(i, j)`. We use `-1` as a sentinel meaning "unreachable".
- `cnt[i][j]` = the number of distinct paths from `S` to `(i, j)` achieving `dp[i][j]`,
  taken modulo `10^9 + 7`.

**Base case.** The journey begins on `S` at `(n-1, n-1)`, so `dp[n-1][n-1] = 0` and
`cnt[n-1][n-1] = 1`.

**Transition.** A move from the previous cell to the current cell goes up, left, or
up-left, which means the current cell `(i, j)` is entered *from* one of its three
"lower/right" neighbors: `(i+1, j)`, `(i, j+1)`, `(i+1, j+1)`. To have those neighbors
already solved, we sweep `i` from `n-1` down to `0` and, inside, `j` from `n-1` down to
`0`.

For each cell (other than the start):

1. If it is an obstacle `'X'`, leave it unreachable.
2. Otherwise scan the three predecessors. Track `best`, the largest `dp` value among the
   reachable ones, and `ways`, the sum (mod `1e9+7`) of the `cnt` values of every
   predecessor whose `dp` equals `best`.
3. If no predecessor is reachable, this cell is unreachable too.
4. Otherwise set `dp[i][j] = best + value`, where `value` is the digit's numeric value,
   or `0` when the cell is `'E'` (or `'S'`). Set `cnt[i][j] = ways`.

**Answer.** The destination is `E` at `(0, 0)`. If `dp[0][0]` is still the sentinel, no
path exists, so return `[0, 0]`. Otherwise return `[dp[0][0], cnt[0][0]]`.

**Worked example.** For `board = ["E23","2X2","12S"]`, the sweep computes a best sum of
`7` reaching `(0,0)` along the single path `S(2,2) → (2,1)=2 → (2,0)=1 → (1,0)=2 →
(0,0)=E`... wait, the actual max path collects `2 + 2 + 3 = 7` via
`S → (0,2)=3 → (0,1)=2 → E` combined with the right column, and exactly one path attains
`7`, giving `[7, 1]`. In `["E12","1X1","21S"]` two different routes both collect `4`, so
the answer is `[4, 2]`, and the count-merging step is what distinguishes them.

## Complexity Analysis

Time Complexity: O(n^2) — each of the `n^2` cells does a constant amount of work
(inspecting three predecessors).
Space Complexity: O(n^2) for the `dp` and `cnt` tables. This can be reduced to O(n) by
keeping only two rows, but O(n^2) is clear and well within the constraints
(`n <= 100`).

## Edge Cases

- **No path exists (example 3, `["E11","XXX","11S"]`).** A wall of obstacles can cut off
  `E` entirely. The sentinel `-1` propagates: since no predecessor is reachable, `E`
  stays unreachable and we return `[0, 0]` rather than a spurious sum.
- **Obstacle handling.** An `'X'` cell must never contribute a path; treating it as
  unreachable (never assigning it a `dp`) prevents it from feeding downstream cells.
- **`'E'` and `'S'` contribute zero.** They are not digits; adding a numeric value for
  them would inflate the sum. Only characters `'1'..'9'` add their value.
- **Counting with ties.** When two or three predecessors tie for the best sum, their path
  counts must be *added*, not overwritten — this is the whole point of the second output.
- **Modulo on the count only.** The maximum sum is small (at most `9 * (2n-2)`), so only
  the path count needs the `10^9 + 7` modulo; applying it to the sum would be wrong.
- **Minimum board size.** With `n = 2`, `E` and `S` are always distinct cells and the
  three-predecessor scan still works because out-of-range neighbors are simply skipped.
