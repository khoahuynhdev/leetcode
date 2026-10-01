# 0020. Valid Parentheses

[LeetCode Link](https://leetcode.com/problems/valid-parentheses/)

Difficulty: Easy
Topics: String, Stack, Bracket Sequences
Acceptance Rate: 45.1%

## Hints

### Hint 1

Think about the order in which brackets must be closed. When you see a closing bracket, which opening bracket must it match: the first one you saw, or the most recent one that hasn't been closed yet? What data structure naturally gives you "the most recent unmatched item"?

### Hint 2

Use a **stack**. Opening brackets go onto the stack. A closing bracket has to match whatever is on top of the stack right now.

### Hint 3

For each closing bracket, check that the stack is non-empty **and** that its top is the matching opener. If so, pop it. If not, the string is invalid right away. When you reach the end, the string is valid only if the stack is **empty**, because leftover openers were never closed. A map from closing to opening bracket keeps the matching check short.

## Approach

Brackets nest. The innermost open bracket has to close first, which is Last-In-First-Out behavior, so a stack fits.

1. Build a lookup `pairs` that maps each closing bracket to its opener: `')' -> '('`, `']' -> '['`, `'}' -> '{'`.
2. Go through the string one character at a time:
   - If the character is an opening bracket, push it onto the stack.
   - If it is a closing bracket:
     - If the stack is empty, there is nothing to match it with. Return `false`.
     - If the top of the stack is not `pairs[c]`, the brackets are interleaved or of the wrong type. Return `false`.
     - Otherwise pop the top. This pair is matched.
3. After the loop, return `true` only if the stack is empty.

**Example:** `s = "([)]"`

| char | action          | stack  |
|------|-----------------|--------|
| `(`  | push            | `(`    |
| `[`  | push            | `( [`  |
| `)`  | top is `[`, need `(`, so mismatch | return false |

**Example:** `s = "([])"`

| char | action             | stack  |
|------|--------------------|--------|
| `(`  | push               | `(`    |
| `[`  | push               | `( [`  |
| `]`  | top `[` matches, pop | `(`  |
| `)`  | top `(` matches, pop | empty |

The stack ends empty, so the result is `true`.

**Small optimization:** a valid string always has even length, so an odd-length input can return `false` immediately.

## Complexity Analysis

Time Complexity: O(n). Each character is pushed and popped at most once.
Space Complexity: O(n). In the worst case the stack holds every character, for example `"(((((("`.

## Edge Cases

- **Only closing brackets** (`"]"`, `"))"`): the stack is empty when a closer arrives. Check for this before reading the top, or you will index out of range.
- **Only opening brackets** (`"(("`, `"{"`): no closer ever fails, but the stack is not empty at the end. You need the final emptiness check.
- **Odd length** (`"(()"`): this can never be valid, so the length check is a cheap early exit.
- **Interleaved brackets** (`"([)]"`): the counts of each type are balanced, but the order is wrong. This is why separate counters for each type don't work and a stack is needed.
- **Deep nesting** (`"{[()]}"`): makes sure the stack correctly handles several levels.
- **Opens after a matched prefix** (`"()("`): the early part matches, but the final check still catches the leftover opener.
