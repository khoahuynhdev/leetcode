# 3903. Smallest Stable Index I

[LeetCode Link](https://leetcode.com/problems/smallest-stable-index-i/)

Difficulty: Easy
Topics: Array, Prefix Sum
Acceptance Rate: 72.5%

## Hints

### Hint 1

The score at index `i` asks two questions that point in opposite directions: something about everything *at or before* `i`, and something about everything *at or after* `i`. Whenever a per-index answer decomposes into "a fact about the prefix" and "a fact about the suffix", think about precomputing both sides once instead of recomputing them inside a loop.

### Hint 2

The constraints are tiny (`n <= 100`), so a brute-force double loop passes. But the interesting version is the linear one: the running maximum of `nums[0..i]` can be maintained with a single left-to-right sweep, since `max(nums[0..i]) = max(max(nums[0..i-1]), nums[i])`. Ask yourself what the mirror-image sweep gives you for `min(nums[i..n-1])`.

### Hint 3

Precompute a suffix-minimum array in one right-to-left pass: `suf[i] = min(nums[i], suf[i+1])`, with `suf[n-1] = nums[n-1]`. Then sweep left to right maintaining the running prefix maximum, and at each `i` the instability score is just `runningMax - suf[i]`. Return the first `i` where that is `<= k`.

One trap worth naming: both `prefMax` and `sufMin` are non-decreasing in `i`, so their difference is **not** monotonic — it can go down and back up. That kills any temptation to binary search for the answer; you genuinely need to scan from the left and stop at the first hit.

## Approach

Note that `max(nums[0..i])` and `min(nums[i..n-1])` are both *inclusive* of index `i`, so every index has a well-defined score and the two ranges overlap at exactly one element. In particular `max(nums[0..i]) >= nums[i] >= min(nums[i..n-1])`, so the instability score is always non-negative.

The naive approach recomputes both extremes for each `i`, costing O(n) per index and O(n²) overall. That is fine for `n <= 100`, but the decomposition into prefix/suffix aggregates gives an O(n) algorithm with almost no extra code.

The algorithm:

1. Handle the trivial empty input by returning `-1` (the constraints guarantee `n >= 1`, but a defensive guard keeps the slice indexing below safe).
2. Build `sufMin` of length `n` in a right-to-left pass: `sufMin[n-1] = nums[n-1]`, and for `i` from `n-2` down to `0`, `sufMin[i] = min(nums[i], sufMin[i+1])`. After this pass, `sufMin[i]` is exactly `min(nums[i..n-1])`.
3. Sweep left to right with a running `prefMax`, initialized so that the first update sets it to `nums[0]`. At index `i`, first fold `nums[i]` into `prefMax` — now `prefMax == max(nums[0..i])` — then test `prefMax - sufMin[i] <= k`.
4. Return the first `i` that passes. If the loop finishes with no hit, return `-1`.

Because we scan indices in increasing order and return immediately, the first success is by construction the *smallest* stable index.

Walking through Example 1, `nums = [5,0,1,4]`, `k = 3`:

| i | nums[i] | prefMax | sufMin[i] | score | <= 3? |
|---|---------|---------|-----------|-------|-------|
| 0 | 5       | 5       | 0         | 5     | no    |
| 1 | 0       | 5       | 0         | 5     | no    |
| 2 | 1       | 5       | 1         | 4     | no    |
| 3 | 4       | 5       | 4         | 1     | yes   |

The scan stops at `i = 3` and returns 3, matching the expected output. Notice the score sequence `5, 5, 4, 1` is decreasing here, but that is a property of this input, not of the problem — for `nums = [0, 5, 0]` the scores are `0, 5, 5`, which is why the left-to-right scan (rather than any search) is the right tool.

The prefix maximum is folded in *before* the comparison rather than after, which is the one place an off-by-one can sneak in: forget it and index 0 would be tested against an uninitialized maximum.

## Complexity Analysis

Time Complexity: O(n) — one right-to-left pass to build the suffix minima, and at most one left-to-right pass to find the answer.
Space Complexity: O(n) — the `sufMin` array. This can be reduced to O(1) auxiliary space only by giving up the single-pass structure (e.g. recomputing suffix minima per index, back to O(n²) time), so O(n) space is the right trade here.

## Edge Cases

- **Single element (`n == 1`)**: `max(nums[0..0]) == min(nums[0..0]) == nums[0]`, so the score is 0. Since `k >= 0`, the answer is always 0. Example 3 is exactly this case.
- **No stable index at all**: as in Example 2, the loop must fall through and return `-1` rather than returning a sentinel like `n` or leaving a zero value.
- **Index 0 and index `n-1` are not special-cased**: at `i = 0` the score is `nums[0] - min(nums)`, and at `i = n-1` it is `max(nums) - nums[n-1]`. Both fall out of the general formula, so no separate branches are needed — but the `sufMin` construction must seed `sufMin[n-1] = nums[n-1]` rather than starting from a "+infinity" placed at index `n-1`.
- **All elements equal**: every score is 0, so the answer is 0 for any `k`. A good sanity check that the ranges are inclusive on both sides.
- **`k == 0`**: only indices where the prefix max equals the suffix min qualify — i.e. `nums` is non-increasing up to `i` and non-decreasing after it, in the sense that `nums[i]` is simultaneously the largest so far and the smallest remaining. Strict `<=` matters; using `<` would break Example 3.
- **Large values**: `nums[i]` and `k` go up to 10⁹. The difference stays well inside Go's `int` (64-bit on all supported platforms, and even 32-bit `int` would hold 10⁹), so no overflow handling is needed — but note the score is a difference of two values, not a sum, which is what keeps it safe.
- **Defensive empty slice**: the constraints forbid it, but returning `-1` up front avoids an out-of-range panic on `nums[n-1]` if the function is ever reused outside the judge.
