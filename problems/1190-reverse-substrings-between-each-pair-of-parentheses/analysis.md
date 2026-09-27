# 1190. Reverse Substrings Between Each Pair of Parentheses

[LeetCode Link](https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/)

Difficulty: Medium
Topics: String, Stack, Bracket Sequences
Acceptance Rate: 72.9%

## Hints

### Hint 1

Whenever you see nested, balanced brackets that must be processed "innermost first", think about a **stack**. Matching each `(` with its `)` is the classic stack job.

### Hint 2

A straightforward approach: keep a stack of partially built strings (or indices). When you hit `)`, reverse the segment since the most recent `(`. This is O(n²) in the worst case but passes easily for n ≤ 2000. Can you avoid doing the reversals physically at all?

### Hint 3

Precompute `pair[i]` — the index of the matching parenthesis for every bracket (one stack pass). Now walk the string with a pointer `i` and a direction `d = +1`. When you land on a parenthesis, **jump to its partner and flip the direction** (`i = pair[i]; d = -d`), then keep stepping. Letters are appended as you visit them. Each character is visited exactly once, so the walk is O(n).

## Approach

**Why the "wormhole" trick works:** Reversing a segment means reading it from right to left. When we reach an opening `(` while moving forward, the segment inside needs to be read backwards — so we teleport to the matching `)` and start moving left. Nested pairs inside are handled the same way: hitting a bracket while moving left teleports us to its partner and flips direction again (two reversals cancel out). When we finally reach the original `(` again from the inside, we teleport back to the `)` and continue forward past it.

**Algorithm:**

1. **Pair brackets.** Scan `s` with a stack of indices. Push on `(`; on `)`, pop `j` and set `pair[i] = j`, `pair[j] = i`.
2. **Walk.** Start at `i = 0`, `d = 1`. While `0 <= i < n`:
   - If `s[i]` is a bracket: `i = pair[i]`, `d = -d`.
   - Otherwise append `s[i]` to the result.
   - Then `i += d`.
3. Return the result.

**Example:** `s = "(u(love)i)"` (indices 0..9, pairs: 0↔9, 2↔7)

- i=0 `(` → jump to 9, d=-1 → i=8 `i` → append "i"
- i=7 `)` → jump to 2, d=+1 → i=3..6 → append "love"
- i=7 `)` → jump to 2, d=-1 → i=1 `u` → append "u"
- i=0 `(` → jump to 9, d=+1 → i=10, stop.

Result: `"iloveu"` ✓

A simpler stack-of-builders solution (reverse on each `)`) is also perfectly acceptable in an interview; mention the O(n) wormhole idea as the optimization.

## Complexity Analysis

Time Complexity: O(n) — one pass to pair brackets, one walk visiting each character a bounded number of times.
Space Complexity: O(n) — the `pair` array, the index stack, and the output buffer.

## Edge Cases

- **No parentheses** (`"abc"`): the string is returned unchanged.
- **Empty parentheses** (`"()"` or `"a()b"`): contribute nothing; the walk jumps in and straight back out.
- **Deep nesting** (`"((ab))"`): even number of reversals cancels out → `"ab"`.
- **Adjacent groups** (`"(ab)(cd)"`): each group is reversed independently → `"badc"`.
- **Letters outside brackets** (`"a(bc)d"`): only the inner segment is reversed → `"acbd"`.
