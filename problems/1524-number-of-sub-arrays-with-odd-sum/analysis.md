# 1524. Number of Sub-arrays With Odd Sum

[LeetCode Link](https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/)

Difficulty: Medium
Topics: Array, Math, Dynamic Programming, Prefix Sum
Acceptance Rate: 56.7%

## Hints

### Hint 1

Checking every subarray is O(n²) and too slow for n = 10^5. Any subarray sum can be written as the difference of two **prefix sums**. Can you use that?

### Hint 2

You don't care about the actual sums, only whether they are odd or even. Think about the **parity** of prefix sums, and what happens to parity when you subtract one number from another.

### Hint 3

`sum(i+1..j) = prefix[j] - prefix[i]` is odd exactly when `prefix[j]` and `prefix[i]` have **different** parity. So as you scan, keep a count of how many earlier prefix sums were even and how many were odd (the empty prefix, 0, counts as even). If the current prefix is odd, add the even count to the answer; if it is even, add the odd count.

## Approach

1. Keep a running prefix parity `prefix` (0 = even, 1 = odd), starting at 0.
2. Keep counters `evenCount = 1` (for the empty prefix) and `oddCount = 0`.
3. For each element:
   - Update `prefix = (prefix + x) % 2`.
   - If `prefix` is odd, every earlier even prefix forms an odd-sum subarray ending here: `ans += evenCount`, then `oddCount++`.
   - Otherwise `ans += oddCount`, then `evenCount++`.
   - Take the answer modulo 10^9 + 7.
4. Return `ans`.

**Why it works:** each subarray `arr[i+1..j]` corresponds to exactly one pair of prefixes `(i, j)` with `i < j`. Odd minus even (or even minus odd) is odd, so counting opposite-parity pairs counts exactly the odd-sum subarrays.

**Example:** `arr = [1, 3, 5]`

| x | prefix parity | even/odd before | added | ans |
|---|---------------|-----------------|-------|-----|
| 1 | odd  | 1 / 0 | +1 (even) | 1 |
| 3 | even | 1 / 1 | +1 (odd)  | 2 |
| 5 | odd  | 2 / 1 | +2 (even) | 4 |

The answer is 4, which matches the expected output.

This is also a small DP: let `odd[j]` / `even[j]` be the number of subarrays ending at `j` with odd/even sum. If `arr[j]` is odd they swap (and the single-element subarray adds to odd); if it's even they carry over. You get the same O(1)-space algorithm from a different angle.

## Complexity Analysis

Time Complexity: O(n), one pass over the array.
Space Complexity: O(1), only a few counters.

## Edge Cases

- **Single element:** odd gives 1, even gives 0. Checks that the empty prefix is counted as even from the start.
- **All even numbers:** every prefix is even, so the answer is 0.
- **All odd numbers:** prefix parity alternates, so about n²/4 subarrays are odd. Make sure you use modular arithmetic and don't overflow.
- **Large n (10^5):** the raw count can reach ~2.5·10^9, which is above 10^9 + 7, so the modulo really matters.
