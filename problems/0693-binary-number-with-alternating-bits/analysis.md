# 0693. Binary Number with Alternating Bits

[LeetCode Link](https://leetcode.com/problems/binary-number-with-alternating-bits/)

Difficulty: Easy
Topics: Bit Manipulation
Acceptance Rate: 65.0%

## Hints

### Hint 1

Think about what "alternating bits" means at the bit level. How do adjacent bits relate to each other? Consider what bitwise operation lets you compare a bit with the one next to it.

### Hint 2

If you right-shift the number by 1, the bit at position `i` moves to position `i-1`. Now each bit is lined up with its former neighbor. What happens if you XOR the original number with its shifted version? What would the result look like for a valid alternating-bits number?

### Hint 3

For a number with alternating bits (like `101` or `1010`), XORing it with itself right-shifted by 1 produces a number where every bit is `1` (e.g., `101 XOR 010 = 111`). A number that is all `1`s in binary has a special property: adding 1 to it makes it a power of 2. So check whether `(n XOR (n >> 1)) + 1` is a power of 2. A power of 2 has exactly one bit set, which you can verify with `x & (x - 1) == 0`.

## Approach

1. Compute `x = n XOR (n >> 1)`.
   - If `n` has alternating bits, every adjacent pair of bits differs, so XOR produces a `1` for each pair. The result `x` will be a sequence of all `1`s (e.g., `0b111` or `0b1111`).
   - If `n` does NOT have alternating bits, at least one adjacent pair will match, producing a `0` in that position.

2. Check whether `x` is all `1`s.
   - A number consisting of all `1`s in binary satisfies `x & (x + 1) == 0`. For example, `0b111` (7) + 1 = `0b1000` (8), and `7 & 8 == 0`.
   - If this condition holds, `n` has alternating bits.

**Example walkthrough with n = 5 (`101`):**
- `n >> 1` = `010` (which is 2)
- `n XOR (n >> 1)` = `101 XOR 010` = `111` (which is 7)
- `7 & (7 + 1)` = `7 & 8` = `0b0111 & 0b1000` = `0` → return `true`

**Example walkthrough with n = 7 (`111`):**
- `n >> 1` = `011` (which is 3)
- `n XOR (n >> 1)` = `111 XOR 011` = `100` (which is 4)
- `4 & (4 + 1)` = `4 & 5` = `0b100 & 0b101` = `4` ≠ `0` → return `false`

## Complexity Analysis

Time Complexity: O(1) — only a fixed number of bitwise operations regardless of input size.
Space Complexity: O(1) — no additional data structures used.

## Edge Cases

- **n = 1 (`1`)**: Single bit is trivially alternating. `1 XOR 0 = 1`, and `1 & 2 = 0` → `true`.
- **n = 2 (`10`)**: Two alternating bits. `2 XOR 1 = 3`, and `3 & 4 = 0` → `true`.
- **n = 3 (`11`)**: Two identical adjacent bits. `3 XOR 1 = 2`, and `2 & 3 = 2` ≠ `0` → `false`.
- **Large values near 2^31 - 1**: The bit manipulation approach works correctly for all values within the 32-bit integer range since Go's `int` is at least 64 bits wide.
