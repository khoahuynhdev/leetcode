# 3121. Count the Number of Special Characters II

[LeetCode Link](https://leetcode.com/problems/count-the-number-of-special-characters-ii/)

Difficulty: Medium
Topics: Hash Table, String

Acceptance Rate: 51.0%

## Hints

### Hint 1

The string only contains the 26 English letters in two cases each. That's a strong signal that you can summarize everything you need to know about the input in a fixed-size table indexed by letter, rather than re-scanning the string for every letter. Think about what single piece of information per letter would let you decide "special or not" in one glance.

### Hint 2

The condition "every lowercase occurrence appears before the first uppercase occurrence" is really a statement about *positions*. For each letter you don't need every position — you only need to compare the **latest** place a lowercase version appeared against the **earliest** place an uppercase version appeared. If the latest lowercase still comes before the first uppercase, all lowercase ones do.

### Hint 3

Do a single pass and record, for each of the 26 letters, two numbers: the index of the **last** lowercase occurrence and the index of the **first** uppercase occurrence. A letter is special exactly when both indices exist and `lastLower < firstUpper`. The "first uppercase" is captured by only writing the uppercase index the first time you see it; the "last lowercase" is captured by overwriting every time you see it.

## Approach

We exploit the fact that a letter is "special" based purely on relative positions:

- It must appear in **both** cases.
- The **last** lowercase occurrence must come before the **first** uppercase occurrence. (If the last lowercase is before the first uppercase, then necessarily *every* lowercase is, so checking the last one is sufficient.)

Keep two integer arrays of length 26:

- `lastLower[c]` — index of the most recent lowercase occurrence of letter `c`, or a sentinel meaning "never seen".
- `firstUpper[c]` — index of the earliest uppercase occurrence of letter `c`, or a sentinel meaning "never seen".

Walk the string once, character by character, tracking the index `i`:

- If `word[i]` is lowercase, set `lastLower[c] = i` (always overwrite, so it ends up holding the last position).
- If `word[i]` is uppercase, set `firstUpper[c] = i` only if it hasn't been set yet (so it keeps the earliest position).

After the pass, iterate over the 26 letters and count those where both indices are present and `lastLower[c] < firstUpper[c]`.

Worked example with `word = "aaAbcBC"` (indices shown):

```
a a A b c B C
0 1 2 3 4 5 6
```

- `a`: lastLower = 1, firstUpper = 2 → 1 < 2 ✓
- `b`: lastLower = 3, firstUpper = 5 → 3 < 5 ✓
- `c`: lastLower = 4, firstUpper = 6 → 4 < 6 ✓

Count = 3, matching the expected output.

This works because the only way a lowercase occurrence could violate the rule is by appearing at or after the first uppercase one — and the latest lowercase index is the worst case to check.

## Complexity Analysis

Time Complexity: O(n) — one linear scan of the string plus a constant 26-letter sweep at the end.
Space Complexity: O(1) — two fixed arrays of 26 integers, independent of input length.

## Edge Cases

- **No uppercase (or no lowercase) at all** (e.g. `"abc"`): one of the two indices stays at its sentinel for every letter, so nothing is counted. Returns 0.
- **Uppercase appears before lowercase** (e.g. `"AbBCab"` for letter `a`/`b`): `lastLower >= firstUpper`, so the letter is correctly rejected.
- **A lowercase after the first uppercase** (e.g. `"aAa"`): `lastLower = 2`, `firstUpper = 1`, and `2 < 1` is false, so it is not special — exactly the "every lowercase before first uppercase" rule.
- **Single character** (e.g. `"a"` or `"A"`): a letter needs both cases, so a length-1 string always yields 0.
- **Large input** (up to 2 × 10^5 characters): the single-pass O(n) approach handles the upper bound comfortably without any nested scanning.
