# 2124. Check if All A's Appears Before All B's

[LeetCode Link](https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/)

Difficulty: Easy
Topics: String
Acceptance Rate: 75.6%

## Hints

### Hint 1

Think about what a "valid" string looks like as a whole. Can you describe its shape in terms of blocks of characters?

### Hint 2

A valid string is some (possibly empty) run of `'a'`s followed by some (possibly empty) run of `'b'`s. What single local pattern would break that shape?

### Hint 3

The string is invalid if and only if a `'b'` is immediately followed by an `'a'` somewhere — i.e., the substring `"ba"` appears. One linear scan over adjacent pairs is enough.

## Approach

A string passes exactly when it has the form `a...ab...b`. If any `'a'` comes after any `'b'`, then somewhere between them the characters must switch from `'b'` to `'a'`, meaning there's an adjacent pair `s[i-1] == 'b'` and `s[i] == 'a'`. Conversely, if such a pair exists, that `'a'` comes after a `'b'`, so the answer is false.

Algorithm:
1. Iterate `i` from 1 to `len(s)-1`.
2. If `s[i-1] == 'b' && s[i] == 'a'`, return `false`.
3. If the loop finishes, return `true`.

Example: `"abab"` — pairs are `ab`, `ba`, `ab`. The `ba` at indices 1–2 exposes the violation, so return `false`. For `"aaabbb"` no `ba` pair exists, so return `true`.

(An equivalent one-liner is `!strings.Contains(s, "ba")`. Another equivalent formulation: track whether you've seen a `'b'`; if you then see an `'a'`, return false.)

## Complexity Analysis

Time Complexity: O(n)
Space Complexity: O(1)

## Edge Cases

- **Single character** (`"a"` or `"b"`): no pairs to check; always true.
- **Only `'b'`s** (`"bbb"`): there are no `'a'`s, so the condition holds vacuously — true.
- **Only `'a'`s** (`"aaa"`): no `'b'`s at all — true.
- **`'b'` first then `'a'`** (`"ba"`): smallest failing case.
- **Violation at the very end** (`"aabba"`): make sure the scan covers the last pair.
