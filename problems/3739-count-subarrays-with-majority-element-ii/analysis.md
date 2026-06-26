# 3739. Count Subarrays With Majority Element II

[LeetCode Link](https://leetcode.com/problems/count-subarrays-with-majority-element-ii/)

Difficulty: Hard
Topics: Array, Hash Table, Divide and Conquer, Segment Tree, Merge Sort, Prefix Sum
Acceptance Rate: 58.2%

## Hints

### Hint 1

The actual values in `nums` don't matter — only whether each element *is* or *is not* the `target`. Whenever a problem cares about "more than half" of something, try collapsing the data into two categories (a `+1` / `-1` world) so you can reason about a running balance instead of raw counts. What does "majority" become once you do that?

### Hint 2

Suppose you map every element equal to `target` to `+1` and every other element to `-1`. For a subarray, `target` is the strict majority exactly when the number of `+1`s exceeds the number of `-1`s — that is, when the **sum of the transformed subarray is strictly positive**. Now you're counting subarrays with a positive sum, which is a classic prefix-sum setup. Be careful: counting subarrays one-by-one is `O(n^2)`, and `n` can be `10^5`.

### Hint 3

Build prefix sums `P[0..n]` of the transformed array. A subarray `(l..r)` has positive sum iff `P[r+1] - P[l] > 0`, i.e. `P[l] < P[r+1]`. So the answer is simply the number of index pairs `(i, j)` with `i < j` and `P[i] < P[j]`. That is a "count pairs in order" problem — sweep left to right and, for each new prefix value, ask "how many earlier prefix values are strictly smaller?" A Fenwick tree (BIT) or a merge-sort inversion count answers every query in `O(log n)`.

## Approach

**1. Reduce to a sign array.** Replace each `nums[i]` by `+1` if it equals `target` and `-1` otherwise. If a subarray contains `c` copies of `target` over a length `L`, its transformed sum is `c - (L - c) = 2c - L`. The condition "`target` is the strict majority" means `c > L/2`, which is equivalent to `2c - L > 0`. So **`target` is the majority of a subarray iff that subarray's transformed sum is strictly positive.**

**2. Move to prefix sums.** Let `P[0] = 0` and `P[k] = transformed[0] + ... + transformed[k-1]`. The sum of subarray `nums[l..r]` (inclusive) equals `P[r+1] - P[l]`. We want this to be `> 0`, i.e. `P[l] < P[r+1]`. Letting `i = l` and `j = r+1`, every valid subarray corresponds to a unique pair of prefix indices `i < j` with `P[i] < P[j]`.

**3. Count ordered pairs efficiently.** We now need: over the sequence `P[0], P[1], ..., P[n]`, how many pairs `(i, j)` satisfy `i < j` and `P[i] < P[j]`? Sweep `j` from left to right while maintaining a frequency structure of all prefix values seen so far. For the current `P[j]`, add the count of stored values strictly less than `P[j]`, then insert `P[j]`. Prefix values lie in `[-n, n]`, so shift them into `[1, 2n+1]` and use a Fenwick tree where `query(x)` returns how many inserted values are `<= x`; we ask for `query(P[j] - 1)`.

**Walkthrough — Example 1:** `nums = [1,2,2,3]`, `target = 2`.
- Transformed: `[-1, +1, +1, -1]`.
- Prefix sums: `P = [0, -1, 0, 1, 0]`.
- Count pairs `i < j` with `P[i] < P[j]`:
  - `j=1 (-1)`: earlier `{0}` smaller than `-1` → 0
  - `j=2 (0)`: earlier `{0,-1}` smaller than `0` → 1 (the `-1`)
  - `j=3 (1)`: earlier `{0,-1,0}` smaller than `1` → 3
  - `j=4 (0)`: earlier `{0,-1,0,1}` smaller than `0` → 1 (the `-1`)
  - Total = `0 + 1 + 3 + 1 = 5` ✓

The same machinery gives `10` for `[1,1,1,1]` (strictly increasing prefix → all `C(5,2)` pairs) and `0` for `[1,2,3]` with `target=4` (strictly decreasing prefix → no pairs).

This is honestly a tricky Hard: the leap from "majority" to "positive sum" to "count ordered pairs" is the whole battle. Once you see it, the implementation is a textbook Fenwick sweep.

## Complexity Analysis

Time Complexity: O(n log n) — one Fenwick update and one Fenwick query per element, each `O(log n)`.
Space Complexity: O(n) — the Fenwick tree has `2n + 2` slots.

## Edge Cases

- **`target` never appears** (Example 3): every transformed value is `-1`, prefix sums strictly decrease, and no pair satisfies `P[i] < P[j]`, so the answer is `0`. The algorithm handles this naturally.
- **All elements equal `target`** (Example 2): prefix sums strictly increase, yielding all `C(n+1, 2)` pairs — verify your counting includes the `P[0] = 0` sentinel so single-element subarrays are counted.
- **Single-element array**: `[target]` → 1, `[other]` → 0. The `P[0]` sentinel plus one step produces the right result.
- **Large answers**: with `n = 10^5`, the count can reach about `5 * 10^9`, which overflows 32-bit integers. Use a 64-bit accumulator (`int64`).
- **Negative prefix values**: prefix sums range from `-n` to `n`; remember to shift indices into `[1, 2n+1]` so the Fenwick tree (1-indexed) is happy.
