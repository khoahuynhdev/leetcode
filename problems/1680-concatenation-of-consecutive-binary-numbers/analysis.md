# 1680. Concatenation of Consecutive Binary Numbers

[LeetCode Link](https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/)

Difficulty: Medium
Topics: Math, Bit Manipulation, Simulation
Acceptance Rate: 58.1%

## Hints

### Hint 1

Think about what "concatenating binary representations" really means in terms of arithmetic operations. If you had the number so far and wanted to append a new binary number to the right, what mathematical operation would achieve that?

### Hint 2

Appending a binary number `i` to an existing value is equivalent to shifting the existing value left by the number of bits in `i`, then adding `i`. The key question becomes: how do you efficiently determine how many bits each number `i` has?

### Hint 3

The number of bits in `i` only increases at powers of 2. You can detect a power of 2 with the bit trick `i & (i - 1) == 0`. Track the current bit length and increment it whenever you hit a power of 2, avoiding any per-number log computation. Apply modular arithmetic at each step to keep values manageable.

## Approach

We build the result incrementally. Starting with `result = 0`, for each number `i` from 1 to n:

1. **Determine the bit length of `i`**: The binary length of `i` increases by 1 exactly when `i` is a power of 2. We can detect this with `i & (i - 1) == 0`. We maintain a running `bitLen` counter and increment it at each power of 2.

2. **Shift and add**: To "concatenate" `i` onto the current result, we left-shift the result by `bitLen` positions and add `i`:
   ```
   result = (result << bitLen) + i
   ```

3. **Apply modulo**: Since the result can grow extremely large, we take `mod 10^9 + 7` at each step.

**Walkthrough with n = 3:**
- i=1: bitLen becomes 1 (1 is a power of 2). result = (0 << 1) + 1 = 1. Binary: "1"
- i=2: bitLen becomes 2 (2 is a power of 2). result = (1 << 2) + 2 = 6. Binary: "110"
- i=3: bitLen stays 2. result = (6 << 2) + 3 = 27. Binary: "11011"

Result: 27, which matches the expected output.

## Complexity Analysis

Time Complexity: O(n) — single pass from 1 to n with O(1) work per iteration.

Space Complexity: O(1) — only a few integer variables are used.

## Edge Cases

- **n = 1**: The simplest case, result is just 1. Verifies the loop handles the first iteration correctly.
- **n = a power of 2**: The bit length increases on the last number. Ensures the power-of-2 detection and bit length update happen before the shift.
- **Large n (n = 100000)**: Tests that modular arithmetic is applied correctly at every step to prevent overflow, and that the solution runs within time limits.
