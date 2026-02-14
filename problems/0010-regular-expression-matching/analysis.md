# 0010. Regular Expression Matching

[LeetCode Link](https://leetcode.com/problems/regular-expression-matching/)

Difficulty: Hard
Topics: String, Dynamic Programming, Recursion
Acceptance Rate: 30.3%

## Hints

### Hint 1

Think about breaking down the problem into smaller subproblems. When you're comparing a string with a pattern, you're essentially making decisions character by character. What information do you need to know about the rest of the string and pattern after making a decision about the current characters?

### Hint 2

The '*' character is the tricky part - it can match zero or more of the preceding character. This means when you encounter a pattern like "a*", you have two choices: either use the star (match one 'a' and keep the pattern "a*" available), or skip the star pattern entirely (match zero 'a's). Consider using dynamic programming or recursion with memoization to avoid recomputing the same subproblems.

### Hint 3

Define your state as dp[i][j], representing whether s[0...i-1] matches p[0...j-1]. The key insight is handling the star case: if p[j-1] is '*', then dp[i][j] is true if either:
1. We skip the pattern (dp[i][j-2] is true - matching zero occurrences)
2. The preceding character matches (s[i-1] matches p[j-2]) AND dp[i-1][j] is true (we consume one character from s and keep trying to match more with the same pattern)

## Approach

We'll use dynamic programming with a 2D table where dp[i][j] represents whether the first i characters of string s match the first j characters of pattern p.

**Base Cases:**
- dp[0][0] = true (empty string matches empty pattern)
- dp[0][j] can be true only if the pattern can match empty string (patterns like "a*", "a*b*", etc.)

**State Transitions:**

For each position (i, j):

1. **If p[j-1] is not '*'**:
   - dp[i][j] = dp[i-1][j-1] AND (s[i-1] == p[j-1] OR p[j-1] == '.')
   - This means the current characters must match, and the previous substrings must match.

2. **If p[j-1] is '*'**:
   - We have two options:
     - **Zero occurrences**: dp[i][j] = dp[i][j-2] (skip the character and '*')
     - **One or more occurrences**: If s[i-1] matches p[j-2] (the character before '*'), then dp[i][j] = dp[i-1][j] (consume one character from s and keep the pattern)
   - dp[i][j] = dp[i][j-2] OR (matches AND dp[i-1][j])
   - Where matches = (s[i-1] == p[j-2] OR p[j-2] == '.')

**Example Walkthrough:**
For s = "aa" and p = "a*":
- dp[0][0] = true
- dp[0][2] = true (because "a*" can match empty string)
- dp[1][2] = true (first 'a' matches "a*")
- dp[2][2] = true (both 'a's match "a*")

## Complexity Analysis

Time Complexity: O(m * n), where m is the length of the string s and n is the length of the pattern p. We fill a 2D table of size (m+1) x (n+1).

Space Complexity: O(m * n) for the DP table. This can be optimized to O(n) by using only two rows, but the standard approach uses O(m * n).

## Edge Cases

1. **Empty string with star patterns**: Patterns like "a*b*c*" should match an empty string.
2. **Single dot**: "." should match any single character.
3. **Dot with star**: ".*" should match any string (zero or more of any character).
4. **Multiple consecutive star patterns**: "a*a*" is a valid pattern that needs proper handling.
5. **Pattern longer than string**: A pattern like "a*b*c*" can match shorter strings or even empty strings.
6. **Exact match without wildcards**: Basic string comparison when pattern has no special characters.
