# 3783. Mirror Distance of an Integer

[LeetCode Link](https://leetcode.com/problems/mirror-distance-of-an-integer/)

Difficulty: Easy
Topics: Math
Acceptance Rate: 89.0%

## Hints

### Hint 1

This is a digit-manipulation problem. Think about how you can break an integer apart into its digits and then put it back together in a different order. The classic toolkit here involves modulo and integer division.

### Hint 2

You don't need strings or any fancy data structure. A simple loop that repeatedly peels off the last digit of `n` (using `n % 10`) and shifts it onto a new accumulator (multiplying by 10 each step) will reverse the digits. Once you have both numbers, the rest is a one-line absolute difference.

### Hint 3

Watch out for trailing zeros. When you reverse `10`, the leading zero disappears and you get `1`, not `01`. The peel-and-shift approach naturally handles this because you only start accumulating once you encounter the first non-zero digit from the right — any leading zeros in the reversed form are simply never added.

## Approach

The problem reduces to two steps: compute `reverse(n)`, then return `|n - reverse(n)|`.

To reverse the digits of a positive integer without converting to a string:

1. Initialize `rev = 0`.
2. While `n > 0`:
   - Extract the last digit: `d = n % 10`.
   - Append it to `rev`: `rev = rev * 10 + d`.
   - Drop the last digit from `n`: `n = n / 10` (integer division).
3. After the loop, `rev` holds the reversed integer.

Because we use a working copy of `n` to compute the reverse, the original `n` is preserved for the final subtraction. The answer is `abs(n - rev)`.

**Worked example (`n = 25`):**

- Iteration 1: `d = 25 % 10 = 5`, `rev = 0*10 + 5 = 5`, `n = 25 / 10 = 2`.
- Iteration 2: `d = 2 % 10 = 2`, `rev = 5*10 + 2 = 52`, `n = 2 / 10 = 0`.
- Loop ends. `|25 - 52| = 27`. ✓

**Worked example (`n = 10`):**

- Iteration 1: `d = 10 % 10 = 0`, `rev = 0*10 + 0 = 0`, `n = 1`.
- Iteration 2: `d = 1 % 10 = 1`, `rev = 0*10 + 1 = 1`, `n = 0`.
- Loop ends. `|10 - 1| = 9`. ✓

The leading-zero issue resolves itself: the trailing `0` in `10` is multiplied into `rev` while `rev` is still `0`, contributing nothing.

Since `n <= 10^9`, the reversed value also fits comfortably in a 32-bit signed integer, so no overflow handling is required in Go's `int`.

## Complexity Analysis

Time Complexity: O(log₁₀ n) — one iteration per decimal digit of `n` (at most 10 digits for the given constraints).
Space Complexity: O(1) — only a couple of integer variables are used.

## Edge Cases

- **Single-digit `n` (e.g., `n = 7`):** `reverse(7) = 7`, so the answer is `0`. The loop runs once and produces the same number.
- **Trailing zeros (e.g., `n = 10`, `n = 1000000000`):** Reversing drops the leading zeros, yielding a much smaller number. The peel-and-shift method handles this naturally.
- **Palindromic numbers (e.g., `n = 121`, `n = 1`):** The reverse equals the original, so the answer is `0`.
- **Maximum value (`n = 10^9 = 1000000000`):** Reverse is `1`, answer is `999999999`. Fits comfortably in `int`.
- **Smallest value (`n = 1`):** A single-digit palindrome — answer is `0`.
