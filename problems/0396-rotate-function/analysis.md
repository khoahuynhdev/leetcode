# 0396. Rotate Function

[LeetCode Link](https://leetcode.com/problems/rotate-function/)

Difficulty: Medium
Topics: Array, Math, Dynamic Programming
Acceptance Rate: 47.3%

## Hints

### Hint 1

A naive simulation rotates the array `n` times and recomputes `F(k)` from scratch for every rotation, giving an `O(n^2)` algorithm. With `n` up to `10^5`, that is too slow. Look at the structure of consecutive `F` values: there is a clean algebraic relationship between `F(k-1)` and `F(k)` that lets you avoid recomputing the entire sum each time.

### Hint 2

Write out `F(k)` and `F(k-1)` side by side and subtract them term by term. Most of the products differ only by a coefficient of `+1`, except for one element whose coefficient drops sharply. Once you spot that pattern, you only need two precomputed quantities to advance from `F(k-1)` to `F(k)` in `O(1)`.

### Hint 3

The identity is `F(k) = F(k-1) + sum(nums) - n * nums[n - k]`. Compute `sum(nums)` and `F(0)` once, then sweep `k` from `1` to `n-1` updating `F` in constant time and tracking the running maximum. The whole algorithm is one pass after the initial setup.

## Approach

Define `S = sum(nums)` and start from
`F(0) = 0 * nums[0] + 1 * nums[1] + ... + (n-1) * nums[n-1]`.

When we move from `F(k-1)` to `F(k)`, every element of `nums` keeps its value but slides one position to the right (in the rotated view), so its coefficient in the dot product increases by `1` — except for `nums[n - k]`, which wraps around from coefficient `n-1` to coefficient `0`. That gives the recurrence:

```
F(k) = F(k-1) + S - n * nums[n - k]
```

Verifying with `nums = [4, 3, 2, 6]`, `n = 4`, `S = 15`:

- `F(0) = 0 + 3 + 4 + 18 = 25`
- `F(1) = 25 + 15 - 4*6 = 16`
- `F(2) = 16 + 15 - 4*2 = 23`
- `F(3) = 23 + 15 - 4*3 = 26`

The maximum is `26`, matching the expected output.

The algorithm:

1. Compute `S` and `F(0)` in one pass.
2. Initialize `best = F(0)` and `f = F(0)`.
3. For `k = 1, 2, ..., n-1`: update `f += S - n * nums[n-k]` and `best = max(best, f)`.
4. Return `best`.

## Complexity Analysis

Time Complexity: O(n) — two linear sweeps over `nums`.
Space Complexity: O(1) — only a handful of integer accumulators.

## Edge Cases

- **Single element (`n == 1`)**: every rotation yields `F(k) = 0 * nums[0] = 0`. The recurrence loop does not execute, and the initial `F(0) = 0` is returned correctly.
- **All zeros**: every `F(k)` is `0`; the maximum is `0`.
- **Negative values**: the answer can be negative (e.g. `nums = [-1, -2, -3]`); seeding `best` with `F(0)` rather than `0` or `math.MinInt` avoids both wrong-answer and overflow issues.
- **Mixed signs**: the optimal rotation is not always `k = 0` or `k = n-1`; the sweep over all `k` is necessary.
- **Large `n` (up to `10^5`)**: the `O(n^2)` simulation would TLE; the recurrence keeps the work linear. The 32-bit guarantee on the answer means `int` is comfortably wide enough in Go.
