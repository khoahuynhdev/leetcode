# 3498. Reverse Degree of a String

[LeetCode Link](https://leetcode.com/problems/reverse-degree-of-a-string/)

Difficulty: Easy
Topics: String, Simulation
Acceptance Rate: 89.7%

## Hints

### Hint 1

There is no clever data structure hiding here. The problem hands you a formula and
asks you to evaluate it. Ask yourself: can every term of the sum be computed while
looking at exactly one character, without needing any other character? If so, a
single left-to-right pass is all you need.

### Hint 2

Two numbers make up each product: the letter's rank in the reversed alphabet, and
the letter's 1-indexed position in the string. The second one comes for free from
the loop counter — just remember Go indexes from 0 while the problem counts from 1.
The first one is a small arithmetic mapping from a byte to a number in `[1, 26]`.

### Hint 3

The reversed-alphabet rank needs no lookup table. `'a'` must map to 26 and `'z'` to
1, which is exactly `'z' - c + 1` (equivalently `26 - (c - 'a')`). Plug that into a
running accumulator and the whole solution is one loop with one multiply-add per
character.

## Approach

Iterate over the string once with an index `i` and a byte `c`.

1. **Convert the character to its reversed-alphabet rank.** The normal rank of a
   lowercase letter is `c - 'a' + 1` (so `'a'` = 1, `'z'` = 26). Reversing that
   range means subtracting from 27: `27 - (c - 'a' + 1)` = `26 - (c - 'a')` =
   `int('z' - c) + 1`. Any of these forms works; the last one avoids a magic 26.

2. **Compute the 1-indexed string position.** Go's `for i, c := range s` gives a
   0-based `i`, so the position is `i + 1`.

3. **Accumulate the product.** Add `rank * (i + 1)` into a running total.

4. **Return the total.**

Walk through `s = "zaza"`:

| i | c   | rank (`'z'-c+1`) | position (`i+1`) | product | running total |
|---|-----|------------------|------------------|---------|---------------|
| 0 | `z` | 1                | 1                | 1       | 1             |
| 1 | `a` | 26               | 2                | 52      | 53            |
| 2 | `z` | 1                | 3                | 3       | 56            |
| 3 | `a` | 26               | 4                | 104     | 160           |

The answer is 160, matching the expected output.

Why it works: the reverse degree is defined as a plain sum of independent per-character
terms, so no ordering trick, prefix sum, or memoization can beat simply evaluating each
term once. A single pass is already optimal.

One implementation note for Go: `for i, c := range s` decodes UTF-8 and yields `c` as a
`rune`. That is harmless for this problem because the input is guaranteed ASCII lowercase,
but indexing bytes with `s[i]` is the more direct expression of "one byte per character"
and avoids any decoding subtlety. Either is fine here.

This one is genuinely easy — the 89.7% acceptance rate is honest. The only real ways to
lose points are an off-by-one on the 1-indexing or getting the alphabet reversal backwards.
The value is in writing it cleanly on the first try.

## Complexity Analysis

Time Complexity: O(n), where n is the length of `s`. Each character is visited once and
does O(1) arithmetic.
Space Complexity: O(1). Only a running integer accumulator is kept; no lookup table or
copy of the string is needed.

## Edge Cases

- **Single character (`"a"`, `"z"`).** The constraints guarantee `1 <= s.length`, so the
  string is never empty, but a length-1 input still exercises the 1-indexing: `"a"` must
  give 26 (not 0), and `"z"` must give 1 (not 0). Both catch a 0-indexed position bug.
- **Alphabet endpoints.** `'a'` → 26 and `'z'` → 1 are the boundaries of the reversal
  mapping. If the mapping is accidentally left un-reversed, `"a"` returns 1 instead of 26,
  which a test on either endpoint exposes immediately.
- **Repeated characters (`"zzz"`, `"aaa"`).** The same letter contributes different
  products at different positions, confirming the position factor is applied per
  occurrence and not cached per letter.
- **Maximum length (1000 characters).** The worst case is 1000 `'a'`s:
  26 * (1 + 2 + ... + 1000) = 26 * 500500 = 13,013,000. That fits comfortably in an `int`
  (even a 32-bit one), so no overflow handling is needed — but it is worth confirming the
  bound rather than assuming it.
- **Empty string.** Excluded by the constraints, though a natural loop returns 0 for it,
  which is the sensible answer anyway.
