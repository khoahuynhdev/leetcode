# 1385. Find the Distance Value Between Two Arrays

[LeetCode Link](https://leetcode.com/problems/find-the-distance-value-between-two-arrays/)

Difficulty: Easy
Topics: Array, Two Pointers, Binary Search, Sorting
Acceptance Rate: 74.9%

## Hints

### Hint 1

The brute force (check every pair) works within the constraints, but think about how to answer "is there any element of `arr2` close to `x`?" faster. What makes "closest value" queries cheap?

### Hint 2

Sort `arr2`. Once it's sorted, you don't need to compare `x` with every element — only with the elements nearest to `x`. Binary search can find them.

### Hint 3

For each `x` in `arr1`, you only need to know whether any value lies in the range `[x - d, x + d]`. Binary search for the first element `>= x - d` in sorted `arr2`. If it exists and is `<= x + d`, then `x` does **not** count; otherwise it does.

## Approach

1. Sort `arr2` in ascending order.
2. For each `x` in `arr1`:
   - Use binary search to find the index `i` of the first element in `arr2` that is `>= x - d` (a lower bound).
   - If `i == len(arr2)` (every element is smaller than `x - d`) or `arr2[i] > x + d`, then no element falls in the window `[x - d, x + d]`, so `x` contributes 1 to the answer.
   - Otherwise, `arr2[i]` is within distance `d` of `x`, so skip it.
3. Return the count.

Why it works: `|x - y| <= d` is equivalent to `x - d <= y <= x + d`. In a sorted array, the smallest value that is `>= x - d` is the only candidate we need to check — if even it exceeds `x + d`, every larger value does too.

Example: `arr1 = [4,5,8]`, `arr2 = [10,9,1,8]`, `d = 2` → sorted `arr2 = [1,8,9,10]`.
- `x = 4`: window `[2,6]`. Lower bound of 2 → `8`, and `8 > 6` → counts.
- `x = 5`: window `[3,7]`. Lower bound of 3 → `8`, and `8 > 7` → counts.
- `x = 8`: window `[6,10]`. Lower bound of 6 → `8`, and `8 <= 10` → doesn't count.

Answer: 2.

An alternative is to sort both arrays and sweep with two pointers, giving O(n log n + m log m) time dominated by sorting.

## Complexity Analysis

Time Complexity: O((n + m) log m), where n = len(arr1) and m = len(arr2) — sorting `arr2` plus one binary search per element of `arr1`.
Space Complexity: O(1) extra (ignoring the sort's internal stack); `arr2` is sorted in place.

## Edge Cases

- **All elements of `arr2` smaller than `x - d`**: the lower bound returns `len(arr2)`; must not index out of bounds.
- **`d = 0`**: only exact matches disqualify an element — the window collapses to a single value.
- **Negative numbers**: values range from -1000 to 1000, so the window can cross zero; the comparison logic must not assume non-negative values.
- **Boundary distance exactly `d`**: `|x - y| == d` disqualifies `x` (the condition is `<= d`), so use inclusive bounds.
- **Duplicates in either array**: each element of `arr1` is counted independently, duplicates included.
