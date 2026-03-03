# 1545. Find Kth Bit in Nth Binary String

[LeetCode Link](https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/)

Difficulty: Medium
Topics: String, Recursion, Simulation
Acceptance Rate: 70.7%

## Hints

### Hint 1

Think about the structure of Sn. Each string is built from the previous one with a predictable pattern. What are the three "sections" of Sn, and how does knowing which section k falls into help you?

### Hint 2

The length of Sn is `2^n - 1`. The middle element is always `'1'`. The left half is exactly S(n-1), and the right half is the reverse-inverse of S(n-1). Can you use this to recurse without ever building the string?

### Hint 3

If k is in the right half of Sn, you can map it back to a position in S(n-1). Specifically, position k in the right half corresponds to position `2^n - k` in S(n-1), but with the bit flipped. This gives you an O(n) recursive solution that peels off one level per call.

## Approach

The key observation is that Sn has a recursive structure we can exploit directly:

```
Sn = S(n-1) + "1" + reverse(invert(S(n-1)))
```

The length of Sn is `2^n - 1`, and the middle position is `mid = 2^(n-1)`.

For any query `(n, k)`, there are three cases:

1. **k == mid**: The middle bit is always `'1'`.
2. **k < mid**: The bit is in the left half, which is exactly S(n-1). So we recurse with `(n-1, k)`.
3. **k > mid**: The bit is in the right half, which is `reverse(invert(S(n-1)))`. Position `k` in the right half maps to position `2^n - k` in S(n-1), and the bit is inverted. So we recurse with `(n-1, 2^n - k)` and flip the result.

**Walkthrough with Example 2 (n=4, k=11):**

- S4 has length 15, mid=8. k=11 > 8, so look at position `16-11=5` in S3, inverted.
- S3 has length 7, mid=4. k=5 > 4, so look at position `8-5=3` in S2, inverted.
- S2 has length 3, mid=2. k=3 > 2, so look at position `4-3=1` in S1, inverted.
- S1, k=1: base case returns `'0'`.
- Unwind: invert `'0'` → `'1'`, invert `'1'` → `'0'`, invert `'0'` → `'1'`.
- Answer: `'1'`.

## Complexity Analysis

Time Complexity: O(n) — each recursive call reduces n by 1, and we do O(1) work per call.

Space Complexity: O(n) — recursion stack depth is at most n. This can be reduced to O(1) with an iterative approach.

## Edge Cases

- **n=1**: The only string is `"0"`, so the answer is always `'0'`. This is the base case of the recursion.
- **k at the exact middle**: The middle bit of any Sn (for n > 1) is always `'1'`.
- **k=1 (first bit)**: The first bit is always `'0'` for any n, since S1 starts with `'0'` and the left half is always preserved.
- **k = 2^n - 1 (last bit)**: The last bit of Sn is always `'1'` for n > 1, since it's the inversion of the first bit of S(n-1) which is `'0'`.
- **Large n (n=20)**: The string would have over 1 million characters, but our recursive approach handles it in just 20 steps.
