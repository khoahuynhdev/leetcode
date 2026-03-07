# 1888. Minimum Number of Flips to Make the Binary String Alternating

[LeetCode Link](https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/)

Difficulty: Medium
Topics: String, Dynamic Programming, Sliding Window
Acceptance Rate: 42.6%

## Hints

### Hint 1

The Type-1 operation (moving the first character to the end) is essentially a rotation of the string. Think about how you can efficiently evaluate all possible rotations without actually performing them one by one.

### Hint 2

A common trick for rotation problems is to double the string (concatenate `s` with itself). Every contiguous substring of length `n` in the doubled string represents a valid rotation of the original. Now you need to find which rotation requires the fewest Type-2 flips. Consider using a sliding window of size `n` over the doubled string.

### Hint 3

There are only two possible alternating target patterns for a string of length `n`: one starting with `'0'` and one starting with `'1'`. For the doubled string, precompute the mismatch against both targets. Then slide a window of size `n` across the doubled string, maintaining a running count of mismatches for each target. The answer is the minimum mismatch count across all windows and both targets.

## Approach

1. **Double the string**: Let `n = len(s)`. Create `ss = s + s`. Any substring of length `n` in `ss` corresponds to a rotation of `s`.

2. **Define two targets**: An alternating string of length `2n` can start with either `'0'` or `'1'`:
   - Target A: `010101...`
   - Target B: `101010...`

3. **Sliding window over the doubled string**: Maintain two counters (`diffA` and `diffB`) tracking how many positions in the current window of size `n` differ from target A and target B respectively.

4. **Initialize**: Count mismatches for the first window (indices `0` to `n-1`).

5. **Slide**: For each new position `i` from `n` to `2n-1`:
   - Add the mismatch at position `i` (the character entering the window).
   - Remove the mismatch at position `i-n` (the character leaving the window).
   - Update the minimum across both `diffA` and `diffB`.

6. **Return** the minimum value found.

This works because each window position simulates a different number of Type-1 rotations, and the mismatch count within that window is exactly the number of Type-2 flips needed.

## Complexity Analysis

Time Complexity: O(n) — we iterate through the doubled string once with the sliding window.

Space Complexity: O(n) — for the doubled string. This can be reduced to O(1) extra space by using modular indexing instead.

## Edge Cases

- **Already alternating**: The string needs zero flips (e.g., `"010"` or `"1010"`). The algorithm naturally returns 0.
- **Single character**: Any single character string is trivially alternating; answer is 0.
- **All same characters**: e.g., `"0000"` or `"1111"`. The algorithm correctly computes the minimum flips needed across all rotations.
- **Even vs odd length**: For even-length strings, rotations don't change the parity alignment, so Type-1 operations have limited effect. For odd-length strings, rotations can shift parity, making Type-1 operations more impactful. The sliding window handles both cases uniformly.
