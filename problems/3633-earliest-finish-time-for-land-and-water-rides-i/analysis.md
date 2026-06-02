# 3633. Earliest Finish Time for Land and Water Rides I

[LeetCode Link](https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-i/)

Difficulty: Easy
Topics: Array, Two Pointers, Binary Search, Greedy, Sorting
Acceptance Rate: 66.9%

## Hints

### Hint 1

The tourist must take exactly one land ride and one water ride, in either order. There are only two "shapes" of plan: land-then-water, or water-then-land. Think about how to evaluate each shape independently, then combine.

### Hint 2

For a fixed ordering (say land first, then water), the finish time depends on a chosen land ride `i` and water ride `j`. Write down the closed-form expression for that finish time and ask yourself: which `i` and `j` minimize it? The `max` of two values (when one is held fixed) is monotonic — that observation lets you decouple the two choices.

### Hint 3

For the land-then-water plan, the finish time is `max(landStart[i] + landDuration[i], waterStart[j]) + waterDuration[j]`. For any fixed `j`, this expression is minimized by picking the `i` that yields the smallest land finish time. So you only need the **minimum land finish time** across all `i`, then sweep over `j`. The water-then-land plan is symmetric. Answer is the min of the two plans.

## Approach

Let `landFinish[i] = landStartTime[i] + landDuration[i]` and `waterFinish[j] = waterStartTime[j] + waterDuration[j]`.

For each pair `(i, j)` we have two candidate finish times:

- **Land then water:** `max(landFinish[i], waterStartTime[j]) + waterDuration[j]`
- **Water then land:** `max(waterFinish[j], landStartTime[i]) + landDuration[i]`

The answer is the minimum of all these candidates.

Naively this is `O(n * m)`, which is fine given the small constraints (`n, m <= 100`). But we can do better by separating the two choices:

- For the **land-then-water** plan, fix `j`. The expression depends on `i` only through `landFinish[i]`, and `max(., waterStartTime[j])` is non-decreasing in its first argument. So the optimal `i` is the one minimizing `landFinish[i]`. Compute `minLandFinish = min_i landFinish[i]` once, then evaluate `max(minLandFinish, waterStartTime[j]) + waterDuration[j]` for every `j` and take the min.
- For the **water-then-land** plan, symmetrically use `minWaterFinish = min_j waterFinish[j]`, then minimize `max(minWaterFinish, landStartTime[i]) + landDuration[i]` over `i`.

Return the minimum of the two best plans.

**Walking through example 1** (`landStartTime = [2,8], landDuration = [4,1], waterStartTime = [6], waterDuration = [3]`):
- `landFinish = [6, 9]`, so `minLandFinish = 6`.
- `waterFinish = [9]`, so `minWaterFinish = 9`.
- Land-then-water best: `max(6, 6) + 3 = 9`.
- Water-then-land best: `min(max(9, 2) + 4, max(9, 8) + 1) = min(13, 10) = 10`.
- Answer: `min(9, 10) = 9`. ✓

## Complexity Analysis

Time Complexity: O(n + m) — one pass to find the minimum finish time in each category, plus one pass over the other category.
Space Complexity: O(1) — only a handful of scalars are tracked.

## Edge Cases

- **Single ride in each category (`n = m = 1`):** Both plans reduce to one candidate each; we still return the min of the two. Example 2 is exactly this case.
- **One category opens much later than the other finishes:** The "wait until open" behavior is captured by `max(prevFinish, otherStartTime)`. Make sure not to start a ride before its `startTime`.
- **All rides in one category share the same start time:** The `min` over that category just picks the shortest duration; nothing special is required.
- **Large durations vs. small start times:** Make sure intermediate sums fit in the integer type. With constraints up to 1000 per value and at most 100 rides per side, a finish time stays well within `int` range.
- **Optimal plan can mix rides differently in each ordering:** The land ride that minimizes the land-then-water plan need not be the same as the one that minimizes the water-then-land plan — that's why we compute each plan independently.
