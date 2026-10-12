# 2897. Apply Operations on Array to Maximize Sum of Squares

[LeetCode Link](https://leetcode.com/problems/apply-operations-on-array-to-maximize-sum-of-squares/)

Difficulty: Hard
Topics: Array, Hash Table, Greedy, Bit Manipulation
Acceptance Rate: 62.2%

## Hints

### Hint 1

Don't try to simulate the operations. Instead, ask: what quantity does the operation *preserve*? Look at the effect of `(a AND b, a OR b)` one bit position at a time.

### Hint 2

For any single bit, check all four cases of `(bit of a, bit of b)`. You'll find that the number of elements with that bit set never changes; the bit can only move from one element to the other. So the array's state is really just "how many of each bit there are."

### Hint 3

Squares are convex: `(x + y)^2 >= x^2 + y^2`. So you want to pile as many bits as possible onto one number, then as many of the leftover bits as possible onto the next number, and so on. Count each bit's frequency, then greedily build the `k` largest numbers.

## Approach

1. **Invariant.** For a fixed bit position, the operation maps `(a, b)` bits as:
   - `(0,0) -> (0,0)`, `(1,1) -> (1,1)`
   - `(1,0) -> (0,1)`, `(0,1) -> (0,1)`

   In every case the number of 1s at that position is unchanged. So `cnt[b]` (how many elements have bit `b` set) is fixed no matter what operations you do.

2. **Reachability.** Repeatedly applying the operation lets you move bits around freely: in effect, you can make the first element take one copy of every available bit, the second take one copy of every bit that still has copies left, etc. (It's like sorting the bit columns so the 1s sit at the "top" of each column.)

3. **Greedy.** Because `f(x) = x^2` is convex, moving a bit from a smaller number to a bigger number never decreases the sum of squares. So the best plan is to make the largest number as large as possible, then the next, and so on:
   - Count `cnt[b]` for each of the ~30 bit positions.
   - Repeat `k` times: build `x` by taking bit `b` whenever `cnt[b] > 0` (and decrement it). Add `x^2 mod (1e9+7)` to the answer.

**Example:** `nums = [2,6,5,8]`, `k = 2`
- Bit counts: bit0 = 1 (from 5), bit1 = 2 (2, 6), bit2 = 2 (6, 5), bit3 = 1 (8).
- 1st number: take bits 0,1,2,3 → 15. Remaining: bit1 = 1, bit2 = 1.
- 2nd number: take bits 1,2 → 6.
- Answer: 15² + 6² = 225 + 36 = 261.

Note that the numbers being built are at most `2^31 - 1`, so `x * x` fits in a 64-bit `int` before taking the modulo.

## Complexity Analysis

Time Complexity: O((n + k) · B), where B ≈ 30 is the number of bits.
Space Complexity: O(B) for the bit counts.

## Edge Cases

- **k = n:** You'll eventually build numbers that are 0 once all bits are used up; they contribute nothing, which is correct.
- **All elements identical:** No operation changes anything; the answer is just `k * x^2`.
- **Disjoint bits (e.g. `[1,2,4]`):** Everything collapses into one number (7). This shows why the answer can exceed the sum of squares of the original elements.
- **Large values (~1e9):** Squares exceed 32-bit range and the answer must be reduced modulo `1e9+7`. Take the modulo after each addition to avoid overflow.
- **Single element:** No operations possible (needs two distinct indices), answer is `nums[0]^2`.
