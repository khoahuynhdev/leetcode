# 0678. Valid Parenthesis String

[LeetCode Link](https://leetcode.com/problems/valid-parenthesis-string/)

Difficulty: Medium
Topics: String, Dynamic Programming, Stack, Greedy, Bracket Sequences
Acceptance Rate: 41.1%

## Hints

### Hint 1

Without the `'*'`, this is the classic "balanced parentheses" check: scan left to right and track how many `'('` are still open. The string is valid if that count never drops below zero and ends at zero. How does the `'*'` change what that single counter can be?

### Hint 2

Each `'*'` has three choices, so trying every combination is exponential. Instead of tracking *one* open count, think about tracking the *set* of open counts that are reachable after each character. Is that set always a contiguous range?

### Hint 3

Keep two numbers: `lo`, the minimum possible number of open `'('`, and `hi`, the maximum possible. On `'('` both go up; on `')'` both go down; on `'*'` `lo` goes down (treat it as `')'`) and `hi` goes up (treat it as `'('`). If `hi` ever drops below 0, there are too many `')'` — fail. Clamp `lo` at 0, since a negative open count is never a valid state. At the end, the string is valid exactly when `lo == 0`.

## Approach

**Greedy range tracking.**

Think of the open-parenthesis balance as a value that must stay `>= 0` at every prefix and end at exactly `0`. Because `'*'` can be `'('`, `')'`, or empty, after any prefix the balance could be one of several values. The key observation is that these reachable values always form a contiguous range `[lo, hi]` (each `'*'` adds -1, 0, or +1, which fills gaps), so we only need to store its two endpoints.

Scan the string:

- `'('`: `lo++`, `hi++`
- `')'`: `lo--`, `hi--`
- `'*'`: `lo--` (use it as `')'`), `hi++` (use it as `'('`)

After each step:

1. If `hi < 0`, even treating every `'*'` as `'('` cannot cover the `')'` seen so far → return `false`.
2. If `lo < 0`, set `lo = 0`. A negative balance is an invalid state, so we discard it; this corresponds to choosing "empty" instead of `')'` for some earlier `'*'`.

At the end, return `lo == 0`: zero must be inside the reachable range (`hi >= 0` is already guaranteed).

**Example:** `s = "(*))"`

| char | lo | hi |
|------|----|----|
| `(`  | 1  | 1  |
| `*`  | 0  | 2  |
| `)`  | 0 (clamped from -1) | 1 |
| `)`  | 0 (clamped from -1) | 0 |

`lo == 0` → `true`. (Here the `'*'` acts as `'('`.)

**Alternatives worth knowing:**
- *Two stacks* (indices of `'('` and of `'*'`): match each `')'` with a `'('` first, else a `'*'`; then pair leftover `'('` with `'*'` that appear *after* them. O(n) time, O(n) space.
- *Interval DP* `dp[i][j]` = whether `s[i..j]` is valid: O(n³) time — fine for n ≤ 100 but much slower.

## Complexity Analysis

Time Complexity: O(n) — a single pass over the string.
Space Complexity: O(1) — only two integer counters.

## Edge Cases

- **Leading `')'`** (e.g. `")("`): `hi` goes negative immediately; must fail fast rather than wait for a later `'('` to "fix" it.
- **Only stars** (e.g. `"***"`): every star can be empty, so it's valid; `lo` stays clamped at 0.
- **Unmatched `'('` with a star before it** (e.g. `"*("`): a `'*'` can't close a `'('` that comes *after* it; `lo` ends at 1 → `false`. This is the order constraint that naive counting misses.
- **Single character**: `"("` and `")"` are invalid, `"*"` is valid.
- **Stars needed as closers** (e.g. `"((*"`): only one star for two opens → `lo` ends at 1 → `false`.
