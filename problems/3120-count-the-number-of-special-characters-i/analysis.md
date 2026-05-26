# 3120. Count the Number of Special Characters I

[LeetCode Link](https://leetcode.com/problems/count-the-number-of-special-characters-i/)

Difficulty: Easy
Topics: Hash Table, String

Acceptance Rate: 71.0%

## Hints

### Hint 1

A letter is "special" only when you have seen *two* facts about it: that it appeared in lowercase **and** that it appeared in uppercase. Whenever a problem asks "does X have property A and also property B", think about tracking membership in two separate sets and then looking at the overlap. A fixed-size lookup table works well here because the alphabet is tiny (only 26 letters).

### Hint 2

You don't need to remember *how many* times a letter showed up, only *whether* it showed up at all in each case. So make a single pass over the string and record, for each of the 26 letters, two booleans: "seen lowercase" and "seen uppercase". After the pass, the answer is just the count of letters where both booleans are true.

### Hint 3

Use two boolean arrays (or two 26-bit bitmasks), `lower[26]` and `upper[26]`. While scanning, set `lower[c-'a'] = true` for lowercase characters and `upper[c-'A'] = true` for uppercase characters. Then loop `i` from 0 to 25 and increment the answer whenever `lower[i] && upper[i]`. The overlap of the two sets is exactly the set of special letters.

## Approach

The core observation is that "special" depends on two independent conditions for the same letter: it must appear at least once in lowercase and at least once in uppercase. Because there are only 26 distinct letters, we can afford a constant-size table instead of a general hash map.

Algorithm, step by step:

1. Create two arrays of 26 booleans: `lower` and `upper`. Index `i` corresponds to the letter `'a'+i` (equivalently `'A'+i`).
2. Make one pass over `word`. For each character `c`:
   - If `c` is between `'a'` and `'z'`, mark `lower[c-'a'] = true`.
   - If `c` is between `'A'` and `'Z'`, mark `upper[c-'A'] = true`.
3. After the pass, iterate over all 26 indices. For each index `i` where both `lower[i]` and `upper[i]` are true, the letter is special, so add 1 to the count.
4. Return the count.

Why it works: marking booleans is idempotent, so duplicate appearances ("aa" vs "a") don't change the result — we only care about presence, never frequency. The final intersection of the two sets is precisely the letters that satisfy both conditions.

Worked example for `word = "aaAbcBC"`:
- Lowercase seen: `a, b, c`.
- Uppercase seen: `A (→a), B (→b), C (→c)`.
- Both lowercase and uppercase: `a, b, c` → 3 special letters. ✓

The repeated `a` at the start contributes nothing extra, which is exactly the behavior we want.

## Complexity Analysis

Time Complexity: O(n), where n is the length of `word`. One pass to fill the tables plus a fixed 26-iteration pass to count, and 26 is a constant.

Space Complexity: O(1). The two boolean arrays have a fixed size of 26 regardless of input length.

## Edge Cases

- **No uppercase letters at all** (e.g., `"abc"`): every `upper[i]` stays false, so no letter is special and the answer is 0. This matches Example 2.
- **Letters appearing in only one case** (e.g., `"abBCab"`): only `b` appears in both cases, so the answer is 1 even though other letters appear multiple times. This matches Example 3 and confirms frequency is irrelevant.
- **Repeated characters** (e.g., the `"aa"` prefix in Example 1): duplicates must not inflate the count; using booleans rather than counters handles this naturally.
- **Single-character string** (e.g., `"a"` or `"A"`): a length-1 word can never be special because a letter cannot be both lowercase and uppercase in one character; the answer is 0.
- **All 26 letters in both cases**: the maximum possible answer is 26; the counting loop bounds guarantee we never exceed it.
