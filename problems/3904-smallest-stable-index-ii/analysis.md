# 3904. Smallest Stable Index II

[LeetCode Link](https://leetcode.com/problems/smallest-stable-index-ii/)

Difficulty: Medium
Topics: Array, Prefix Sum
Acceptance Rate: 75.2%

## Hints

### Hint 1

The brute-force reading of the problem is "for each index `i`, scan left for a max and scan right for a min," which is O(n²). But notice that as `i` moves from left to right, the left window only ever *grows* and the right window only ever *shrinks*. Whenever a quantity over a growing or shrinking window is asked for at *every* index, think about precomputing it in one sweep rather than recomputing it from scratch.

### Hint 2

This is the prefix/suffix aggregate pattern (the same family as prefix sums, just with `max`/`min` instead of `+`). Define two arrays:

- `prefMax[i] = max(nums[0..i])`
- `sufMin[i] = min(nums[i..n-1])`

Each of these can be built in a single pass, because `prefMax[i]` depends only on `prefMax[i-1]` and `nums[i]`, and `sufMin[i]` depends only on `sufMin[i+1]` and `nums[i]`. Once you have both, the instability score at `i` is just one subtraction.

### Hint 3

The critical realization is that you should **not** try to binary search the answer. Both `prefMax` and `sufMin` are non-decreasing, so their difference is *not* monotonic — a stable index can be followed by an unstable one. For example, `nums = [2, 1, 3]` gives scores `1, 1, 0`, while `nums = [1, 3, 2]` gives scores `0, 1, 1`. So the predicate "index `i` is stable" has no monotone structure to exploit, and you must scan left to right and return the first index that satisfies it.

The other half of the insight is a space trick: since you walk left to right, `prefMax` never needs to be stored as an array — a single running variable suffices. Only `sufMin` has to be materialized, because it is built in the opposite direction.

## Approach

The whole problem reduces to evaluating `prefMax[i] - sufMin[i] <= k` for every `i` in increasing order and returning the first `i` that passes.

**Step 1 — build the suffix minimum array.** Walk from the right end to the left:

```
sufMin[n-1] = nums[n-1]
sufMin[i]   = min(nums[i], sufMin[i+1])   for i from n-2 down to 0
```

After this pass, `sufMin[i]` holds the smallest value in `nums[i..n-1]`. This costs O(n) time and O(n) space.

**Step 2 — sweep left to right with a running prefix maximum.** Maintain a variable `prefMax` initialized to `nums[0]`; before testing index `i`, fold `nums[i]` into it:

```
prefMax = max(prefMax, nums[i])
if prefMax - sufMin[i] <= k:
    return i
```

The first index that satisfies the condition is by construction the smallest stable index. If the loop finishes without a hit, return `-1`.

**Worked example** on `nums = [5, 0, 1, 4]`, `k = 3`:

| i | nums[i] | prefMax | sufMin[i] | score | stable? |
|---|---------|---------|-----------|-------|---------|
| 0 | 5       | 5       | 0         | 5     | no      |
| 1 | 0       | 5       | 0         | 5     | no      |
| 2 | 1       | 5       | 1         | 4     | no      |
| 3 | 4       | 5       | 4         | 1     | **yes** |

We return `3`, matching the expected output. Note how index 3 becomes stable only at the very end — a good reminder of why the scan cannot stop early or binary search.

For `nums = [3, 2, 1]`, `k = 1`, the scores are `2, 2, 2`; nothing is `<= 1`, so the answer is `-1`.

Two details worth calling out. First, both endpoints are *inclusive* and the two ranges *overlap* at `i` itself, so `nums[i]` participates in both the max and the min. That guarantees the score is always `>= 0`, since `prefMax[i] >= nums[i] >= sufMin[i]`. Second, this means `i = n-1` has score `max(nums) - nums[n-1]` and `i = 0` has score `nums[0] - min(nums)` — no index is special-cased.

This problem is genuinely on the easier side for a Medium (hence the ~75% acceptance rate), but it is a clean, honest exercise in the prefix/suffix precomputation pattern, and the "don't binary search this" trap is a real one worth internalizing.

## Complexity Analysis

Time Complexity: O(n) — one right-to-left pass to build `sufMin`, plus at most one left-to-right pass.
Space Complexity: O(n) — for the `sufMin` array. The prefix maximum needs only O(1) extra space because it is consumed in the same direction it is produced.

## Edge Cases

- **Single element (`n == 1`).** The score is `nums[0] - nums[0] = 0`, which is `<= k` for any `k >= 0`, so the answer is always `0`. The suffix-min base case must be initialized from `nums[n-1]` rather than from a sentinel to make this fall out naturally.
- **No stable index at all.** A strictly decreasing array such as `[3, 2, 1]` with small `k` yields a constant score of `max - min` at every index. The function must return `-1` rather than defaulting to `0` or `n-1`.
- **`k = 0`.** Requires an exact match `prefMax[i] == sufMin[i]`. A non-decreasing array like `[1, 2, 3, 4]` satisfies this at every index (answer `0`), whereas `[2, 1]` satisfies it nowhere.
- **All elements equal.** Every score is `0`, so the answer is `0` for any valid `k`. A useful sanity check that the two windows overlapping at `i` is handled correctly.
- **Large values.** With `nums[i]` up to `10^9`, the difference fits comfortably in `int32`, but Go's `int` is 64-bit on all target platforms, so no overflow concern arises. The subtraction never goes negative, as argued above.
- **Answer only at the last index.** As in Example 1, a large leading element can keep every early index unstable. Make sure the loop actually runs through `i = n-1` and does not stop one short.
- **Non-monotone stability.** Inputs like `[1, 3, 2]` (scores `0, 1, 1`) prove a stable index can be followed by an unstable one; do not try to short-circuit the scan based on an assumed trend.
