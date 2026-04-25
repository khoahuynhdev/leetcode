# 3464. Maximize the Distance Between Points on a Square

[LeetCode Link](https://leetcode.com/problems/maximize-the-distance-between-points-on-a-square/)

Difficulty: Hard
Topics: Array, Math, Binary Search, Geometry, Sorting
Acceptance Rate: 33.9%

## Hints

### Hint 1

The phrase "maximize the minimum" is the giveaway. Whenever the answer space
is monotonic (if a distance `d` is feasible, every smaller `d` is also
feasible), think binary search on the answer. The remaining question is:
given a target distance `d`, can we *check feasibility* — pick `k` points so
every pair is at least `d` apart — quickly?

### Hint 2

Points live on the boundary of the square, not anywhere in the plane. That
extra structure is what makes the problem tractable. Try parameterizing each
point by a single number: how far it sits along the perimeter, walking
clockwise from `(0, 0)`. Once every point is a number on a circle of length
`4 * side`, the problem starts to look like "place `k` items on a ring with a
minimum gap." That's a classic feasibility-check pattern.

### Hint 3

The crucial geometric fact: when the candidate distance `d` is at most `side`
(which it always is for `k >= 4`, because `k * d <= 4 * side` from the
pigeonhole on perimeter), then **perimeter distance >= d implies Manhattan
distance >= d**. For two points on the same or adjacent edges, the two
distances are equal; for two points on opposite edges, the Manhattan distance
is at least `side >= d` automatically. So the feasibility check collapses to:
sort by perimeter position, fix a starting point, and greedily pick the next
point whose perimeter position is at least `d` ahead. If you reach `k` picks
without overshooting one full lap, `d` is feasible.

## Approach

**Step 1 — Parameterize by perimeter position.** Walk clockwise from
`(0, 0)`. Map each point to a single number `t in [0, 4 * side)`:

- Bottom edge (`y == 0`): `t = x`
- Right edge (`x == side`): `t = side + y`
- Top edge (`y == side`): `t = 3*side - x`
- Left edge (`x == 0`): `t = 4*side - y`

Corners get a single canonical value because we check the edges in the order
above (so `(0, 0)` becomes `0`, never `4*side`).

**Step 2 — Sort by `t`.** Now the points lie on a ring of length
`P = 4 * side` and we want to select `k` of them so every cyclic gap is at
least `d`.

**Step 3 — Binary search on `d`.** Lower bound `0`, upper bound `P / k`
(by averaging cyclic gaps). For each candidate `d`, run a feasibility check.

**Step 4 — Feasibility check.** Duplicate the sorted perimeter array shifted
by `P` (the standard "double the array" trick for cyclic problems). For each
starting index `i` in `[0, n)`:

1. Set `last = perim[i]`.
2. Repeat `k - 1` times: binary-search for the smallest index `j` (greater
   than `i`) in the doubled array with `perim2[j] >= last + d`.
3. If `perim2[j] > perim[i] + P - d`, this start fails (the cyclic wrap-back
   gap would be smaller than `d`).
4. Otherwise update `last = perim2[j]` and continue.

If any starting point completes `k - 1` jumps without failing, `d` is
feasible. The doubled array lets us treat wrap-around naturally: a pick that
falls in the second copy represents a pick that wraps past the start.

**Why "perimeter gap >= d" suffices.** This is the heart of the problem.
For consecutive picks (in perimeter order):

- Same edge or adjacent edges: the perimeter path *is* the Manhattan path,
  so Manhattan distance equals the perimeter gap, which is `>= d`.
- Opposite edges: Manhattan distance is `|x1 - x2| + side >= side`, and
  since `k >= 4` forces `d <= 4*side / k <= side`, we get Manhattan
  distance `>= side >= d` for free.

For non-consecutive picks (i, j) with several picks between them, perimeter
distance is the sum of intermediate gaps `>= 2d`, so the same argument
gives Manhattan distance `>= d` either directly or via the opposite-edge
fallback. So consecutive-perimeter feasibility implies pairwise feasibility.

**Worked example (Example 1).** `side = 2`, points
`[[0,2],[2,0],[2,2],[0,0]]`, `k = 4`. Perimeter positions: `0, 2, 4, 6`
after sorting. With `d = 2`: starting at `0`, pick `2` (>= 0+2), then `4`
(>= 2+2), then `6` (>= 4+2). Wrap-back gap = `0 + 8 - 6 = 2 >= 2`. Feasible.
`d = 3` fails because three jumps of size 3 would land at `9 > 8 - 3 = 5`.
Answer = `2`.

## Complexity Analysis

Let `n` be the number of points and `V = 4 * side` the perimeter length.

Time Complexity: `O(n * k * log n * log V)`. For each of `O(log V)` binary
search rounds, the feasibility check tries `n` starting points and performs
`k - 1` binary searches of `O(log n)` each. With `n <= 15000`, `k <= 25`,
`log V <= 32`, this is well under `10^8` operations.

Space Complexity: `O(n)` for the sorted perimeter array and its doubled
copy.

## Edge Cases

- **Corner points.** `(0, 0)`, `(side, 0)`, `(side, side)`, `(0, side)` lie
  on two edges each. The perimeter mapping must pick one canonical value
  per corner; checking edges in the order bottom, right, top, left handles
  this without producing duplicates at `0` and `4*side`.
- **`k == 4` on a tiny square (`side == 1`).** Only the four corners exist,
  pairwise Manhattan distances are `1, 1, 1, 1, 2, 2`, so the answer is
  `1`. The upper-bound `P / k = 1` keeps the binary search tight.
- **`d == 0` is always feasible** when `n >= k`. Make sure the binary search
  starts with `lo = 0` and the feasibility check returns `true` for that.
- **Cyclic wrap.** Without the "doubled array" trick (or an equivalent
  modular comparison), a feasible-but-wrap-around configuration would be
  missed. The check `perim2[j] <= first + P - d` enforces the wrap-back
  gap explicitly.
- **Integer overflow.** With `side` up to `10^9`, perimeter values can
  reach `4 * 10^9`, and the doubled array reaches `8 * 10^9`. Use 64-bit
  integers throughout the perimeter math.
