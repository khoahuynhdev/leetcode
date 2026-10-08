# 1021. Remove Outermost Parentheses

[LeetCode Link](https://leetcode.com/problems/remove-outermost-parentheses/)

Difficulty: Easy
Topics: String, Stack, Bracket Sequences
Acceptance Rate: 87.8%

## Hints

### Hint 1

Validating parentheses usually means using a stack. Here, though, every `(` is the same as every other `(`. Do you need to store the characters, or is it enough to know how many are open?

### Hint 2

Keep a counter for the current nesting depth: add one for `(` and subtract one for `)`. Think about what the depth is when you reach the first and last character of each primitive piece.

### Hint 3

A primitive starts with the `(` that takes the depth from 0 to 1, and it ends with the `)` that takes the depth from 1 back to 0. Those are the only two characters you drop. Copy every other character to the output. If you check the depth at the right moment (before increasing it on `(`, after decreasing it on `)`), the rule becomes "keep the character when the depth is greater than 0."

## Approach

Go through the string once and track `depth`, the number of currently open parentheses.

- On `(`: if `depth > 0`, this paren is inside a primitive, so append it. Then increment `depth`. When `depth` was 0, this is the opening paren of a new primitive, and we skip it.
- On `)`: decrement `depth` first. If `depth > 0` after that, append it. When `depth` becomes 0, this is the closing paren of the current primitive, and we skip it.

Why this works: a primitive is a balanced substring where the depth never drops to 0 before its last character. So depth 0 marks the boundaries between primitives, and the outermost parens are exactly the characters at those boundaries.

Example: `s = "(()())(())"`

| char | depth before | action | depth after | output |
|------|--------------|--------|-------------|--------|
| (    | 0            | skip   | 1           |        |
| (    | 1            | keep   | 2           | (      |
| )    | 2            | keep   | 1           | ()     |
| (    | 1            | keep   | 2           | ()(    |
| )    | 2            | keep   | 1           | ()()   |
| )    | 1            | skip   | 0           | ()()   |
| (    | 0            | skip   | 1           | ()()   |
| (    | 1            | keep   | 2           | ()()(  |
| )    | 2            | keep   | 1           | ()()() |
| )    | 1            | skip   | 0           | ()()() |

The result is `"()()()"`.

A full stack would also work, but the counter does the same job in O(1) extra space because only the stack's size is ever used.

## Complexity Analysis

Time Complexity: O(n), with one pass over the string.
Space Complexity: O(n) for the output buffer, and O(1) extra space besides that.

## Edge Cases

- **Only flat primitives (`"()()"`)**: every character is an outer paren, so the answer is the empty string. Make sure you return `""` and not nil or garbage.
- **One deeply nested primitive (`"((()))"`)**: only the first and last characters are removed, which gives `"(())"`. This checks that you drop parens only at depth 0, not at every level.
- **Minimum input (`"()"`)**: the smallest valid input, with output `""`.
- **Many primitives in a row**: the depth goes back to 0 between pieces. Your logic has to handle every boundary, not just the first and last character of the whole string.
- **Order of increment/decrement**: checking the depth on the wrong side of the update gives an off-by-one error that either keeps the outer parens or drops inner ones.
