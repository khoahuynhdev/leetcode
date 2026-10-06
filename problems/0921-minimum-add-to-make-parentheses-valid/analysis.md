# 0921. Minimum Add to Make Parentheses Valid

[LeetCode Link](https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/)

Difficulty: Medium
Topics: String, Stack, Greedy, Bracket Sequences
Acceptance Rate: 74.2%

## Hints

### Hint 1

The classic tool for matching brackets is a stack. Think about what happens when you try to validate the string with a stack. Which characters are left over or fail to match?

### Hint 2

You only have one kind of bracket, so you don't need to store the actual characters on the stack. A single counter of "currently open" parentheses works just as well. Scan left to right and track it.

### Hint 3

Every `)` that shows up when the open counter is already zero can never be matched by anything later. It needs an inserted `(`, so count it right away. Any `(` still open at the end of the scan needs an inserted `)`. The answer is the sum of those two counts.

## Approach

Scan the string once and keep two counters:

- `open`: the number of `(` seen so far that haven't been matched yet.
- `adds`: the number of insertions needed for `)` characters that had nothing to match.

For each character:

1. If it is `(`, increment `open`.
2. If it is `)`:
   - If `open > 0`, it matches one pending `(`, so decrement `open`.
   - Otherwise there's no `(` to its left to match it, so we must insert one. Increment `adds`.

At the end, each of the `open` unmatched `(` needs its own `)`, so the answer is `adds + open`.

**Why this is minimal (greedy argument):** An unmatched `)` at position `i` means the prefix `s[0..i]` has more `)` than `(`. Characters later in the string can't fix that prefix, so at least one `(` must be inserted before it. Likewise, every leftover `(` has no `)` after it to pair with. Each insertion fixes exactly one unmatched character, and matching greedily (each `)` pairs with any pending `(`) never wastes a pair. So the count is a lower bound and also achievable.

**Example:** `s = "())("`

| char | open | adds |
|------|------|------|
| `(`  | 1    | 0    |
| `)`  | 0    | 0    |
| `)`  | 0    | 1    |
| `(`  | 1    | 1    |

Answer = `1 + 1 = 2`. For example, `"(())()"` uses 2 insertions to make it valid.

## Complexity Analysis

Time Complexity: O(n), with a single pass over the string.
Space Complexity: O(1), using only two integer counters.

## Edge Cases

- **Already valid string** (e.g. `"()"`, `"(())"`): the answer should be 0, so the counters must end balanced.
- **All opening** (e.g. `"((("`): nothing ever closes, so the answer is the full length, all coming from the leftover `open`.
- **All closing** (e.g. `")))"`): every `)` is unmatched, so the answer is the full length, all coming from `adds`.
- **Reversed pair** (`")("`): the total counts of `(` and `)` are equal, but order matters, so the answer is 2. Comparing totals alone gives the wrong answer here.
- **Single character**: the minimum input length is 1, and the answer is 1 either way.
