# 0190. Reverse Bits

[LeetCode Link](https://leetcode.com/problems/reverse-bits/)

Difficulty: Easy
Topics: Divide and Conquer, Bit Manipulation
Acceptance Rate: 66.0%

## Hints

### Hint 1

Think about how you would reverse a string character by character. Can you apply a similar idea to the individual bits of an integer? Consider how to extract a single bit and how to place it into a new position.

### Hint 2

You can extract the least significant bit of `n` using `n & 1`, then shift your result left to make room for the next bit. After extracting a bit, shift `n` right to expose the next one. Repeat this process for all 32 bits.

### Hint 3

Build the reversed number from scratch: initialize `result = 0`. In each of 32 iterations, left-shift `result` by 1, OR in the lowest bit of `n`, then right-shift `n` by 1. After exactly 32 iterations the bits are fully reversed — no early termination, because leading zeros in `n` become trailing zeros in `result`.

## Approach

The algorithm processes all 32 bits one at a time from least significant to most significant:

1. Start with `result = 0`.
2. For each of the 32 bit positions:
   - Left-shift `result` by 1 to make room for the next bit.
   - Extract the lowest bit of `n` with `n & 1`.
   - OR that bit into `result`.
   - Right-shift `n` by 1 to move to the next bit.
3. Return `result`.

**Why it works:** On iteration `i` (0-indexed), the bit originally at position `i` in `n` gets placed at position `31 - i` in `result`. This is exactly what "reversing" means: position 0 goes to 31, position 1 goes to 30, and so on.

**Follow-up optimization:** If the function is called many times, you can split the 32-bit integer into four 8-bit chunks, pre-compute a lookup table that maps each byte to its reversed value, and combine the four reversed bytes in reverse order. This trades O(256) memory for O(4) work per call instead of O(32).

## Complexity Analysis

Time Complexity: O(1) — always exactly 32 iterations regardless of input.
Space Complexity: O(1) — only a fixed number of variables are used.

## Edge Cases

- **n = 0**: All bits are zero. The reversed result is also 0. Important to verify the loop still runs 32 times and returns 0.
- **n = 2 (smallest even > 0)**: Binary `...10`. Reversed, the `1` moves to bit position 30, producing a large number. Tests that single-bit values reverse correctly.
- **n = 2^31 - 2 (maximum constraint value)**: Near the upper bound. All bits except bit 0 and bit 31 are set. Verifies correct handling of large inputs.
- **Even constraint**: The problem guarantees `n` is even, meaning bit 0 is always 0. This means bit 31 of the result is always 0, so the result always fits in a signed 32-bit integer.
