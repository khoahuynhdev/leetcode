# 3614. Process String with Special Operations II

[LeetCode Link](https://leetcode.com/problems/process-string-with-special-operations-ii/)

Difficulty: Hard
Topics: String, Simulation

Acceptance Rate: 26.9%

## Hints

### Hint 1

Read the constraints before reaching for the obvious approach. The string `s` is short
(at most `10^5`), but the final `result` can grow to `10^15` characters because of the
duplication operator `'#'`. Actually building `result` is impossible — you would run out
of memory long before you finished. Ask yourself: do you really need the *whole* string,
or just *one* character of it?

### Hint 2

Since you only need the `k`-th character, think about working **backwards**. If you knew
the length of `result` after every operation, could you trace where index `k` "came from"
as you undo each operation one at a time? This is a classic trick for problems where the
structure explodes in size but only a single position is queried.

### Hint 3

Do two passes. **Pass 1 (forward):** compute the length of `result` after each operation
— this is cheap and never builds the string. **Pass 2 (backward):** start with the target
index `pos = k` and walk the operations from last to first, rewriting `pos` to be the index
it occupied in the *previous* `result`. Each operator has a simple inverse mapping on the
index:

- A letter appended at the end occupies index `prevLen`; if `pos == prevLen`, that letter
  *is* your answer.
- `'#'` (duplicate) doubled the string, so if `pos >= prevLen`, subtract `prevLen` to fold
  it back into the first half.
- `'%'` (reverse) maps `pos` to `prevLen - 1 - pos`.
- `'*'` (remove) only chopped the tail, so a valid `pos` is unchanged.

When `pos` lands exactly on an appended letter, you are done.

## Approach

The whole difficulty here is the size of `result`. Letters add one character, `'*'` removes
one, `'%'` is length-preserving, and `'#'` *doubles* the length. After enough `'#'`
operations the string is astronomically large, so simulation is out. But the question only
asks for a single index `k`, which lets us avoid materializing the string entirely.

**Pass 1 — record lengths.** Walk `s` left to right and maintain a running length `cur`:

- lowercase letter: `cur++`
- `'*'`: `cur--` if `cur > 0` (otherwise it stays 0)
- `'#'`: `cur *= 2`
- `'%'`: unchanged

Store `lengths[i]` = length of `result` immediately after processing `s[i]`. The final
length is `cur`. If `k >= cur`, index `k` is out of bounds and the answer is `'.'`.

**Pass 2 — trace the index backwards.** Set `pos = k` (a valid index into the final
`result`). Iterate `i` from `n-1` down to `0`. Let `prev` be the length *before* operation
`i`, which is `lengths[i-1]` (or `0` when `i == 0`). Rewrite `pos` according to the inverse
of `s[i]`:

- **letter:** appending put that letter at index `prev`. If `pos == prev`, the letter
  `s[i]` is exactly the character we are tracking — return it. Otherwise `pos < prev`, and
  `pos` refers to a character already present before the append, so leave it unchanged.
- **`'*'`:** the operation removed the last character, so every surviving index is identical
  to its index in the previous string. Leave `pos` unchanged.
- **`'#'`:** the previous string was concatenated with itself. Indices `[0, prev)` come from
  the first copy; indices `[prev, 2*prev)` come from the second copy. If `pos >= prev`,
  subtract `prev` to map it back onto the original copy.
- **`'%'`:** the string was reversed, so index `pos` came from index `prev - 1 - pos`.

Because every valid `pos` either resolves to an appended letter or is carried back to the
start, by the time we finish the loop we will have returned a letter. (If we somehow fall
through, returning `'.'` is a safe guard.)

**Why it works:** at every step `pos` is maintained as "the index in the current `result`
of the character we ultimately want." The inverse maps above are exact, so following them
to the beginning lands us on the precise letter that produced that position. We never store
more than the `lengths` array.

**Walkthrough (Example 2):** `s = "cd%#*#"`, `k = 3`. Forward lengths are
`[1, 2, 2, 4, 3, 6]`, final length `6`, so `k = 3` is in bounds. Tracing backward with
`pos = 3`:

- `i=5 '#'`, `prev=3`: `pos >= 3` → `pos = 0`
- `i=4 '*'`, `prev=4`: unchanged → `pos = 0`
- `i=3 '#'`, `prev=2`: `pos < 2` → `pos = 0`
- `i=2 '%'`, `prev=2`: `pos = 2 - 1 - 0 = 1`
- `i=1 'd'`, `prev=1`: `pos == 1` → return `'d'`. ✓

## Complexity Analysis

Time Complexity: O(n), where `n = len(s)` — one forward pass to build lengths and one
backward pass to trace the index.
Space Complexity: O(n) for the `lengths` array.

## Edge Cases

- **`k` out of bounds:** if `k >= finalLength` (including the case where `result` ends up
  empty), return `'.'`. Example 3 (`"z*#"`, `k = 0`) ends with an empty string.
- **`'*'` on an empty string:** removing from an empty `result` is a no-op; guard with
  `if cur > 0` so the length never goes negative.
- **`'#'` on an empty string:** duplicating an empty string stays empty (`0 * 2 = 0`), which
  the length recurrence handles naturally.
- **Large indices:** `k` can be up to `10^15` and the length up to `10^15`, so use 64-bit
  integers (`int64`) for both the running length and `pos` to avoid overflow.
- **`'%'` mapping near the boundary:** the reverse map `prev - 1 - pos` must use `prev`, the
  length *before* the reverse (which equals the length after, since reverse preserves
  length), to stay within bounds.
