# 2657. Find the Prefix Common Array of Two Arrays

[LeetCode Link](https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/)

Difficulty: Medium
Topics: Array, Hash Table, Bit Manipulation
Acceptance Rate: 87.3%

## Hints

### Hint 1

The naive idea is, for every prefix length `i`, to re-scan both prefixes and count the shared values. That works but does redundant work. Ask yourself: as `i` grows by one, what *new* information appears, and can the answer for `i` be built incrementally from the answer for `i - 1`? Think about what data structure lets you answer "have I seen this value before?" in constant time.

### Hint 2

You only add two new elements when moving from index `i - 1` to `i`: `A[i]` and `B[i]`. A value becomes "common" the moment it has appeared in *both* the A-prefix and the B-prefix. So instead of recounting, keep a tally of how many times you have encountered each value across the two arrays so far. A frequency table (array of size `n + 1`, since values are in `1..n`) is a perfect hash table here.

### Hint 3

The key insight: a value `v` is common at index `i` exactly when it has been seen **twice** in the combined stream — once in `A[0..i]` and once in `B[0..i]` (the array values are unique within each permutation, so a second sighting must come from the other array). Maintain a running `common` counter. Each step, bump the tally for `A[i]` and `B[i]`; whenever a tally crosses to `2`, increment `common`. Then `C[i] = common`. This carries the previous answer forward, so it never decreases.

## Approach

We process the arrays left to right, one index at a time, and reuse the count from the previous index instead of recomputing it.

Maintain:
- `seen` — an integer array of length `n + 1`, where `seen[v]` is how many times value `v` has appeared so far across both `A` and `B` (values range from `1` to `n`).
- `common` — a running total of how many distinct values have shown up in both prefixes.

For each index `i` from `0` to `n - 1`:
1. Increment `seen[A[i]]`. If it just became `2`, that value has now appeared in both arrays' prefixes, so increment `common`.
2. Increment `seen[B[i]]`. If it just became `2`, increment `common`.
3. Record `C[i] = common`.

Why it works: within a single permutation every value is unique, so `seen[v]` can reach at most `2` — at most one sighting from `A` and one from `B`. The moment `seen[v]` hits `2`, value `v` is guaranteed to be present in both prefixes, and it stays common for all later indices. Because `common` only ever grows, the prefix-common array is non-decreasing, which matches the problem's definition.

Walking through Example 1, `A = [1,3,2,4]`, `B = [3,1,2,4]`:
- `i = 0`: see `1` and `3`, both first sightings → `common = 0`, `C[0] = 0`.
- `i = 1`: see `3` (now twice → `common = 1`) and `1` (now twice → `common = 2`) → `C[1] = 2`.
- `i = 2`: see `2` (first), then `2` again (twice → `common = 3`) → `C[2] = 3`.
- `i = 3`: see `4` (first), then `4` again (twice → `common = 4`) → `C[3] = 4`.

Result: `[0, 2, 3, 4]`. The case where `A[i] == B[i]` is handled naturally: the first increment makes the tally `1`, the second makes it `2` and bumps `common`.

A note on the "Bit Manipulation" topic tag: with `n <= 50` you can also keep two bitmasks (one for A's prefix, one for B's prefix), OR each new value in, and use `popcount(maskA & maskB)` for `C[i]`. That is elegant but conceptually the same idea; the counting approach is just as fast and easier to read.

This problem is on the gentler side for a Medium — the trick is recognizing the incremental-counting pattern rather than re-scanning. Once that clicks, the implementation is short.

## Complexity Analysis

Time Complexity: O(n) — a single pass over the arrays, with constant work per index.
Space Complexity: O(n) — the `seen` tally of size `n + 1`. (The output array `C` is also O(n) but is required by the problem.)

## Edge Cases

- **`n = 1`**: Both arrays are `[1]`. At `i = 0`, value `1` is seen in both, so `C = [1]`. The single-element permutation always yields `[1]`.
- **Identical arrays (`A == B`)**: Every value matches at its own index, so each prefix shares everything seen so far, giving `C = [1, 2, 3, ..., n]`. Confirms the `A[i] == B[i]` path increments correctly.
- **Reversed / fully shuffled arrays**: Values become common at different times; the running counter must never decrease. This stresses that once a value is counted it stays counted.
- **1-indexed values**: Values are in `1..n`, never `0`, so sizing `seen` as `n + 1` (indices `0..n`) avoids any out-of-range access while leaving index `0` unused.
