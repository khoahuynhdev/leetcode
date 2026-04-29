# 3225. Maximum Score From Grid Operations

[LeetCode Link](https://leetcode.com/problems/maximum-score-from-grid-operations/)

Difficulty: Hard
Topics: Array, Dynamic Programming, Matrix, Prefix Sum
Acceptance Rate: 45.6%

## Hints

### Hint 1

Each operation is fully described by one number per column: how many cells from
the top of that column are colored black. So the entire game reduces to choosing
a height `h[j] ∈ {0, 1, ..., n}` for every column `j`. Once you accept that, the
problem becomes: pick `n` heights to maximize a score. That should immediately
suggest dynamic programming over columns. Think about what state you need to
carry from one column to the next.

### Hint 2

A white cell `(i, j)` scores if it has a black neighbor to its left or right.
Translated to heights: `(i, j)` contributes `grid[i][j]` iff `i >= h[j]` and
(`i < h[j-1]` or `i < h[j+1]`). The contribution of column `j` therefore depends
on three consecutive heights: `h[j-1]`, `h[j]`, and `h[j+1]`. That three-way
dependency tells you the DP state needs to remember the last *two* heights, not
just one. Per-column prefix sums make computing each contribution O(1).

### Hint 3

With state `(p, c) = (h[j-1], h[j])`, the transition adds *column j-1's*
contribution, which depends on `q = h[j-2]`, `p`, and `c`:

  `contribute(j-1, q, p, c) = sum of grid[i][j-1] for p <= i < max(q, c)`
                            `= max(0, ps[j-1][max(q,c)] - ps[j-1][p])`.

Naively, computing `dp[j][p][c] = max_q (dp[j-1][q][p] + contribute(j-1, q, p, c))`
is `O(n^4)`. The key trick: since prefix sums are monotone,
`max(0, ps[max(q,c)] - ps[p]) = max(K_q, K_c)` where `K_x = max(0, ps[x] - ps[p])`.
Therefore

  `max_q (dp[j-1][q][p] + max(K_q, K_c)) = max(M1, M2 + K_c)`

with `M1 = max_q (dp[j-1][q][p] + K_q)` and `M2 = max_q dp[j-1][q][p]`. Both `M1`
and `M2` depend only on `(j, p)`, so each `(j, p)` is `O(n)` and the total is
`O(n^3)`.

## Approach

1. **Reformulate**. Replace the operation set with a height vector
   `h[0..n-1]`, `h[j] ∈ {0..n}`. White cell `(i, j)` scores iff `i >= h[j]`
   and `i < h[j-1]` or `i < h[j+1]` (treat out-of-range heights as 0).

2. **Per-column prefix sums.** Build `ps[j][k] = grid[0][j] + ... + grid[k-1][j]`.
   Then the white contribution of column `j` is
   `max(0, ps[j][max(h[j-1], h[j+1])] - ps[j][h[j]])`, computable in O(1).

3. **DP state.** `dp[p][c]` holds the maximum cumulative score when the last
   two chosen heights are `(p, c) = (h[j-1], h[j])`, and contributions from
   columns `0..j-1` are already accounted for. Initialize `dp[0][c] = 0` for
   every `c` (a virtual `h[-1] = 0`).

4. **Transition.** Moving to the next column adds column `j-1`'s contribution.
   For each `(p, c)`,
   `new_dp[p][c] = max_q (dp[q][p] + max(0, ps[j-1][max(q,c)] - ps[j-1][p]))`.
   Decompose the inner `max(K_q, K_c)`: precompute, per `(j, p)`,
   `M1 = max_q (dp[q][p] + K_q)` and `M2 = max_q dp[q][p]`, then
   `new_dp[p][c] = max(M1, M2 + K_c)`.

5. **Final column.** After processing column `n-1`, dp encodes
   `(h[n-2], h[n-1])` but column `n-1`'s own white contribution still has to be
   added, with virtual `h[n] = 0`:
   `answer = max_{p,c} dp[p][c] + max(0, ps[n-1][p] - ps[n-1][c])`.

**Walking through Example 1** (`grid` is 5×5, expected 11): the optimal choice
is `h = [0, 4, 0, 0, 5]`. Column 0's white rows 0..3 contribute `grid[3][0]=5`
because they sit left of column 1 which has height 4. Column 2's white row 1
contributes `grid[1][2]=3` for the same reason. Column 3's white row 3
contributes `grid[3][3]=3` because column 4 has height 5. Total: `5+3+3=11`. The
DP arrives at this through transitions `(0,0) → (0,4) → (4,0) → (0,0) → (0,5)`.

## Complexity Analysis

Time Complexity: `O(n^3)` — `n` columns, `O(n)` outer height `p`, and `O(n)` to
compute `M1`, `M2`, and the per-`c` updates. For `n = 100` that's `10^6`
operations.

Space Complexity: `O(n^2)` for the rolling DP table plus `O(n^2)` for prefix
sums.

## Edge Cases

- **`n = 1`**: A single column has no horizontal neighbors, so no cell can ever
  satisfy the "white with adjacent black" rule. Answer is 0. Handle with an
  early return so `n - 1 = 0` indexing into `ps[col]` stays sane.
- **All-zero grid**: every contribution is 0. The DP returns 0, never going
  negative because all `K_x` clamp at 0.
- **Big values (`grid[i][j] ≤ 10^9`)**: cumulative sums can reach `10^4 · 10^9 =
  10^13`, well past `int32`. Use `int64` throughout (which is also what the
  return type demands).
- **Heights at the extremes**: `h[j] = 0` (no operation) and `h[j] = n` (entire
  column black) are both valid choices and the DP must consider them. The state
  range `0..n` (inclusive) covers this — note the extra `+1` in array sizes.
- **Two-column grids**: there is no "middle" column, but the DP still works
  because the final-column step explicitly handles the missing right neighbor
  with a virtual zero height.
