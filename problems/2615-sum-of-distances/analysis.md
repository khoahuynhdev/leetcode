# 2615. Sum of Distances

[LeetCode Link](https://leetcode.com/problems/sum-of-distances/)

Difficulty: Medium
Topics: Array, Hash Table, Prefix Sum
Acceptance Rate: 36.8%

## Hints

### Hint 1

The brute-force approach is O(n^2): for every `i`, scan every `j` with the same value and add `|i - j|`. With `n` up to 10^5, that is far too slow. Notice that the work for each `i` only involves indices that share the same value. What data structure lets you quickly collect all indices that share a value?

### Hint 2

Once you group indices by value, each group is an independent subproblem: given a sorted list of indices `g[0] < g[1] < ... < g[k-1]`, for every `p` compute `sum over all q != p of |g[p] - g[q]|`. The absolute value is the tricky part — can you eliminate it by splitting the sum into "indices to the left of `p`" and "indices to the right of `p`"?

### Hint 3

Because the group is sorted, for index `g[p]`:
- Left contribution: `sum_{q < p} (g[p] - g[q]) = p * g[p] - prefixSum[p]`
- Right contribution: `sum_{q > p} (g[q] - g[p]) = (totalSum - prefixSum[p+1]) - (k - 1 - p) * g[p]`

Maintain a running prefix sum as you iterate the group once. Each index is then answered in O(1). Also remember that the answers can exceed `int32` range (distances up to ~10^10), so use 64-bit integers.

## Approach

1. Walk through `nums` once and bucket indices by value into a hash map `groups: map[value] -> []int` (the indices list for each value is naturally in increasing order because we insert left-to-right).
2. For each group of indices `g` of length `k`:
   - If `k == 1`, every answer for this group is `0` (the problem says set `arr[i] = 0` when no other matching `j` exists) — skip.
   - Otherwise, compute `totalSum = g[0] + g[1] + ... + g[k-1]` once.
   - Walk the group with a running `prefixSum` (sum of indices strictly before position `p`). For each `p`:
     - `leftSum = p * g[p] - prefixSum`
     - `rightSum = (totalSum - prefixSum - g[p]) - (k - 1 - p) * g[p]`
     - `arr[g[p]] = leftSum + rightSum`
     - Add `g[p]` to `prefixSum` before moving on.

### Worked example

`nums = [1, 3, 1, 1, 2]`

Groups: `{1: [0, 2, 3], 3: [1], 2: [4]}`.

For group `1: [0, 2, 3]`, `totalSum = 5`:
- `p=0`, `g[p]=0`: left `= 0*0 - 0 = 0`, right `= (5 - 0 - 0) - 2*0 = 5`. → `arr[0] = 5`
- `p=1`, `g[p]=2`: left `= 1*2 - 0 = 2`, right `= (5 - 0 - 2) - 1*2 = 1`. → `arr[2] = 3`
- `p=2`, `g[p]=3`: left `= 2*3 - 2 = 4`, right `= (5 - 2 - 3) - 0*3 = 0`. → `arr[3] = 4`

Groups `3` and `2` each have a single index, so those answers stay `0`. Final: `[5, 0, 3, 4, 0]`.

The key insight is that once a group is sorted (which it already is when built by scanning left-to-right), the absolute value collapses into two prefix-sum-friendly sums.

## Complexity Analysis

Time Complexity: O(n) — each index is inserted into its group once, then visited once inside its group. Hash-map operations are amortized O(1).
Space Complexity: O(n) — the hash map stores all n indices across its groups, plus the output array of size n.

## Edge Cases

- **Single element** (`nums.length == 1`): the only element has no other matching index, so `arr = [0]`.
- **All distinct values**: every group has size 1, so `arr[i] = 0` for all `i` (Example 2).
- **All identical values**: a single large group drives all answers; make sure the prefix-sum math handles endpoints (`p = 0` and `p = k-1`) correctly — at the endpoints one of the two sums is zero, not negative.
- **Large indices / large groups**: at `n = 10^5` with most elements equal, distances can reach roughly `n^2/4 ≈ 2.5 * 10^9`, which exceeds `int32`. Use `int64` (Go's `int64`) for the running sums and output to avoid overflow.
- **Values up to 10^9**: values are used only as map keys, never arithmetic, so their magnitude is irrelevant to correctness.
