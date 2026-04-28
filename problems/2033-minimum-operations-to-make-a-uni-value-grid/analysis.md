# 2033. Minimum Operations to Make a Uni-Value Grid

[LeetCode Link](https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/)

Difficulty: Medium
Topics: Array, Math, Sorting, Matrix
Acceptance Rate: 68.2%

## Hints

### Hint 1

The grid's 2D shape is a distraction. Every operation only touches a single cell, and the cells don't interact with each other, so flatten the grid into a 1D list and reason about that list. Now ask: under what condition can two numbers `a` and `b` ever become equal using only `+x` and `-x` steps?

### Hint 2

Two values `a` and `b` can be made equal with `±x` operations only when `(a - b)` is divisible by `x` — equivalently, when `a % x == b % x`. So first, sanity-check feasibility: if any two flattened values disagree on their remainder mod `x`, the answer is `-1`. Otherwise, every value reaches some shared target `T` in exactly `|a - T| / x` operations, and you want to choose `T` to minimize the total.

### Hint 3

You are minimizing `Σ |a_i - T|` (then dividing by `x`). This is the classic "minimize sum of absolute deviations" problem, and the minimizer is the **median** of the values. Sort the flattened array, pick the middle element as the target, then sum `(|a_i - median|) / x`. Note the target must be one of the existing values (since they all share the same residue mod `x`, and the median is one of them — so feasibility is automatic).

## Approach

1. **Flatten** the `m x n` grid into a single slice of length `m * n`. The 2D structure adds nothing — operations are independent per cell.
2. **Feasibility check**: pick any reference value (e.g. the first cell) and compute `r = grid[0][0] % x`. Every other cell must satisfy `cell % x == r`. If not, return `-1`. This works because `±x` operations preserve the value mod `x`.
3. **Find the median**: sort the flattened slice (or use quickselect). The median minimizes the sum of absolute distances to all points — a standard result. Because the median is one of the actual grid values, it automatically shares the common residue, so it is a reachable target.
4. **Count operations**: for each value `v`, the number of operations needed is `|v - median| / x`. Sum these up and return.

**Worked example** (Example 1): `grid = [[2,4],[6,8]]`, `x = 2`.
- Flatten → `[2, 4, 6, 8]`. All have remainder `0 mod 2`. Feasible.
- Sort → `[2, 4, 6, 8]`. Median (lower middle) = `4`.
- Operations: `|2-4|/2 + |4-4|/2 + |6-4|/2 + |8-4|/2 = 1 + 0 + 1 + 2 = 4`. ✓

**Why the median, intuitively?** If you're standing at point `T` on a number line and shift right by 1, you move closer to every point on your right and farther from every point on your left. The total distance decreases as long as more points are on the right, and increases once more are on the left. The break-even is the median.

## Complexity Analysis

Let `N = m * n`.

Time Complexity: O(N log N) — dominated by sorting the flattened array. Quickselect would bring this to O(N) average, but sorting is simple and the constraints (`N ≤ 10^5`) make it fast enough.
Space Complexity: O(N) — for the flattened slice. (O(1) extra if you sort in place and view the grid as a 1D buffer.)

## Edge Cases

- **Single element grid (`1 x 1`)**: already uni-value, answer is `0`. The algorithm handles this naturally — median equals the only element, sum is `0`.
- **Infeasible grid**: any two cells with different residues mod `x` → return `-1`. Don't forget this check; otherwise you'd compute a nonsensical positive answer.
- **`x = 1`**: every grid is feasible because every integer is `0 mod 1`... wait, every integer has *some* remainder mod 1 (namely `0`), so all residues match. The algorithm still works.
- **All elements already equal**: median equals each value, answer is `0`. Handled implicitly.
- **Even-length array (median ambiguity)**: any value between the two middle elements minimizes the sum, but using the lower-middle (the conventional `arr[N/2]` for 0-indexed sort) gives the same total cost and is always one of the grid values, keeping feasibility safe.
- **Large values / large `N`**: use `int` (64-bit on most platforms) for the running sum to avoid overflow worries. With `N ≤ 10^5` and values `≤ 10^4`, the sum fits comfortably in int64.
