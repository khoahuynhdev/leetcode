# 1541. Minimum Insertions to Balance a Parentheses String

[LeetCode Link](https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/)

Difficulty: Medium
Topics: String, Stack, Greedy, Bracket Sequences
Acceptance Rate: 56.6%

## Hints

### Hint 1

This is a variation of the classic "valid parentheses" problem. In the classic version you can get away with just a counter instead of a real stack. Can you do the same here, even though one `(` now needs **two** closing characters?

### Hint 2

Scan left to right greedily. Keep a counter of how many `)` characters you still *need* to close every `(` you've seen so far. Each `(` adds 2 to that need. Whenever the need can't be satisfied (a `)` shows up with nothing open), fix it on the spot by counting an insertion.

### Hint 3

The tricky part is that `))` must be **consecutive**. When you see a `)`, look at the next character:
- If it's also `)`, you have a full `))` pair — consume both.
- If not, you only have a lone `)`, so you must insert one `)` to complete the pair.

Then, for that `))` pair: if there's an open `(` waiting, match it; otherwise insert a `(`. At the end, every still-open `(` needs two `)` inserted.

## Approach

Treat `))` as a single closing token and process the string greedily from left to right with two counters:

- `open` — number of `(` that are still unmatched.
- `ans` — number of insertions made so far.

Walk with index `i`:

1. If `s[i] == '('`: increment `open`, advance `i` by 1.
2. If `s[i] == ')'`:
   - If `s[i+1]` exists and is `)`, we have a complete `))` token; advance `i` by 2.
   - Otherwise it's a lone `)`; insert a `)` (`ans++`) to form `))`, and advance `i` by 1.
   - Now match this `))` token: if `open > 0`, decrement `open`; otherwise insert a `(` before it (`ans++`).
3. After the scan, each remaining open `(` needs `))`: `ans += 2 * open`.

**Why greedy works:** a lone `)` that isn't immediately followed by another `)` can never become part of a consecutive `))` with any later existing character (something non-`)` sits between them), so inserting a `)` right there is forced. Similarly, a `))` with no unmatched `(` before it can only be fixed by inserting a `(` earlier — again forced. Every insertion we make is unavoidable, so the total is minimal.

**Example:** `s = "))())("`
- `i=0`: `))` → no open `(`, insert `(` → `ans=1`. `i=2`.
- `i=2`: `(` → `open=1`. `i=3`.
- `i=3`: `))` → match, `open=0`. `i=5`.
- `i=5`: `(` → `open=1`. `i=6`.
- End: `ans += 2*1` → `ans=3`. ✅

## Complexity Analysis

Time Complexity: O(n) — single pass over the string.
Space Complexity: O(1) — only two integer counters.

## Edge Cases

- **Lone `)` between `(`s, e.g. `"()()"`:** each `)` is followed by `(`, so each needs an extra `)` inserted — must check the *next* character, not just count `)`s.
- **Leading closers, e.g. `")))"`:** first `))` needs a `(` inserted; the final lone `)` needs both a `)` and a `(` → 3.
- **Only openers, e.g. `"((("`:** nothing closes them; add `2 * open` at the end.
- **Last character is `)`:** be careful with `i+1` going out of bounds.
- **Single character inputs `"("` / `")"`:** answers 2 and 2 respectively.
