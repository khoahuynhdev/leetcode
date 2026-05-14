# 2784. Check if Array is Good

[LeetCode Link](https://leetcode.com/problems/check-if-array-is-good/)

Difficulty: Easy
Topics: Array, Hash Table, Sorting
Acceptance Rate: 51.6%

## Hints

### Hint 1

Think about what a "good" array must look like structurally. If the array is a permutation of `base[n]`, what does that say about its length, its maximum value, and which values must appear? Try to write down the structural properties any good array must satisfy before reaching for a data structure.

### Hint 2

The candidate value of `n` is forced — there is exactly one possible `n` for any given input. Once you know `n`, the check reduces to verifying multiset equality between `nums` and `base[n] = [1, 2, ..., n-1, n, n]`. Two common tools for multiset equality are: (a) sorting both sides, or (b) counting frequencies with a hash map / fixed-size array.

### Hint 3

The maximum element of `nums` must equal `n`, and the length of `nums` must equal `n + 1`. After that, you only need to confirm that every value from `1` to `n - 1` appears exactly once and the value `n` appears exactly twice. If either the length check fails or any count is off, return false. Sorting gives a particularly clean one-pass verification: after sorting, position `i` (0-indexed) must hold the value `i + 1` for `i` in `[0, n-1]`, and the last slot must also hold `n`.

## Approach

The key observation is that `base[n]` is fully determined by `n`, and `n` is fully determined by `nums` (it must be `max(nums)`). So the algorithm is:

1. Find `n = max(nums)`.
2. If `len(nums) != n + 1`, return false — the lengths cannot match.
3. Sort `nums` ascending.
4. Walk indices `0..n-1` and confirm `nums[i] == i + 1`. This guarantees `1, 2, ..., n-1` each appear and that the first `n` slots end with the value `n` (since `nums[n-1] == n`).
5. The final slot `nums[n]` must also equal `n` (it is the largest after sorting, equal to `max(nums)`, so this falls out — but we still ensure the second-to-last slot is `n` via step 4).

Example: `nums = [1, 3, 3, 2]`. `n = 3`, `len = 4 = n + 1`. Sorted: `[1, 2, 3, 3]`. Check `nums[0]=1`, `nums[1]=2`, `nums[2]=3`. Last slot is `3 = n`. All good → true.

Counter-example: `nums = [3, 4, 4, 1, 2, 1]`. `n = 4`, but `len = 6 ≠ 5`, so we return false immediately.

A frequency-array variant works equally well: bump counts and verify `count[v] == 1` for `v` in `[1, n-1]` and `count[n] == 2`. The sort version is slightly easier to get right in fewer lines.

## Complexity Analysis

Time Complexity: O(k log k) where k = len(nums), dominated by sorting. (A counting-based variant achieves O(k + n).)
Space Complexity: O(1) extra if we sort in place, ignoring the sort's internal stack. O(n) for the counting variant.

## Edge Cases

- Smallest possible input `nums = [1, 1]`: this is `base[1]`, must return true.
- Length mismatch (too short or too long): e.g. `nums = [2, 1, 3]` — max is 3 but length is 3, not 4. Early-return on length is what saves you here.
- Duplicates of a non-max value: e.g. `nums = [1, 1, 2, 3]` — max is 3, length 4 matches `n + 1`, but after sorting `nums[1]` is 1 not 2, so the per-index check catches it.
- Missing values in the `1..n-1` range: e.g. `nums = [1, 3, 3, 3]` — length matches, max is 3, but `2` is missing; the sorted check detects the gap.
- Values that exceed `n` are impossible by construction (since `n = max`), but the value `n` appearing more than twice (e.g. `[3, 3, 3, 3]`) is caught because `nums[1]` would be 3 instead of 2.
