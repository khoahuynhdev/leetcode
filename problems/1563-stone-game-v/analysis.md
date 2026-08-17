# 1563. Stone Game V

[LeetCode Link](https://leetcode.com/problems/stone-game-v/)

Difficulty: Hard
Topics: Array, Math, Dynamic Programming, Game Theory
Acceptance Rate: 44.7%

## Hints

### Hint 1

Despite the "Game Theory" tag, only one player actually makes choices. Bob's move is
forced: he always throws away the heavier row. So this is not a minimax problem at all —
it is a pure maximization over Alice's choices.

Every round leaves Alice with a *contiguous* sub-row of the original array, and the game
that follows depends only on that sub-row, not on how she got there. That "state = interval,
future is independent of the past" property is the signature of interval DP. Think about
what `dp[i][j]` should mean.

### Hint 2

Let `dp[i][j]` be the maximum score Alice can still collect starting from the sub-row
`stoneValue[i..j]`. With prefix sums you can get any range sum in O(1), so for each split
point `k` (left = `i..k`, right = `k+1..j`) you can compare the two halves in constant time
and write a recurrence:

- left sum < right sum → Bob discards the right, Alice banks `sum(i,k)` and continues on `i..k`
- left sum > right sum → Bob discards the left, Alice banks `sum(k+1,j)` and continues on `k+1..j`
- the sums are equal → Alice picks whichever side has the better future

Base case: `dp[i][i] = 0`, since a single stone cannot be split. That gives an O(n³)
algorithm, which for `n = 500` is about 1.25 × 10⁸ split evaluations — usually enough to
pass, but it is worth pushing further.

### Hint 3

Here is the insight that removes a whole factor of `n`.

Rewrite the recurrence in terms of a single quantity: define

```
f(x, y) = sum(x, y) + dp[x][y]
```

which is exactly "the score Alice gets if the segment `x..y` is the one that survives this
round." Then every branch of the recurrence is just a maximum of `f` values over a
*contiguous range of split points*:

- keeping the left half contributes `max f(i, k)` over all `k` where left ≤ right
- keeping the right half contributes `max f(k+1, j)` over all `k` where left ≥ right

The equality case needs no special handling — when the halves are equal, that split point
belongs to *both* ranges, which is precisely Alice choosing the better side.

Two more observations finish it:

1. Because `sum(i,k)` grows with `k` while `sum(k+1,j)` shrinks, the boundary between
   "left is lighter" and "left is heavier" is a single point, and for a fixed `i` that
   boundary only moves *right* as `j` grows. A monotone two-pointer finds it in amortized
   O(1).
2. Those range maxima are prefix/suffix maxima along a row and a column, so they can be
   maintained incrementally in O(1) as the DP table fills.

## Approach

Let `prefix` be the prefix-sum array so that `sum(i, j) = prefix[j+1] - prefix[i]`, and let

```
dp[i][j]       = best score Alice can still obtain from stoneValue[i..j]
f(i, j)        = sum(i, j) + dp[i][j]
maxLeft[i][j]  = max over t in [i, j] of f(i, t)     // prefix max along row i
maxRight[i][j] = max over t in [i, j] of f(t, j)     // suffix max along column j
```

`f(i, j)` reads as: "if segment `i..j` is the half that Bob spares, Alice immediately banks
`sum(i, j)` and then plays out `i..j` optimally." `maxLeft` and `maxRight` are just running
maxima of `f` over one fixed endpoint, which is exactly the shape the recurrence needs.

**Iteration order.** Loop `i` from `n-1` down to `0` and, inside, `j` from `i+1` up to `n-1`.
Computing `dp[i][j]` needs `f(i, t)` for `t < j` (same row, already filled because `j`
increases) and `f(t, j)` for `t > i` (rows below `i`, already fully filled because `i`
decreases). Both dependencies are satisfied.

**Finding the boundary.** For a fixed `i`, keep a pointer `split` and advance it while
`2 * sum(i, split) < sum(i, j)`. When it stops, `split` is the smallest index whose left
half is *not* lighter than its right half. Since `sum(i, j)` only grows as `j` grows, the
pointer never moves backwards, so the total work across the whole row is O(n).

From `split` derive the two ranges of usable split points:

- if `split == j`, no split makes the left half heavy, so left-keeping covers `k ∈ [i, j-1]`
  and right-keeping covers nothing;
- if `2 * sum(i, split) == sum(i, j)` the halves tie there, so that index belongs to both
  ranges (`k0 = k1 = split`);
- otherwise the left half at `split` is strictly heavier, so left-keeping covers
  `k ∈ [i, split-1]` and right-keeping covers `k ∈ [split, j-1]`.

Then

```
dp[i][j] = max( maxLeft[i][k0],            // if k0 >= i
                maxRight[k1+1][j] )        // if k1 <= j-1
```

and finally extend the running maxima with the value just computed:

```
v := sum(i, j) + dp[i][j]
maxLeft[i][j]  = max(maxLeft[i][j-1],  v)
maxRight[i][j] = max(maxRight[i+1][j], v)
```

The answer is `dp[0][n-1]`.

**Worked example: `[6,2,3,4,5,5]`, total 25.** For `i = 0, j = 5` the pointer stops at
`split = 3` (`sum(0,3) = 15`, and `2·15 > 25`, while `2·sum(0,2) = 22 < 25`). So
left-keeping ranges over `k ∈ [0, 2]` and right-keeping over `k ∈ [3, 4]`. The best entry in
`maxLeft[0][2]` is `f(0, 2) = 11 + dp[0][2] = 11 + 7 = 18` — that is Alice cutting into
`[6,2,3] | [4,5,5]`, banking 11, and then extracting 7 more from `[6,2,3]` (split off
`[6]`, bank 5, then split `[2],[3]`, bank 2). Right-keeping tops out lower, so
`dp[0][5] = 18`.

Note that `dp[i][i] = 0` is left at its zero value, while the bases
`maxLeft[i][i] = maxRight[i][i] = stoneValue[i]` must be set explicitly — they encode
`f(i, i) = stoneValue[i] + 0`.

## Complexity Analysis

Time Complexity: O(n²) — each of the O(n²) intervals does O(1) work, and the monotone
pointer contributes O(n) amortized per row. (The straightforward recurrence is O(n³).)

Space Complexity: O(n²) for the three `n × n` tables. At `n = 500` that is about 6 MB with
64-bit ints, comfortably within limits.

## Edge Cases

- **Single stone (`n == 1`)**: the game is already over, so the answer is `0`. Guard this
  before allocating, otherwise `dp[0][0]` is returned uninitialized-by-accident (it happens
  to be `0`, but relying on that is fragile). Example 3 tests exactly this.
- **Empty input**: not allowed by the constraints, but the same `n < 2` guard returns `0`
  instead of indexing out of bounds.
- **Two stones**: the only split is forced. If they are equal (`[5,5]`) Alice picks a side
  and scores `5`; if not (`[10,1]`) she is stuck with the smaller one and scores `1`. This
  is the case most likely to expose an off-by-one in the boundary logic, because `k0` can
  legitimately fall to `i-1` (no left option at all).
- **Ties everywhere (`[7,7,7,7,7,7,7]`)**: the equality branch is where naive
  implementations go wrong, either by forgetting Alice gets to choose or by double-counting.
  The range formulation handles it for free because a tied index sits in both ranges.
- **All values identical but odd length**: the boundary never lands exactly on a tie for
  odd-length windows, so both the strict-inequality paths get exercised.
- **Large values**: `500 × 10⁶ = 5 × 10⁸` fits in `int32`, but intermediate `2 * sum(...)`
  comparisons reach `10⁹`. Go's `int` is 64-bit on all realistic targets, so this is safe —
  just don't port the comparison to a 32-bit type without care.
- **Monotone pointer reset**: `split` must be reset to `i` at the start of each row. Sharing
  it across rows silently produces wrong answers on some inputs while still looking correct
  on the samples.
