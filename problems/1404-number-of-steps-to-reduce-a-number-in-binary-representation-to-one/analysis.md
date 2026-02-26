# 1404. Number of Steps to Reduce a Number in Binary Representation to One

[LeetCode Link](https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/)

Difficulty: Medium
Topics: String, Bit Manipulation, Simulation
Acceptance Rate: 61.8%

## Hints

### Hint 1

Think about what "divide by 2" and "add 1" mean in binary. You don't need to convert the string to a number — the operations have very simple effects on the binary digits themselves.

### Hint 2

Dividing by 2 in binary is just removing the last bit. Adding 1 to an odd number flips the last bit from 1 to 0, but may cause a carry that propagates leftward. Can you process the string from right to left, tracking a carry, and count the steps without modifying the string?

### Hint 3

Process bits from the least significant (rightmost) to the most significant (leftmost), maintaining a carry. For each bit position (excluding the MSB): if the effective bit (original + carry) is even (0 or 2), that's one step (divide by 2); if it's odd (1), that's two steps (add 1, then divide by 2) and generates a carry. After processing all bits, check if there's a remaining carry on the MSB — if so, that's one extra step.

## Approach

We simulate the process directly on the binary string by scanning from right to left. The key observations are:

1. **Divide by 2** removes the last (rightmost) bit — this always costs 1 step.
2. **Add 1** to an odd number flips the trailing `1` to `0` and may carry into higher bits — this costs 1 step, then dividing the now-even result costs another step (2 total).

Rather than mutating the string, we walk from index `n-1` down to `1`, keeping a `carry` variable:

- Compute `bit = s[i] + carry`.
- If `bit` is even (0 or 2): one step to divide. If `bit == 2`, the carry persists; if `bit == 0`, the carry clears.
- If `bit` is odd (1): two steps (add 1, then divide). This always produces a carry.

After the loop, if `carry == 1`, the MSB has overflowed from `1` to `10`, needing one more divide step.

**Walkthrough with `"1101"` (decimal 13):**

| Position (i) | s[i] | carry | bit | Steps added | Total steps | New carry |
|---|---|---|---|---|---|---|
| 3 | 1 | 0 | 1 (odd) | +2 | 2 | 1 |
| 2 | 0 | 1 | 1 (odd) | +2 | 4 | 1 |
| 1 | 1 | 1 | 2 (even) | +1 | 5 | 1 |
| MSB | 1 | 1 | 2 (overflow) | +1 | 6 | — |

Result: **6 steps** ✓

## Complexity Analysis

Time Complexity: O(n), where n is the length of the binary string. We scan each bit exactly once.

Space Complexity: O(1). Only a few integer variables are used regardless of input size.

## Edge Cases

- **Already "1"**: The number is already 1, so the answer is 0. The loop body doesn't execute since there are no bits to process beyond the MSB.
- **"10" (power of 2)**: Only one divide step is needed. The trailing `0` adds one step and no carry propagates.
- **All ones (e.g. "1111")**: Every bit triggers an add-then-divide (2 steps) with carry propagation, plus a final step for the MSB overflow. This is the worst case per bit.
- **Very long strings (up to 500 characters)**: The O(n) approach handles this efficiently without any big-integer conversion.
