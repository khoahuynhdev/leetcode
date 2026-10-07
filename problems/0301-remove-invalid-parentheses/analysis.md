# 0301. Remove Invalid Parentheses

[LeetCode Link](https://leetcode.com/problems/remove-invalid-parentheses/)

Difficulty: Hard
Topics: String, Backtracking, Breadth-First Search
Acceptance Rate: 50.8%

## Hints

### Hint 1

Before you try to build any answers, ask yourself: can you figure out *how many* characters must be removed, without knowing *which* ones? A single left-to-right scan with a counter (like a stack of depth only) tells you a lot about validity.

### Hint 2

Scan once while tracking the open-paren balance. Every `)` that arrives when the balance is zero is a `)` that *must* be removed; every `(` still open at the end is a `(` that *must* be removed. These two counts give you the exact minimum. Now you only need to explore removals that respect those budgets — a natural fit for backtracking (or BFS level-by-level, if you prefer).

### Hint 3

Do a DFS over the string carrying `(index, openBalance, leftToRemove, rightToRemove)`. At each parenthesis, branch on "remove it" (only if that type's budget is > 0) and "keep it" (only if keeping a `)` doesn't make the balance negative). When you reach the end with balance 0 and both budgets at 0, you have a valid answer with the minimum removals. A set handles duplicates like removing either of two adjacent `)` characters.

## Approach

**Step 1 — compute the minimum removals.**
Walk the string with a counter `leftRem`:
- On `(`, increment `leftRem`.
- On `)`, if `leftRem > 0` it matches an earlier `(`, so decrement; otherwise this `)` can never be matched, so increment `rightRem`.

At the end, `leftRem` is the number of unmatched `(` and `rightRem` the number of unmatched `)`. No valid string can be made with fewer than `leftRem + rightRem` removals, and that many is always enough.

**Step 2 — backtrack with budgets.**
`dfs(i, open, lRem, rRem)` builds the candidate in a shared byte buffer:
- If `lRem + rRem` exceeds the characters left, prune — we can't spend the budget.
- At the end of the string, if `open == 0`, the buffer is a valid string with exactly the minimum removals; add it to a set.
- For a parenthesis at `s[i]`, try removing it if the matching budget is positive.
- Then try keeping `s[i]`: letters always kept; `(` increases `open`; `)` only allowed when `open > 0`.

Because we never let `open` go negative and finish only with `open == 0`, every result is valid. Because we spend exactly `leftRem + rightRem` removals, every result is minimal.

**Example:** `s = "()())()"`. The scan gives `leftRem = 0`, `rightRem = 1`. The DFS may only delete one `)`. Deleting the `)` at index 1 gives `(())()`; deleting index 3 or 4 gives `()()()` (twice — the set deduplicates). Deleting index 6 would leave `()())(`, which fails the balance check. Result: `["(())()", "()()()"]`.

**Alternative (BFS):** Start from `s`, generate all strings with one character removed, and stop at the first level that contains any valid string. Simpler to reason about, but explores many more strings than the budget-guided DFS.

## Complexity Analysis

Time Complexity: O(2^n · n) in the worst case — each parenthesis can be kept or removed, and building each result costs O(n). The budget and balance pruning cut this down dramatically in practice (n ≤ 25 with at most 20 parentheses).
Space Complexity: O(n) for the recursion depth and buffer, plus the output size.

## Edge Cases

- **Already valid string** (e.g. `"(a)(b)"`): both budgets are 0, so the only result is the input itself.
- **No parentheses at all** (e.g. `"abc"`): return the string unchanged — letters are never removed.
- **Everything must go** (e.g. `")("`): the answer is `[""]`, not an empty list. Make sure the empty string gets emitted.
- **Unmatched on both sides** (e.g. `")a("`): need to remove both a `)` and a `(`; budgets must be tracked separately.
- **Duplicate results** from removing different but identical adjacent characters (e.g. `"())"`): use a set (or skip-duplicate logic) so each string appears once.
- **Closing paren before any opening** (e.g. `")(()"`): the balance check must reject keeping a `)` when `open == 0`.
