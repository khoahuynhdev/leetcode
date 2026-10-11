# 2778. Sum of Squares of Special Elements 

[LeetCode Link](https://leetcode.com/problems/sum-of-squares-of-special-elements/)

Difficulty: Easy
Topics: Array, Enumeration
Acceptance Rate: 84.4%

## Hints

### Hint 1

Whether an element is "special" depends only on its position, not on its value. Think about which positions qualify before you think about the values stored there.

### Hint 2

Go through every possible 1-indexed position `i` from `1` to `n` and check one simple divisibility condition. Watch out for the gap between the problem's 1-based indexing and Go's 0-based slices.

### Hint 3

Position `i` is special exactly when `n % i == 0`, i.e. `i` is a divisor of `n`. For each such `i`, add `nums[i-1] * nums[i-1]` to a running sum. (If you want to be fancy, you can enumerate divisors in pairs `i` and `n/i` up to `sqrt(n)`, but with `n <= 50` a plain loop is perfectly fine.)

## Approach

1. Let `n = len(nums)`.
2. Loop `i` from `1` to `n` (1-indexed positions).
3. If `n % i == 0`, the element at position `i` is special. In Go it lives at `nums[i-1]`, so add `nums[i-1]²` to the answer.
4. Return the sum.

Example: `nums = [2,7,1,19,18,3]`, `n = 6`. The divisors of 6 are `1, 2, 3, 6`, so the special elements are `2, 7, 1, 3`. Sum of squares: `4 + 49 + 1 + 9 = 63`.

Why it works: the definition of "special" is literally "the index divides `n`", so checking every index against that condition finds exactly the special elements — nothing is missed and nothing extra is counted.

Optional optimization: divisors come in pairs `(i, n/i)`. Iterating `i` only up to `sqrt(n)` and adding both `nums[i-1]²` and `nums[n/i-1]²` (taking care not to double-count when `i == n/i`) brings the time down to O(√n). It's a nice trick to know, but unnecessary here.

## Complexity Analysis

Time Complexity: O(n) — one pass over the positions (O(√n) with the divisor-pair trick).
Space Complexity: O(1) — only a running sum.

## Edge Cases

- **Single element (`n = 1`)**: position 1 divides 1, so the answer is just `nums[0]²`.
- **Prime length**: only positions `1` and `n` are special — make sure you include the last element.
- **Off-by-one indexing**: the problem is 1-indexed; using `i % n` or starting at `i = 0` (division/modulo by zero) is a common mistake.
- **Perfect-square length** (if using the √n optimization): don't count the middle divisor `√n` twice.
- **Overflow**: max sum is `50 * 50² = 125,000`, which fits comfortably in an `int`.
