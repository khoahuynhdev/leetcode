# 2333. Minimum Sum of Squared Difference

[LeetCode Link](https://leetcode.com/problems/minimum-sum-of-squared-difference/)

Difficulty: Medium
Topics: Array, Binary Search, Greedy, Sorting, Heap (Priority Queue)
Acceptance Rate: 33.6%

## Hints

### Hint 1

Does it matter whether you spend an operation on `nums1[i]` or on `nums2[i]`? Each one changes `|nums1[i] - nums2[i]|` by exactly 1. Try to simplify the problem down to a single array and a single budget.

### Hint 2

With diffs `d[i] = |nums1[i] - nums2[i]|` and a budget `k = k1 + k2`, each operation lowers one diff by 1. Squares grow fast, so lowering a large value saves more than lowering a small one: `x² - (x-1)² = 2x - 1`. A greedy choice will work, but simulating one operation at a time is far too slow when `k` can be 2·10⁹.

### Hint 3

The best you can do is "flatten the peaks": repeatedly reduce all elements tied at the current maximum down together, like water draining from the top. Diff values are at most 10⁵, so count how many diffs have each value (a bucket array) and sweep from the maximum value downward, moving the whole group at level `v` down to `v - 1` in one step. When the budget can't cover the whole group, only part of it drops by one more, and you're done.

## Approach

1. **Reduce to one array.** Compute `d[i] = |nums1[i] - nums2[i]|`. An operation on either array can shrink `d[i]` by 1, so the two budgets merge into `k = k1 + k2`.
2. **Greedy rationale.** Reducing `x` to `x - 1` saves `2x - 1`, which is larger for larger `x`. So every operation should go to a current maximum. The final result is that all values above some threshold `T` get lowered to `T` (with some left at `T - 1` if budget remains partially).
3. **Bucket counting.** Since diffs are at most 10⁵, build `cnt[v]` = number of diffs equal to `v`.
4. **Sweep from the top.** For `v` from max down to 1:
   - `cnt[v]` is the number of elements currently at level `v` (including ones moved down from above).
   - If `k >= cnt[v]`, lower all of them to `v - 1`: `cnt[v-1] += cnt[v]`, `k -= cnt[v]`, `cnt[v] = 0`.
   - Otherwise lower only `k` of them: `cnt[v-1] += k`, `cnt[v] -= k`, `k = 0`, and stop.
5. If the sweep reaches 0 (budget covers every unit of difference), the answer is 0. Operations are "at most" `k1`/`k2`, so leftover budget is simply unused.
6. **Sum up** `cnt[v] * v * v` over all `v` using 64-bit integers (the result can reach 10⁵ · (10⁵)² = 10¹⁵).

**Example 2:** diffs are `[4, 4, 4, 3]`, `k = 2`. At level 4 there are 3 elements but only 2 operations, so two drop to 3: diffs become `[3, 3, 4, 3]` → `9 + 9 + 16 + 9 = 43`.

## Complexity Analysis

Time Complexity: O(n + M), where M = max diff ≤ 10⁵
Space Complexity: O(M) for the bucket array

(An alternative is sorting + binary search on the threshold in O(n log n), or a heap-based simulation of groups.)

## Edge Cases

- **k1 = k2 = 0:** No changes allowed; just return the original sum.
- **Budget exceeds total difference:** If `k >= sum(d)`, the answer is 0 — don't let values go negative or keep spending.
- **All diffs already zero:** The sweep never runs; answer 0.
- **Partial group at the threshold:** When `k` can't cover the whole tied group, only `k` of them drop one level — getting this split right is the crux.
- **Overflow:** `k1 + k2` up to 2·10⁹ and the sum up to 10¹⁵ both require 64-bit integers.
