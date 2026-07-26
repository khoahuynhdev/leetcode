# 0628. Maximum Product of Three Numbers

[LeetCode Link](https://leetcode.com/problems/maximum-product-of-three-numbers/)

Difficulty: Easy
Topics: Array, Math, Sorting
Acceptance Rate: 46.4%

## Hints

### Hint 1

The brute force is three nested loops over every triple — O(n^3), far too slow for `n = 10^4`. But notice you are not asked *which* three numbers, only the product. That means you never need to inspect most of the array: only a handful of numbers near the extremes of the sorted order can possibly participate in the answer. Ask yourself which positions those are.

### Hint 2

Sorting makes the candidates visible. After sorting, the three largest values sit at the far right. That is clearly one candidate. But the array can contain negatives, and negatives are exactly what makes this problem more than "take the top three." Think about what happens to a product's sign when you multiply two negative numbers together.

### Hint 3

Two negatives multiply to a positive, and the *most negative* pair produces the *largest* positive magnitude. So the answer is always one of exactly two candidates:

- `nums[n-1] * nums[n-2] * nums[n-3]` — the three largest, and
- `nums[0] * nums[1] * nums[n-1]` — the two smallest (most negative) paired with the single largest.

Take the maximum of the two. Any other combination is dominated: mixing one negative with two positives gives a negative result, and pairing the two smallest with anything other than the largest positive value gives a smaller product. Once you see that only five numbers matter — the three largest and the two smallest — you can also drop the sort entirely and find them in one linear pass.

## Approach

The answer is determined by at most five values: the three largest and the two smallest elements. The reasoning is a short case analysis on how many negative numbers the winning triple contains.

A product of three numbers is positive when it has zero or two negative factors, and negative when it has one or three. To maximize:

- **Zero negatives in the triple.** Then we want the three largest values overall, giving `max1 * max2 * max3`.
- **Two negatives in the triple.** The two negative factors multiply to a positive, and to make that positive as large as possible we pick the two *most negative* numbers — the two smallest, `min1` and `min2`. The remaining factor must then be the single largest value, `max1`.
- **One or three negatives.** The product is negative. If any non-negative triple exists at all, the two cases above beat it. If the whole array is negative (so every triple is negative), we want the product closest to zero, which is the three largest — again `max1 * max2 * max3`, since those are the three negatives of smallest magnitude. Case 1 already covers it.

So `answer = max(max1*max2*max3, min1*min2*max1)`, and both branches are always safe to evaluate because the constraint `n >= 3` guarantees all five slots exist (values may repeat across slots only when `n` is small, but the two expressions still only reference distinct array positions).

The straightforward implementation sorts and reads the endpoints. The version in `solution.go` instead walks the array once, maintaining the running top three (`max1 >= max2 >= max3`) and bottom two (`min1 <= min2`). For each value `v`, cascade it down the max chain — if `v > max1`, shift `max1` to `max2` and `max2` to `max3` before assigning; otherwise try `max2`, then `max3`. Do the mirrored thing for the minimums. Initialize the maxes to a value below any legal input and the mins to a value above it, so the first few elements settle in naturally.

Worked example on `[-100, -98, -1, 2, 3, 4]`. After the scan, `max1, max2, max3 = 4, 3, 2` and `min1, min2 = -100, -98`. The two candidates are `4 * 3 * 2 = 24` and `(-100) * (-98) * 4 = 39200`. The answer is `39200` — the negative pair wins decisively, which is exactly the trap that the naive "sort and take the top three" solution falls into.

## Complexity Analysis

Time Complexity: O(n) — a single pass with constant work per element. (The sort-based variant is O(n log n).)
Space Complexity: O(1) — five scalar accumulators, regardless of input size.

## Edge Cases

- **Two large negatives plus positives**, e.g. `[-100, -98, -1, 2, 3, 4]`. The headline case. Skipping the `min1 * min2 * max1` candidate returns `24` instead of `39200`.
- **All negative numbers**, e.g. `[-1, -2, -3]`. Every product is negative, so the answer `-6` is the one closest to zero. This falls out of the three-largest branch without special handling.
- **Exactly three elements**, e.g. `[1, 2, 3]`. Both candidate expressions reduce to the same product here, so no guard is needed — but it does mean your min/max trackers must be seeded loosely enough to fill all five slots from just three inputs.
- **Zeros present**, e.g. `[-1, 0, 1, 2]`. Zero can legitimately be part of the maximum when everything else would produce a negative product. No special-casing required; the max of the two candidates handles it.
- **Duplicate values**, e.g. `[1, 1, 1]` or `[-1, -1, -1, -1]`. Use strict-vs-non-strict comparisons carefully in the cascade so equal elements still occupy separate slots rather than collapsing into one.
- **Integer range**. With `|nums[i]| <= 1000`, the largest magnitude product is `10^9`, which fits in a 32-bit int and comfortably in Go's `int` on any platform — no overflow concerns, but worth confirming rather than assuming.
- **Sentinel initialization**. If you seed the max trackers with `0` instead of something below `-1000`, an all-negative array silently returns `0`. Use `math.MinInt` / `math.MaxInt` or bounds derived from the constraints.
