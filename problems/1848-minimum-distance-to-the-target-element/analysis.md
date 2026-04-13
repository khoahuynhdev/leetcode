# 1848. Minimum Distance to the Target Element

[LeetCode Link](https://leetcode.com/problems/minimum-distance-to-the-target-element/)

Difficulty: Easy
Topics: Array
Acceptance Rate: 57.6%

## Hints

### Hint 1

This is a classic "scan the array and track a best value" problem. Think about what information you need to maintain as you walk through the array, and what you need to compute for each element you encounter.

### Hint 2

You only care about positions `i` where `nums[i] == target`. For every such position, you can compute a candidate answer `abs(i - start)`. Your job is to keep the minimum such candidate as you iterate.

### Hint 3

A single linear pass is enough. Initialize a "best so far" with something larger than any possible distance (e.g. `len(nums)`), then for each index `i` check whether `nums[i] == target`; if it does, update the best with `min(best, abs(i - start))`. The problem guarantees `target` exists, so you will always update at least once.

## Approach

Walk through `nums` once with index `i`:

1. Skip indices where `nums[i] != target`.
2. For indices where `nums[i] == target`, compute `d = abs(i - start)`.
3. Maintain the smallest `d` seen so far in a variable `best`.

Since `target` is guaranteed to appear in `nums`, `best` will be updated at least once, and we simply return it at the end.

**Why it works.** The answer is the minimum of a finite set of candidates `{abs(i - start) : nums[i] == target}`. A linear pass that takes the minimum of this set is both correct and optimal — we cannot decide the answer without inspecting every element at least once in the worst case (e.g., the matching index could be the last one, or a closer match could appear anywhere).

**Example walkthrough.** `nums = [1,2,3,4,5]`, `target = 5`, `start = 3`.
- `i=0..3`: no match.
- `i=4`: `nums[4] == 5`, `abs(4 - 3) = 1`, update `best = 1`.
- Return `1`.

A small optimization: because indices move away from `start` in both directions, once you are past `start` you could early-exit as soon as the distance starts to exceed an already-found match on the other side. For `n ≤ 1000`, though, the plain linear scan is already trivially fast and much clearer.

## Complexity Analysis

Time Complexity: O(n), where n = len(nums). We inspect each element once.
Space Complexity: O(1). Only a single integer `best` is tracked.

## Edge Cases

- **Single-element array** (`nums = [1]`, `start = 0`): the only index matches, return `0`.
- **`start` itself matches `target`**: distance is `0`, which is the global minimum — nothing can beat it.
- **Multiple matches**: we must pick the one closest to `start`, not the first or last (example 3 covers this with all-equal values).
- **Target at the far end from `start`**: ensures the initial `best` sentinel must be large enough; using `len(nums)` (strictly greater than any valid distance `≤ n-1`) works safely.
- **No match**: the problem guarantees `target` is present, so we don't need to handle "not found"; still, seeding `best` sensibly keeps the code robust.
