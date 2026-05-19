# 2540. Minimum Common Value

[LeetCode Link](https://leetcode.com/problems/minimum-common-value/)

Difficulty: Easy
Topics: Array, Hash Table, Two Pointers, Binary Search
Acceptance Rate: 58.9%

## Hints

### Hint 1

The problem asks for the smallest integer that appears in both arrays. A brute-force scan that checks every pair would be O(n*m), which is too slow for inputs up to 10^5. Take a moment to look at the problem constraints — what extra property of the inputs can you exploit to do better?

### Hint 2

Both input arrays are already sorted in non-decreasing order. Sorted input is a strong nudge toward either binary search or a two-pointer sweep. Since we want the *minimum* common value, we are naturally interested in the *earliest* match we encounter when walking from left to right.

### Hint 3

Use two indices, one per array, both starting at 0. Compare the values at the two pointers:
- If they are equal, you have found the smallest common value (because we have only ever advanced past strictly smaller values on both sides).
- If they are unequal, the smaller one cannot be the answer (the other array has no values that small left to see), so advance the pointer pointing at the smaller value.
Stop when either pointer falls off the end.

## Approach

We use a classic two-pointer sweep over the two sorted arrays.

1. Initialize `i = 0` for `nums1` and `j = 0` for `nums2`.
2. While both `i < len(nums1)` and `j < len(nums2)`:
   - If `nums1[i] == nums2[j]`, return `nums1[i]`. This is guaranteed to be the smallest common value, because everything we have skipped on each side was strictly less than the values we are now comparing.
   - If `nums1[i] < nums2[j]`, increment `i`. The value `nums1[i]` cannot appear later in `nums2`, since `nums2[j]` is already larger and the rest of `nums2` is non-decreasing.
   - Otherwise, increment `j` by the symmetric argument.
3. If the loop ends without a match, return `-1`.

Walk-through for `nums1 = [1,2,3,6]`, `nums2 = [2,3,4,5]`:
- `(i,j) = (0,0)`: `1 < 2`, advance `i`.
- `(i,j) = (1,0)`: `2 == 2`, return `2`.

We could also use a hash set built from one array and probe with the other — that gives O(n+m) time and O(min(n,m)) space, but requires a second pass to find the *minimum*. The two-pointer approach is strictly better here because it leverages the sorted property to find the minimum directly without extra space.

## Complexity Analysis

Time Complexity: O(n + m), where `n = len(nums1)` and `m = len(nums2)`. Each iteration advances at least one pointer, so we do at most `n + m` steps.
Space Complexity: O(1). Only two integer indices are used.

## Edge Cases

- **No common element**: e.g. `nums1 = [1,2]`, `nums2 = [3,4]`. The pointers walk to the ends without ever matching; return `-1`.
- **Single-element arrays**: e.g. `nums1 = [5]`, `nums2 = [5]` (match) or `nums1 = [5]`, `nums2 = [6]` (no match). The loop handles both correctly.
- **Common value at the very start**: e.g. `nums1 = [1, ...]`, `nums2 = [1, ...]`. The first comparison returns immediately.
- **Common value at the very end**: e.g. `nums1 = [1,2,9]`, `nums2 = [3,5,9]`. Pointers advance until both reach the matching tail.
- **Duplicates within an array**: e.g. `nums1 = [1,1,2]`, `nums2 = [1,1,1]`. Duplicates do not break the algorithm — the first equal comparison wins, and the answer is still the minimum.
- **Disjoint ranges**: e.g. `nums1 = [1,2,3]`, `nums2 = [100,200]`. The shorter side runs out first and we correctly return `-1`.
