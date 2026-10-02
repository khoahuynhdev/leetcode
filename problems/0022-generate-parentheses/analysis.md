# 0022. Generate Parentheses

[LeetCode Link](https://leetcode.com/problems/generate-parentheses/)

Difficulty: Medium
Topics: String, Dynamic Programming, Backtracking, Bracket Sequences
Acceptance Rate: 79.3%

## Hints

### Hint 1

You need to produce *every* valid string, not count them. Whenever a problem asks you to enumerate all combinations, think about building the answer one character at a time and exploring choices recursively.

### Hint 2

Use backtracking: at each position you can either place `(` or `)`. Instead of generating all `2^(2n)` strings and filtering, try to prune invalid branches *while* building. What information do you need to track to know whether a choice is still legal?

### Hint 3

Track two counters: `open` (how many `(` used) and `close` (how many `)` used).
- You may add `(` only if `open < n`.
- You may add `)` only if `close < open` (there is an unmatched `(` to close).

With these two rules, every completed string of length `2n` is automatically valid — no filtering needed.

## Approach

We build strings recursively with a shared byte buffer of length `2n`.

1. Start with an empty buffer, `open = 0`, `close = 0`.
2. If the buffer is full (`open + close == 2n`), record a copy of it as an answer.
3. Otherwise:
   - If `open < n`, write `(` at the current position and recurse with `open + 1`.
   - If `close < open`, write `)` at the current position and recurse with `close + 1`.
4. Because we overwrite the same position on each branch, no explicit "undo" step is needed.

**Why it works:** A parentheses string is well-formed iff every prefix has at least as many `(` as `)`, and the totals are equal. Rule `close < open` enforces the prefix condition; rule `open < n` plus the length `2n` forces the totals to be equal. Every branch we explore leads to at least one valid string, so we never waste work on dead ends.

**Example (n = 2):**

```
""  -> "("  -> "((" -> "(()" -> "(())"
            -> "()" -> "()(" -> "()()"
```

Trying `(` before `)` at each step yields results in lexicographic order, matching the example output.

## Complexity Analysis

Time Complexity: O(4^n / √n) — the number of valid strings is the n-th Catalan number `C(n) ~ 4^n / (n^{3/2}·√π)`, and each takes O(n) to copy into the result.
Space Complexity: O(n) auxiliary (recursion depth and buffer of size `2n`), excluding the O(4^n / √n) output.

## Edge Cases

- **n = 1:** The smallest input; only `"()"` is valid. Make sure the base case triggers correctly.
- **Never closing before opening:** Strings like `")("` must never be generated — guarded by `close < open`.
- **Too many opens:** Strings like `"((("` for n = 2 must be cut off — guarded by `open < n`.
- **Larger n (up to 8):** Output grows to 1430 strings; the pruned backtracking handles this easily, but naive generate-and-filter (`2^16` strings) is wasteful.
- **Buffer reuse:** When reusing a byte slice, convert to `string` at the leaf so each result is an independent copy.
