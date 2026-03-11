# 1009. Complement of Base 10 Integer

[LeetCode Link](https://leetcode.com/problems/complement-of-base-10-integer/)

Difficulty: Easy
Topics: Bit Manipulation
Acceptance Rate: 60.9%

## Hints

### Hint 1

Think about what "complement" really means at the bit level. You need to flip every bit, but only the bits that matter — not all 32 (or 64) bits of the integer. What distinguishes the "meaningful" bits from the leading zeros?

### Hint 2

If you XOR a number with a mask of all 1s that has the same bit-length as the number, you effectively flip every bit. The key question becomes: how do you construct that mask?

### Hint 3

A mask of all 1s with `k` bits is simply `2^k - 1`. You can find `k` by determining the bit-length of `n` — the position of the highest set bit. Then `n XOR mask` gives you the complement. Don't forget the special case when `n = 0`.

## Approach

The complement of a number flips all bits within its binary representation. For example, `5 = 101` becomes `010 = 2`. The trick is that we only want to flip the bits that are part of the number's actual binary representation, not the full 32-bit or 64-bit width.

**Algorithm:**

1. Handle the edge case: if `n == 0`, return `1` (the complement of `0` in binary is `1`).
2. Find a bitmask that has all 1s and is the same length as the binary representation of `n`. We can do this by starting with `mask = 1` and left-shifting while `mask <= n`, which doubles the mask each time. After the loop, `mask - 1` gives us a value like `111...1` with the right number of bits.
3. XOR `n` with the mask to flip all meaningful bits: `n ^ (mask - 1)`.

**Example walkthrough with n = 5:**
- Binary of 5: `101`
- We build mask: start at 1, shift to 2, 4, 8. Stop when mask (8) > 5.
- `mask - 1 = 7 = 111`
- `5 XOR 7 = 101 XOR 111 = 010 = 2`

## Complexity Analysis

Time Complexity: O(log n) — we iterate once through the bits of n to build the mask.
Space Complexity: O(1) — only a few integer variables are used.

## Edge Cases

- **n = 0**: The complement of `0` is `1`. This must be handled explicitly since the bit-length of 0 is technically 0, and the mask-building loop wouldn't execute.
- **n = 1**: The complement of `1` (`1` in binary) is `0` (`0` in binary). The mask is `1`, and `1 XOR 1 = 0`.
- **n is all 1s in binary (e.g., 7 = 111)**: The complement should be `0`. The mask equals `n`, so `n XOR n = 0`.
- **Large values near 10^9**: Ensure the mask variable doesn't overflow. Using standard int in Go (64-bit) handles this comfortably.
