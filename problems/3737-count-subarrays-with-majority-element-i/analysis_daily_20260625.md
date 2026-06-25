# 3737. Count Subarrays With Majority Element I

[LeetCode Link](https://leetcode.com/problems/count-subarrays-with-majority-element-i/)

Difficulty: Medium
Topics: Array, Hash Table, Divide and Conquer, Segment Tree, Merge Sort, Counting, Prefix Sum
Acceptance Rate: 70.4%

## Hints

### Hint 1

Only one thing about each element matters for this problem: is it equal to `target` or not? The actual values are irrelevant once you know that. Try mentally relabeling the array so every element is either a "good" element (equals `target`) or a "bad" element (everything else). What is the condition that makes `target` the majority of a subarray, expressed only in terms of how many good vs. bad elements it contains?

### Hint 2

"`target` appears strictly more than half the time" is the same as "the number of good elements is strictly greater than the number of bad elements" in that subarray. So assign `+1` to every good element and `-1` to every bad element. The majority condition becomes: **the sum of the transformed subarray is strictly positive.** Now you are just counting subarrays whose sum is `> 0`. Does a familiar tool for "subarray sums" come to mind?

### Hint 3

Use prefix sums. Let `pre[0] = 0` and `pre[k] = pre[k-1] + (+1 or -1)`. A subarray `nums[i..j-1]` has positive transformed sum exactly when `pre[j] - pre[i] > 0`, i.e. `pre[j] > pre[i]` with `i < j`. So the answer is the number of index pairs `i < j` where the later prefix value is **strictly larger** than the earlier one. With `n <= 1000` a direct double loop works; for the optimal version, count these ordered pairs as you scan left to right using a Fenwick tree (BIT) keyed by prefix value.

## Approach

**Step 1 — Reduce to a sign problem.** Replace each element with `+1` if it equals `target` and `-1` otherwise. `target` is the strict majority of a subarray iff (count of `+1`) > (count of `-1`) in that subarray, which is precisely (transformed sum) `> 0`.

**Step 2 — Prefix sums.** Define `pre[0] = 0` and `pre[k] = pre[k-1] + sign(nums[k-1])`. The transformed sum of `nums[i..j-1]` equals `pre[j] - pre[i]`. We want to count pairs `0 <= i < j <= n` with `pre[j] - pre[i] > 0`, equivalently `pre[i] < pre[j]`.

**Step 3 — Count ordered increasing pairs.** Scan `j` from `1` to `n`. Before processing `j`, all prefixes `pre[0..j-1]` have been inserted into a Fenwick tree that stores how many times each prefix value has appeared. For the current `pre[j]`, the number of earlier prefixes strictly smaller than it is a prefix-count query on the BIT; add that to the answer, then insert `pre[j]`.

Prefix values lie in `[-n, n]`, so shift them by `n` to get non-negative indices `[0, 2n]` and use a 1-indexed Fenwick tree of size `2n + 2`.

**Worked example:** `nums = [1,2,2,3]`, `target = 2`. Signs are `[-1, +1, +1, -1]`, so `pre = [0, -1, 0, 1, 0]`. Counting, for each `j`, how many earlier prefixes are strictly smaller:
- `j=1` (`pre=-1`): earlier `{0}`, smaller count `0`.
- `j=2` (`pre=0`): earlier `{0,-1}`, smaller `{-1}` → `1`.
- `j=3` (`pre=1`): earlier `{0,-1,0}`, smaller all three → `3`.
- `j=4` (`pre=0`): earlier `{0,-1,0,1}`, smaller `{-1}` → `1`.

Total `0 + 1 + 3 + 1 = 5`, matching the expected output.

Because the count of subarrays can be large (up to `~n^2/2`), accumulate the answer in a 64-bit integer.

## Complexity Analysis

Time Complexity: O(n log n) — one left-to-right scan with an `O(log n)` Fenwick query and update per index. (A plain double loop would be `O(n^2)`, which also passes given `n <= 1000`.)
Space Complexity: O(n) — the prefix array implicitly plus a Fenwick tree of size `O(n)`.

## Edge Cases

- **`target` absent from `nums` (Example 3):** every sign is `-1`, all prefixes are strictly decreasing, so no pair satisfies `pre[i] < pre[j]` and the answer is `0`.
- **All elements equal `target` (Example 2):** every sign is `+1`, prefixes strictly increase, so every one of the `n*(n+1)/2` subarrays counts.
- **Single-element subarrays:** an element equal to `target` is its own majority (sum `+1 > 0`); a non-target single element is not. The prefix formulation handles this automatically.
- **Negative prefix values:** prefix sums can go as low as `-n`, so indices must be shifted before indexing the Fenwick tree to avoid negative/out-of-range access.
- **Large counts:** the result may exceed what feels small at a glance; use `int64`-capable accumulation (Go's `int` is 64-bit on common platforms, but returning an `int` that fits is fine since `n <= 1000` keeps the max around `500500`).
