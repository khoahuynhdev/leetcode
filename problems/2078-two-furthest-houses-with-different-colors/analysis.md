# 2078. Two Furthest Houses With Different Colors

[LeetCode Link](https://leetcode.com/problems/two-furthest-houses-with-different-colors/)

Difficulty: Easy
Topics: Array, Greedy
Acceptance Rate: 67.2%

## Hints

### Hint 1

The brute force answer is an `O(n^2)` scan of every pair — perfectly fine for `n <= 100`, but not very interesting. Before you reach for two nested loops, ask yourself: *which pairs of indices can even possibly produce the maximum distance?* The structure of "maximum distance" in a linear array is very restrictive.

### Hint 2

Think about where the optimal pair must be anchored. If the answer pair is `(i, j)` with `i < j`, can both `i` and `j` sit in the interior (neither touches index `0` nor `n-1`)? Play with a small example: replace `i` with `0` and ask whether that still gives a valid "different colors" pair. If yes, you just improved (or tied) the distance. That observation tells you one endpoint is effectively fixed.

### Hint 3

At least one endpoint of the optimal answer is the first house or the last house. So you only need to consider two candidates:

1. Fix the left end at index `0` and sweep from the right to find the furthest index whose color differs from `colors[0]`.
2. Fix the right end at index `n-1` and sweep from the left to find the furthest index whose color differs from `colors[n-1]`.

The answer is the larger of those two distances. This reduces the problem to a single linear pass.

## Approach

**Key observation.** The maximum distance between two houses with different colors must involve either the first house (`index 0`) or the last house (`index n-1`).

Why? Consider any valid pair `(i, j)` with `i < j` and `colors[i] != colors[j]`. Look at `colors[0]`:

- If `colors[0] != colors[j]`, then `(0, j)` is a valid pair and `j - 0 >= j - i`, so it's at least as good.
- Otherwise `colors[0] == colors[j]`, which means `colors[0] != colors[i]` (since `colors[i] != colors[j]`). Now look at `colors[n-1]`:
  - If `colors[n-1] != colors[i]`, then `(i, n-1)` is a valid pair and `(n-1) - i >= j - i`, at least as good.
  - Otherwise `colors[n-1] == colors[i] != colors[0]`, so `(0, n-1)` itself is valid, with distance `n-1` — the global maximum.

In every case, some pair anchored at `0` or `n-1` matches or beats `(i, j)`. So two sweeps suffice.

**Algorithm.**

1. Sweep from the right end inward: find the largest `j` such that `colors[j] != colors[0]`. Call that distance `d1 = j`.
2. Sweep from the left end inward: find the smallest `i` such that `colors[i] != colors[n-1]`. Call that distance `d2 = (n-1) - i`.
3. Return `max(d1, d2)`.

The problem guarantees at least two different colors exist, so both sweeps are guaranteed to find a valid index.

**Example walkthrough** — `colors = [1, 8, 3, 8, 3]`:

- `colors[0] = 1`. Scan right-to-left: `colors[4] = 3 != 1`, so `d1 = 4`.
- `colors[4] = 3`. Scan left-to-right: `colors[0] = 1 != 3`, so `d2 = 4 - 0 = 4`.
- Answer: `max(4, 4) = 4`. ✓

## Complexity Analysis

Time Complexity: O(n) — two single passes over the array in the worst case.
Space Complexity: O(1) — only a handful of indices are tracked.

## Edge Cases

- **Minimum length (`n = 2`).** The two houses must have different colors (guaranteed by constraints), so the answer is `1`. Both sweeps will immediately find the opposite endpoint.
- **All-but-one same color, odd color at an endpoint.** Example: `[6, 1, 1, 1, 1]`. Fixing the right end fails to find a different color until it reaches index `0`, giving distance `4`. Fixing the left end finds a different color at index `1`, giving distance `1`. The right-anchored sweep wins — showing why both sweeps are necessary.
- **All-but-one same color, odd color in the middle.** Example: `[1, 1, 6, 1, 1]`. Neither endpoint anchors the optimal pair directly on its own, but the observation still holds: `colors[0] == colors[n-1]`, and the odd middle index `2` is different from both. `d1 = 2` (distance from index 0 to the middle), `d2 = 2` (distance from middle to `n-1`), answer is `2`.
- **First and last colors already differ.** Example: `[0, 1]` or `[1, 2, 2, 2, 3]`. Both sweeps immediately return `n-1`, the global maximum.
