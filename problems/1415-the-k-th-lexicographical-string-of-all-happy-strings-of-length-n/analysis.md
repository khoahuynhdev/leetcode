# 1415. The k-th Lexicographical String of All Happy Strings of Length n

[LeetCode Link](https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/)

Difficulty: Medium
Topics: String, Backtracking
Acceptance Rate: 85.5%

## Hints

### Hint 1

Think about how you would generate all strings of length `n` using only the characters `'a'`, `'b'`, and `'c'`. What constraint makes a string "happy", and how does that constraint affect the number of choices you have at each position?

### Hint 2

At each position in the string, you have exactly 2 valid choices (any character that is not the same as the previous one). This means you can use backtracking to generate all happy strings in lexicographical order. If you always try characters in order `'a'`, `'b'`, `'c'`, the results will naturally be sorted.

### Hint 3

Since the first character has 3 choices and every subsequent character has exactly 2 choices, the total number of happy strings of length `n` is `3 * 2^(n-1)`. You can either generate all of them and pick the k-th one, or use this mathematical structure to directly compute the k-th string without generating them all — each position can be determined by dividing `k` into groups of size `2^(remaining positions)`.

## Approach

**Backtracking approach:** We build happy strings character by character using backtracking. At each step, we try each of the three characters `'a'`, `'b'`, `'c'` in order, skipping any character that matches the previous character. Since we iterate characters in alphabetical order, the generated strings are automatically in lexicographical order.

We maintain a counter that increments each time we complete a string of length `n`. When the counter reaches `k`, we've found our answer.

**Step-by-step walkthrough (n=3, k=9):**
1. Start with empty string, try 'a' first
2. Build: "a" → "ab" → "aba" (count=1), "abc" (count=2)
3. Build: "a" → "ac" → "aca" (count=3), "acb" (count=4)
4. Continue with 'b': "bab" (5), "bac" (6), "bca" (7), "bcb" (8)
5. Continue with 'c': "cab" (count=9) — this is our answer!

The backtracking naturally produces strings in lexicographical order because we always try 'a' before 'b' before 'c'.

## Complexity Analysis

Time Complexity: O(k) — we generate at most k complete happy strings before finding the answer. In the worst case this is O(3 * 2^(n-1)) if k equals the total count.

Space Complexity: O(n) — for the recursion stack depth and the current string being built.

## Edge Cases

- **k exceeds total happy strings:** When `k > 3 * 2^(n-1)`, return an empty string. For example, n=1, k=4 should return "".
- **n = 1:** Only three happy strings exist: "a", "b", "c". No adjacency constraint applies since the string has only one character.
- **k = 1:** Always returns "a...a" alternating pattern — specifically a string starting with 'a' and alternating with 'b'.
- **Maximum constraints (n=10, k=100):** Total happy strings = 3 * 512 = 1536, so k=100 is always valid for n=10.
