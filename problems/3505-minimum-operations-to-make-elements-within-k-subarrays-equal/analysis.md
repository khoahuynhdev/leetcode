# 3505. Minimum Operations to Make Elements Within K Subarrays Equal

[LeetCode Link](https://leetcode.com/problems/minimum-operations-to-make-elements-within-k-subarrays-equal/)

Difficulty: Hard
Topics: Array, Hash Table, Math, Dynamic Programming, Sliding Window, Heap (Priority Queue)
Acceptance Rate: 51.3%

## Hints

### Hint 1

Split the problem in two independent parts: (a) how much does it cost to make one particular window of size `x` all-equal, and (b) given those costs, which `k` non-overlapping windows should you choose?

### Hint 2

For part (a), recall the classic fact: the value `t` that minimizes `sum |a_i - t|` is the **median**. Since every window has the same size `x`, slide it across the array and maintain the median dynamically (two heaps, or a Fenwick tree over compressed values). For part (b), `k <= 15` is small — think of a DP indexed by (number of windows chosen, prefix length).

### Hint 3

Let `cost[i]` be the cost of the window starting at `i`. Then
`dp[j][i] = min(dp[j][i-1], dp[j-1][i-x] + cost[i-x])` — either position `i-1` isn't the end of a chosen window, or it is, and the window occupies `[i-x, i-1]`. Computing each `cost[i]` in `O(log n)` with the sliding median makes the whole thing `O(n log n + n·k)`.

## Approach

**Step 1 — Window costs via sliding median.**
For a fixed window, moving every element to the median `med` costs

```
med * cntLeft - sumLeft + sumRight - med * cntRight
```

where "left" are elements `<= med` and "right" are the rest. We need the median and these prefix aggregates for each of the `n - x + 1` windows.

We coordinate-compress the values and keep two Fenwick trees (BITs): one for counts, one for sums. Sliding the window = one add + one remove. The median is the `ceil(x/2)`-th smallest element, which we find with Fenwick binary lifting in `O(log n)`. A prefix query at the median's rank gives `cntLeft` and `sumLeft`; the right side follows from the window total.

(Two heaps with lazy deletion, or a pair of sorted multisets, work equally well — the BIT is just compact to write in Go.)

**Step 2 — DP for choosing k windows.**
`dp[j][i]` = minimal cost to place `j` non-overlapping windows inside `nums[0..i-1]`.

- `dp[0][i] = 0`
- `dp[j][i] = min(dp[j][i-1], dp[j-1][i-x] + cost[i-x])` for `i >= x`
- otherwise `dp[j][i] = +inf`

Answer: `dp[k][n]`. Only the previous row is needed, so we roll two arrays.

**Example 2:** `nums = [9,-2,-2,-2,1,5], x = 2`. Window costs: `[11, 0, 0, 3, 4]`. Picking windows at 1 and 3 costs `0 + 3 = 3`; windows at 1 and 2 overlap, so they can't both be taken. DP yields `3`.

Why it works: window cost is independent of other windows (non-overlapping means no element is shared), so the global problem is just a weighted interval selection with fixed-length intervals — exactly what the DP enumerates.

## Complexity Analysis

Time Complexity: O(n log n + n·k)
Space Complexity: O(n)

## Edge Cases

- **`k * x == n`**: the windows must tile the entire array; the DP must still find that unique layout (make sure `+inf` states don't leak).
- **Even `x`**: any value between the two middle elements is optimal; picking the lower median (`ceil(x/2)`-th smallest) is fine.
- **Large magnitudes**: values up to `1e6`, windows up to `1e5` long — costs reach ~`1e11`, so use `int64`.
- **Duplicate values**: compression must map duplicates to the same rank; the count tree handles multiplicity.
- **Already-equal windows**: cost `0`; answer may be `0`.
- **Single window covering the whole array** (`x == n, k == 1`): only one window exists.
