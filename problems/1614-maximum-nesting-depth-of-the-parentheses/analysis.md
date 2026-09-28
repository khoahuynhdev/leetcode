# 1614. Maximum Nesting Depth of the Parentheses

[LeetCode Link](https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/)

Difficulty: Easy
Topics: String, Stack, Bracket Sequences
Acceptance Rate: 85.4%

## Hints

### Hint 1

This is a bracket-matching problem. Think about the classic tool for tracking open brackets as you scan left to right. The digits and operators are just noise.

### Hint 2

A stack of `'('` works: push on `'('`, pop on `')'`. The stack's size at any moment is how deeply nested you are right now. What do you need to record while scanning?

### Hint 3

You never look at what's *inside* the stack, only at its size. So you can drop the stack and keep a single integer counter: `+1` on `'('`, `-1` on `')'`. The answer is the largest value that counter ever reaches.

## Approach

Scan the string once and keep two integers:

- `depth`: how many parentheses are currently open
- `best`: the largest `depth` seen so far

For each character:

1. If it's `'('`, increment `depth` and update `best = max(best, depth)`.
2. If it's `')'`, decrement `depth`.
3. Ignore everything else (digits and operators).

**Why it works:** the input is guaranteed to be a valid parentheses string, so `depth` never goes negative and always equals the number of unmatched `'('` before the current position. That is exactly the nesting depth at that point. The overall nesting depth is the maximum over all positions, and it can only increase right after an `'('`, so that's the only place we need to update `best`.

**Example:** `s = "(1)+((2))+(((3)))"`

| char | depth | best |
|------|-------|------|
| `(`  | 1     | 1    |
| `)`  | 0     | 1    |
| `(`  | 1     | 1    |
| `(`  | 2     | 2    |
| `)`,`)` | 0  | 2    |
| `(`,`(`,`(` | 3 | 3 |
| `)`,`)`,`)` | 0 | 3 |

Result: `3`.

This counter trick is worth remembering. Whenever a bracket problem only cares about *how many* brackets are open and not *which kind*, a counter can replace the stack.

## Complexity Analysis

Time Complexity: O(n), one pass over the string.
Space Complexity: O(1), just two integers.

## Edge Cases

- **No parentheses at all** (e.g. `"1+2"`): the depth stays 0, so return 0. Make sure `best` starts at 0.
- **Single character** (e.g. `"8"`): the minimum length is 1, and the answer is 0.
- **Sibling groups vs. nested groups** (e.g. `"()()()"` vs. `"((()))"`): siblings stay at depth 1 while nested groups reach 3. This checks that you track the *current* depth rather than just counting `'('`.
- **Deep nesting at the end** (e.g. `"1+(2+(3))"`): the maximum can occur anywhere, so update `best` on every `'('`, not only at the end.
