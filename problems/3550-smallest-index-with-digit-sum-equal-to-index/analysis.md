# 3550. Smallest Index With Digit Sum Equal to Index

[LeetCode Link](https://leetcode.com/problems/smallest-index-with-digit-sum-equal-to-index/)

Difficulty: Easy
Topics: Array, Math
Acceptance Rate: 82.2%

## Hints

### Hint 1

You're asked for the *smallest* index that satisfies a condition. What's the most natural order to visit indices if you want the first one that works?

### Hint 2

Split the problem into two pieces: a helper that computes the digit sum of a single number, and a loop that compares that digit sum with the current index.

### Hint 3

Scan left to right and return immediately on the first match — no need to look further. The digit sum is computed with the classic `n % 10` / `n / 10` loop.

## Approach

1. Write a helper `digitSum(n)` that repeatedly adds `n % 10` to a running total and divides `n` by 10 until `n` becomes 0.
2. Iterate over `nums` with index `i` from 0 upward.
3. If `digitSum(nums[i]) == i`, return `i` right away. Because we scan in increasing index order, the first match is guaranteed to be the smallest.
4. If the loop finishes without a match, return `-1`.

Example: `nums = [1, 10, 11]`
- `i = 0`: digit sum of 1 is 1 ≠ 0
- `i = 1`: digit sum of 10 is 1 = 1 → return 1

A small observation: since `nums[i] <= 1000`, a digit sum is at most 27 (for 999), so only indices 0..27 can ever match. This could be used to stop early, but it isn't necessary given the tiny input size.

## Complexity Analysis

Time Complexity: O(n · d), where d is the number of digits (at most 4), so effectively O(n)
Space Complexity: O(1)

## Edge Cases

- **`nums[0] = 0`**: digit sum of 0 is 0, which matches index 0. Make sure your digit-sum loop returns 0 for input 0 rather than misbehaving.
- **Multiple matching indices**: only the smallest should be returned — early return handles this.
- **No match at all**: return `-1`.
- **Single element array**: only index 0 is possible, so the answer is 0 only when `nums[0] == 0`.
- **Large values like 1000**: digit sum is 1, showing that large numbers can still match small indices.
