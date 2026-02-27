# 0762. Prime Number of Set Bits in Binary Representation

[LeetCode Link](https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/)

Difficulty: Easy
Topics: Math, Bit Manipulation
Acceptance Rate: 73.2%

## Hints

### Hint 1

Think about what tools you have for counting bits in a number. Once you know the bit count, you just need a fast way to check a property of that count. What property? And how large can that count actually get given the constraints?

### Hint 2

Since numbers go up to 10^6 (roughly 2^20), the number of set bits is at most 20. You only need to determine which values from 0 to 20 are prime. Consider pre-computing or hardcoding that small set of primes so the check is O(1).

### Hint 3

Use a language built-in (like `bits.OnesCount` in Go) to count set bits efficiently, then check membership in the set {2, 3, 5, 7, 11, 13, 17, 19}. A bitmask where bit `i` is set when `i` is prime gives you a constant-time lookup: `(1 << bitCount) & mask != 0`.

## Approach

1. **Identify the prime set:** Since `right <= 10^6 < 2^20`, any number in the range has at most 20 set bits. The primes up to 20 are {2, 3, 5, 7, 11, 13, 17, 19}.

2. **Build a bitmask for fast lookup:** Encode those primes into a single integer where bit `p` is set for each prime `p`:
   ```
   mask = (1<<2) | (1<<3) | (1<<5) | (1<<7) | (1<<11) | (1<<13) | (1<<17) | (1<<19)
   ```
   This equals `665772`.

3. **Iterate and count:** For each number `i` in `[left, right]`:
   - Count its set bits using `bits.OnesCount(uint(i))`.
   - Check if `(mask >> bitCount) & 1 == 1`.
   - If so, increment the result counter.

**Example walkthrough (left=6, right=10):**
- 6 = 110 → 2 bits → 2 is prime ✓
- 7 = 111 → 3 bits → 3 is prime ✓
- 8 = 1000 → 1 bit → 1 is not prime ✗
- 9 = 1001 → 2 bits → 2 is prime ✓
- 10 = 1010 → 2 bits → 2 is prime ✓
- Result: 4

## Complexity Analysis

Time Complexity: O(n) where n = right - left + 1. Each number requires O(1) for popcount and O(1) for the prime check.

Space Complexity: O(1). Only a fixed-size bitmask and a counter are used.

## Edge Cases

- **left == right:** Only one number to check. The loop still works correctly.
- **All numbers have prime bit counts:** Every number in the range qualifies (e.g., left=6, right=7).
- **No numbers have prime bit counts:** For example, powers of two greater than 2 have exactly 1 set bit, which is not prime.
- **Maximum range (right = 10^6):** The solution handles this efficiently since popcount and bitmask lookup are both O(1).
- **Numbers with 0 or 1 set bits:** 0 is not prime, 1 is not prime. These are correctly excluded by the bitmask.
