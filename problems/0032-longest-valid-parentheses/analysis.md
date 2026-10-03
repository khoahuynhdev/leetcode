# 0032. Longest Valid Parentheses

[LeetCode Link](https://leetcode.com/problems/longest-valid-parentheses/)

Difficulty: Hard
Topics: String, Dynamic Programming, Stack, Bracket Sequences
Acceptance Rate: 40.2%

## Hints

### Hint 1

Checking whether a whole string is balanced is a classic **stack** problem. Here you need the longest balanced *substring*, so think about how a stack could also tell you where each valid stretch *starts*.

### Hint 2

Instead of pushing characters onto the stack, push **indices**. When a `)` matches a `(`, the length of the valid run ending here can be computed from an index that is still on the stack.

### Hint 3

Keep a "boundary" index at the bottom of the stack: the position just before the current valid run could begin. Start with `-1` as the boundary. On `(`, push its index. On `)`, pop. If the stack becomes empty, this `)` is unmatched, so push its index as the new boundary. Otherwise, the valid run ending at `i` has length `i - stack.top()`.

## Approach

We scan the string once, keeping a stack of indices.

1. Push `-1` onto the stack. It acts as a sentinel boundary: "the last position that can't be part of a valid substring".
2. For each index `i`:
   - If `s[i] == '('`, push `i`. It might be matched later.
   - If `s[i] == ')'`, pop the top.
     - If the stack is now **empty**, that `)` had nothing to match. It becomes the new boundary, so push `i`.
     - Otherwise, the top of the stack is the index just before the current valid substring, so update `best = max(best, i - top)`.

**Why it works:** At any time, the stack holds the indices of unmatched `(` characters, sitting on top of the index of the most recent unmatched `)` (or `-1`). Everything between the top of the stack and `i` has been matched, so it forms a valid substring.

**Example:** `s = ")()())"`

| i | char | action | stack | best |
|---|------|--------|-------|------|
| - | - | init | [-1] | 0 |
| 0 | ) | pop → empty, push 0 | [0] | 0 |
| 1 | ( | push 1 | [0, 1] | 0 |
| 2 | ) | pop → top 0, len 2 | [0] | 2 |
| 3 | ( | push 3 | [0, 3] | 2 |
| 4 | ) | pop → top 0, len 4 | [0] | 4 |
| 5 | ) | pop → empty, push 5 | [5] | 4 |

Answer: `4`.

**Alternative (O(1) space):** Scan left to right counting `open` and `close`. When they're equal, record `2 * close`. When `close > open`, reset both to zero. That misses cases like `"(()"`, where there are always more opens than closes, so do a second scan right to left with the reset rule flipped (`open > close`). There is also an O(n) DP where `dp[i]` is the length of the longest valid substring ending at `i`.

## Complexity Analysis

Time Complexity: O(n), one pass over the string, with each index pushed and popped at most once.
Space Complexity: O(n) for the stack in the worst case (e.g. `"(((((("`). The two-counter variant uses O(1).

## Edge Cases

- **Empty string**: there's nothing to match, so the answer is 0. The loop simply doesn't run.
- **Only `(` or only `)`**: nothing ever matches, so the answer is 0. The sentinel and boundary logic must not produce false lengths.
- **Leading unmatched `)`**: for example `")()"`. The empty-stack reset makes the `)` the new boundary.
- **Unmatched `(` in the middle or at the start**: for example `"(()"`. The leftover `(` stays on the stack and serves as the boundary for later matches.
- **Nested and adjacent groups combined**: for example `"()(())"`. These must be counted as one run of length 6, which the boundary index handles naturally.
- **Whole string valid**: the sentinel `-1` gives length `n - 1 - (-1) = n`.
