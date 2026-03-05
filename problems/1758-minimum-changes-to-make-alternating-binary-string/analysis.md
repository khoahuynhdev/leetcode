# 1758. Minimum Changes To Make Alternating Binary String

[LeetCode Link](https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/)

Difficulty: Easy
Topics: String
Acceptance Rate: 64.5%

## Hints

### Hint 1

Think about what a valid alternating binary string looks like. How many possible alternating strings of a given length exist?

### Hint 2

There are only two possible alternating patterns for any length: one starting with '0' (like "0101...") and one starting with '1' (like "1010..."). Compare the input against both patterns and count mismatches.

### Hint 3

You don't need to actually build the two target strings. At each index `i`, the expected character for the "starts with 0" pattern is `'0'` if `i` is even and `'1'` if `i` is odd. Count mismatches for one pattern, and the mismatches for the other pattern is simply `len(s) - count`. Return the minimum of the two.

## Approach

Since an alternating binary string has no two adjacent characters equal, there are exactly two valid alternating strings of length `n`:
- Pattern A: `"010101..."` (starts with '0')
- Pattern B: `"101010..."` (starts with '1')

These two patterns are complements of each other -- wherever pattern A has a '0', pattern B has a '1', and vice versa. This means if a character in `s` mismatches pattern A, it must match pattern B, and the other way around.

We iterate through the string once, counting how many positions mismatch pattern A (starts with '0'). Call this `countA`. Then the number of mismatches against pattern B is `len(s) - countA`. The answer is `min(countA, len(s) - countA)`.

For example, with `s = "0100"`:
- Pattern A: `"0101"` -- mismatches at index 3 -> countA = 1
- Pattern B: `"1010"` -- mismatches at indices 0, 1, 2 -> countB = 3
- Answer: min(1, 3) = 1

## Complexity Analysis

Time Complexity: O(n), where n is the length of the string. We make a single pass.
Space Complexity: O(1), only a counter variable is used.

## Edge Cases

- **Already alternating**: The string may already be valid (e.g., "10"), requiring 0 changes. Both patterns are checked so this is handled naturally.
- **All same characters**: A string like "1111" requires changes to half the characters (for even length) or (n-1)/2 or (n+1)/2 depending on the pattern chosen.
- **Single character**: Any single character string is already alternating, so the answer is 0.
